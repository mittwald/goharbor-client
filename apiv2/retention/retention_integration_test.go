////go:build integration

package retention

import (
	"context"
	"net/url"
	"testing"

	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
	pc "github.com/mittwald/goharbor-client/v4/apiv2/project"
)

var (
	u, _                      = url.Parse(integrationtest.Host)
	legacySwaggerClient       = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient           = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                  = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimit        int64 = 1
	opts                      = config.Options{}
	defaultOpts               = opts.Defaults()
)

const (
	projectName string = "test-project"
)

func newTestRetention(projectID int64) modelv2.RetentionPolicy {
	return modelv2.RetentionPolicy{
		Algorithm: AlgorithmOr,
		Rules: []*modelv2.RetentionRule{{
			Action:   "retain",
			Disabled: false,
			Params: map[string]interface{}{
				PolicyTemplateDaysSinceLastPush.String(): 1,
			},
			ScopeSelectors: map[string][]modelv2.RetentionSelector{
				"repository": {{
					Decoration: ScopeSelectorRepoMatches.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
					Extras:     "", // The "Extras" field is unused for scope selectors.
				}},
			},
			TagSelectors: []*modelv2.RetentionSelector{{
				Decoration: TagSelectorMatches.String(),
				Extras:     ToTagSelectorExtras(true),
				Kind:       SelectorTypeDefault,
				Pattern:    "**",
			}},
			Template: PolicyTemplateDaysSinceLastPush.String(),
		}},
		Scope: &modelv2.RetentionPolicyScope{
			Level: "project",
			Ref:   projectID,
		},
		Trigger: &modelv2.RetentionRuleTrigger{
			Kind:     "Schedule",
			Settings: map[string]interface{}{"cron": "0 * * * *"},
		},
	}
}

func TestAPIRetentionNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	pc := pc.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	ret := newTestRetention(int64(p.ProjectID))

	p, err = pc.GetProject(ctx, projectName)

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)

	require.Nil(t, err)
}

func TestAPIRetentionUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	pc := pc.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProject(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)

	require.NoError(t, err)
	require.Nil(t, err)

	changed := rp

	changed.Rules = []*modelv2.RetentionRule{
		{
			Action:   "retain",
			Disabled: true,
			Params: map[string]interface{}{
				PolicyTemplateDaysSinceLastPull.String(): 2,
			},
			ScopeSelectors: map[string][]modelv2.RetentionSelector{
				"repository": {{
					Decoration: ScopeSelectorRepoExcludes.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
					Extras:     "", // The "Extras" field is unused for scope selectors.
				}},
			},
			TagSelectors: []*modelv2.RetentionSelector{{
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

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	pc := pc.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := pc.NewProject(ctx, projectName, &storageLimit)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProject(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)

	require.NoError(t, err)
	require.Nil(t, err)

	err = c.DeleteRetentionPolicyByID(ctx, rp.ID)

	require.NoError(t, err)

	deleted, err := c.GetRetentionPolicyByProject(ctx, p)

	require.Error(t, err)
	require.ErrorIs(t, err, &ErrRetentionInternalErrors{})
	require.Nil(t, deleted)
}
