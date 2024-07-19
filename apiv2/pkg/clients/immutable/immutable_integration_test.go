//go:build integration

package immutable

import (
	"context"
	"testing"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

var projectName = "test-project"

func TestAPIImmutableListImmutableRules(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &model.ProjectReq{
		ProjectName: projectName,
	})
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	immutableRule := model.ImmutableRule{
		ScopeSelectors: map[string][]model.ImmutableSelector{},
		TagSelectors: []*model.ImmutableSelector{{
			Decoration: "matches",
			Kind:       "doublestar",
			Pattern:    "**",
		}},
	}

	err = c.CreateImmuRule(ctx, projectName, &immutableRule)
	require.NoError(t, err)

	listedImmutableRules, err := c.ListImmuRules(ctx, projectName)

	require.NoError(t, err)

	listedImmutableRuleTag := listedImmutableRules[0].TagSelectors

	require.Equal(t, listedImmutableRuleTag[0], immutableRule.TagSelectors[0])

	immuRuleID := listedImmutableRules[0].ID

	c.DeleteImmuRule(ctx, projectName, immuRuleID)

	checkDeletedImmuRules, err := c.ListImmuRules(ctx, projectName)

	require.Empty(t, checkDeletedImmuRules)
}

func TestAPIImmutableUpdateImmutableRules(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &model.ProjectReq{
		ProjectName: projectName,
	})
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	createImmutableRule := model.ImmutableRule{
		ScopeSelectors: map[string][]model.ImmutableSelector{},
		TagSelectors: []*model.ImmutableSelector{{
			Decoration: "matches",
			Kind:       "doublestar",
			Pattern:    "1.0.0",
		}},
	}

	updateImmutableRule := model.ImmutableRule{
		ScopeSelectors: map[string][]model.ImmutableSelector{},
		TagSelectors: []*model.ImmutableSelector{{
			Decoration: "matches",
			Kind:       "doublestar",
			Pattern:    "2.0.0",
		}},
	}

	err = c.CreateImmuRule(ctx, projectName, &createImmutableRule)
	require.NoError(t, err)

	listedImmutableRules, err := c.ListImmuRules(ctx, projectName)

	require.NoError(t, err)

	immuRuleID := listedImmutableRules[0].ID

	err = c.UpdateImmuRule(ctx, projectName, &updateImmutableRule, immuRuleID)

	require.NoError(t, err)

	listedImmutableRules, err = c.ListImmuRules(ctx, projectName)

	require.NoError(t, err)

	listedImmutableRuleTag := listedImmutableRules[0].TagSelectors

	require.Equal(t, listedImmutableRuleTag[0], updateImmutableRule.TagSelectors[0])


	c.DeleteImmuRule(ctx, projectName, immuRuleID)

	checkDeletedImmuRules, err := c.ListImmuRules(ctx, projectName)

	require.Empty(t, checkDeletedImmuRules)
}