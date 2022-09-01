//go:build !integration

package retention

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	projectmeta "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/project_metadata"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/retention"
	"github.com/testwill/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/common"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
)

var ctx = context.Background()

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Retention:       mocks.MockRetentionClientService{},
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestEvaluateRetentionRuleParams(t *testing.T) {
	t.Run("WithParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{
			PolicyTemplateLatestPushedArtifacts: 1,
			PolicyTemplateLatestPulledArtifacts: 2,
			PolicyTemplateDaysSinceLastPush:     3,
			PolicyTemplateDaysSinceLastPull:     4,
		}
		e, err := evaluateRetentionRuleParams(params)
		require.NoError(t, err)
		require.NotNil(t, e)
	})

	t.Run("WithoutParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{}

		e, err := evaluateRetentionRuleParams(params)

		require.Error(t, err)
		require.Nil(t, e)
	})

	t.Run("InvalidParams", func(t *testing.T) {
		params := map[PolicyTemplate]interface{}{
			"foo": "bar",
		}

		e, err := evaluateRetentionRuleParams(params)

		require.Error(t, err)
		require.Nil(t, e)
	})
}

func TestRESTClient_NewRetentionPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &retention.CreateRetentionParams{
		Policy: &modelv2.RetentionPolicy{
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
					}},
				},
				TagSelectors: []*modelv2.RetentionSelector{{
					Decoration: TagSelectorMatches.String(),
					Kind:       SelectorTypeDefault,
					Pattern:    "**",
				}},
				Template: PolicyTemplateDaysSinceLastPush.String(),
			}},
			Scope: &modelv2.RetentionPolicyScope{
				Level: "project",
				Ref:   0,
			},
			Trigger: &modelv2.RetentionRuleTrigger{
				Kind:     "Schedule", // Trigger kind is _always_ 'Schedule'.
				Settings: map[string]interface{}{"cron": "0 * * * *"},
			},
		},
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Retention.On("CreateRetention", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.CreateRetentionCreated{}, &runtime.APIError{Code: http.StatusCreated})

	err := apiClient.NewRetentionPolicy(ctx, createParams.Policy)

	require.NoError(t, err)
	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_GetRetentionPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	retentionIDPtr := "1"

	project := &modelv2.Project{
		Deleted: false,
		Metadata: &modelv2.ProjectMetadata{
			RetentionID: &retentionIDPtr,
		},
		Name:      "test-project",
		ProjectID: 1,
	}

	getMetaParams := &projectmeta.GetProjectMetadataParams{
		MetaName:        common.ProjectMetadataKeyRetentionID.String(),
		ProjectNameOrID: project.Name,
		Context:         ctx,
	}

	getMetaParams.WithTimeout(apiClient.Options.Timeout)

	getRetentionParams := &retention.GetRetentionParams{
		ID:      1,
		Context: ctx,
	}

	getRetentionParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("GetProjectMetadata", getMetaParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.GetProjectMetadataOK{Payload: map[string]string{
			common.ProjectMetadataKeyRetentionID.String(): retentionIDPtr,
		}}, nil)

	mockClient.Retention.On("GetRetention", getRetentionParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.GetRetentionOK{}, nil)

	_, err := apiClient.GetRetentionPolicyByProject(ctx, project.Name)

	require.NoError(t, err)
}

func TestRESTClient_UpdateRetentionPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     nil,
		Scope:     nil,
		Trigger:   nil,
	}

	updateParams := &retention.UpdateRetentionParams{
		ID:      1,
		Policy:  policy,
		Context: ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Retention.On("UpdateRetention", updateParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.UpdateRetentionOK{}, nil)

	err := apiClient.UpdateRetentionPolicy(ctx, updateParams.Policy)

	require.NoError(t, err)
	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_DeleteRetentionPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	ctx := context.Background()

	policy := &modelv2.RetentionPolicy{
		Algorithm: "",
		ID:        1,
		Rules:     []*modelv2.RetentionRule{},
	}

	deleteParams := &retention.DeleteRetentionParams{
		ID:      1,
		Context: ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Retention.On("DeleteRetention", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&retention.DeleteRetentionOK{}, nil)

	err := apiClient.DeleteRetentionPolicyByID(ctx, policy.ID)

	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}
