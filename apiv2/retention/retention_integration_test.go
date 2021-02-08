// +build integration

package retention

import (
	"context"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	pc "github.com/mittwald/goharbor-client/v3/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                      = url.Parse(integrationtest.Host)
	legacySwaggerClient       = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient           = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                  = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimit        int64 = 1
)

const (
	projectName string = "test-project"
)

func newTestRetention(projectID int64) model.RetentionPolicy {
	return model.RetentionPolicy{
		Algorithm: AlgorithmOr,
		Rules: []*model.RetentionRule{{
			Action:   "retain",
			Disabled: false,
			Params: map[string]interface{}{
				PolicyTemplateDaysSinceLastPush.String(): 1,
			},
			ScopeSelectors: map[string][]model.RetentionSelector{
				"repository": {{
					Decoration: ScopeSelectorRepoMatches.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
					Extras:     "", // The "Extras" field is unused for scope selectors.
				}}},
			TagSelectors: []*model.RetentionSelector{{
				Decoration: TagSelectorMatches.String(),
				Extras:     ToTagSelectorExtras(true),
				Kind:       SelectorTypeDefault,
				Pattern:    "**",
			}},
			Template: PolicyTemplateDaysSinceLastPush.String(),
		}},
		Scope: &model.RetentionPolicyScope{
			Level: "project",
			Ref:   projectID,
		},
		Trigger: &model.RetentionRuleTrigger{
			Kind:     "Schedule",
			Settings: map[string]interface{}{"cron": "0 * * * *"},
		},
	}
}

func TestAPIRetentionNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	ret := newTestRetention(int64(p.ProjectID))

	p, err = pc.GetProjectByName(ctx, projectName)

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)

	require.Nil(t, err)
}

func TestAPIRetentionGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProjectByName(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)
	require.NoError(t, err)
	require.NotNil(t, rp)
}

func TestAPIRetentionUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProjectByName(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)

	require.NoError(t, err)
	require.Nil(t, err)

	changed := rp

	changed.Rules = []*model.RetentionRule{{
		Action:   "retain",
		Disabled: true,
		Params: map[string]interface{}{
			PolicyTemplateDaysSinceLastPull.String(): 2,
		},
		ScopeSelectors: map[string][]model.RetentionSelector{
			"repository": {{
				Decoration: ScopeSelectorRepoExcludes.String(),
				Kind:       SelectorTypeDefault,
				Pattern:    "**",
				Extras:     "", // The "Extras" field is unused for scope selectors.
			}}},
		TagSelectors: []*model.RetentionSelector{{
			Decoration: TagSelectorExcludes.String(),
			Extras:     ToTagSelectorExtras(false),
			Kind:       SelectorTypeDefault,
			Pattern:    "**",
		}},
		Template: PolicyTemplateDaysSinceLastPull.String(),
	},
	}

	err = c.UpdateRetentionPolicy(ctx, changed)
	require.NoError(t, err)
}

func TestAPIRetentionDelete(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProjectByName(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)

	require.NoError(t, err)
	require.Nil(t, err)

	err = c.DisableRetentionPolicy(ctx, rp)

	require.NoError(t, err)

	disabled, err := c.GetRetentionPolicyByProject(ctx, p)

	require.NoError(t, err)
	require.Equal(t, 0, len(disabled.Rules))
}
