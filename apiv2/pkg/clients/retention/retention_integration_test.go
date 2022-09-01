//go:build integration

package retention

import (
	"context"
	"testing"

	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"

	pc "github.com/testwill/goharbor-client/v5/apiv2/pkg/clients/project"
)

var storageLimit int64 = 1

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

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := pc.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimit,
	})
	require.NoError(t, err)

	p, err := pc.GetProject(ctx, projectName)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	p, err = pc.GetProject(ctx, projectName)

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)

	require.Nil(t, err)
}

func TestAPIRetentionUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := pc.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimit,
	})
	require.NoError(t, err)

	p, err := pc.GetProject(ctx, projectName)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	p, err = pc.GetProject(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, projectName)

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

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := pc.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimit,
	})
	require.NoError(t, err)

	p, err := pc.GetProject(ctx, projectName)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	ret := newTestRetention(int64(p.ProjectID))

	err = c.NewRetentionPolicy(ctx, &ret)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProject(ctx, projectName)

	require.NoError(t, err)
	require.Nil(t, err)

	err = c.DeleteRetentionPolicyByID(ctx, rp.ID)

	require.NoError(t, err)

	deleted, err := c.GetRetentionPolicyByProject(ctx, projectName)

	require.Error(t, err)
	require.ErrorIs(t, err, &ErrRetentionInternalErrors{})
	require.Nil(t, deleted)
}
