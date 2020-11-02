// +build integration

package retention

import (
	"context"
	"flag"
	"net/url"
	"testing"

	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	pc "github.com/mittwald/goharbor-client/v3/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	harborVersion       = flag.String("version", "2.1.0",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 2.1.0")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

const (
	projectName string = "test-project"
)

func TestAPIRetentionNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, 1)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProjectByName(ctx, projectName)

	rep := &model.RetentionPolicy{
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
			Ref:   int64(p.ProjectID),
		},
		Trigger: &model.RetentionRuleTrigger{
			Kind:     "Schedule",
			Settings: map[string]interface{}{"cron": "0 * * * *"},
		},
	}

	err = c.NewRetentionPolicy(ctx, rep)

	require.NoError(t, err)

	require.Nil(t, err)
}

func TestAPIRetentionGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, 1)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProjectByName(ctx, projectName)

	rep := &model.RetentionPolicy{
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
			Ref:   int64(p.ProjectID),
		},
		Trigger: &model.RetentionRuleTrigger{
			Kind:     "Schedule",
			Settings: map[string]interface{}{"cron": "0 * * * *"},
		},
	}

	err = c.NewRetentionPolicy(ctx, rep)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, p)
	require.NoError(t, err)
	require.NotNil(t, rp)
}
