//go:build !integration

package immutable

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/immutable"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

var (
	ctx                       = context.Background()
	projectID                 = 1
	immutableRuleID			  = 100
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Immutable: mocks.MockImmutableClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_CreateImmuRule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.CreateImmuRuleParams{
		ProjectNameOrID: strconv.Itoa(projectID),
		ImmutableRule: &model.ImmutableRule{
			TagSelectors: []*model.ImmutableSelector{{
				Decoration: "matches",
				Kind:       "doublestar",
				Pattern:    "1.0.0",
			}},
			ScopeSelectors: map[string][]model.ImmutableSelector{
				"repository": {{
					Decoration: "repoMatches",
					Kind:       "doublestar",
					Pattern:    "**",
				}},
			},
		},
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("CreateImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&immutable.CreateImmuRuleCreated{}, nil)

	err := apiClient.CreateImmuRule(ctx, strconv.Itoa(projectID), params.ImmutableRule)

	require.NoError(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_CreateImmuRuleError(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.CreateImmuRuleParams{
		ProjectNameOrID: strconv.Itoa(projectID),
		ImmutableRule: &model.ImmutableRule{
			TagSelectors: []*model.ImmutableSelector{{
				Decoration: "matches",
				Kind:       "doublestar",
				Pattern:    "1.0.0",
			}},
			ScopeSelectors: map[string][]model.ImmutableSelector{
				"repository": {{
					Decoration: "repoMatches",
					Kind:       "doublestar",
					Pattern:    "**",
				}},
			},
		},
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("CreateImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &immutable.CreateImmuRuleBadRequest{})

	err := apiClient.CreateImmuRule(ctx, strconv.Itoa(projectID), params.ImmutableRule)

	require.Error(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_UpdateImmuRule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.UpdateImmuRuleParams{
		ImmutableRuleID: int64(immutableRuleID),
		ProjectNameOrID: strconv.Itoa(projectID),
		ImmutableRule: &model.ImmutableRule{
			TagSelectors: []*model.ImmutableSelector{{
				Decoration: "matches",
				Kind:       "doublestar",
				Pattern:    "1.0.0",
			}},
			ScopeSelectors: map[string][]model.ImmutableSelector{
				"repository": {{
					Decoration: "repoMatches",
					Kind:       "doublestar",
					Pattern:    "**",
				}},
			},
		},
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("UpdateImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&immutable.UpdateImmuRuleOK{}, nil)

	err := apiClient.UpdateImmuRule(ctx, strconv.Itoa(projectID), params.ImmutableRule, int64(immutableRuleID))

	require.NoError(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_UpdateImmuRuleNotFound(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.UpdateImmuRuleParams{
		ImmutableRuleID: int64(immutableRuleID),
		ProjectNameOrID: strconv.Itoa(projectID),
		ImmutableRule: &model.ImmutableRule{
			TagSelectors: []*model.ImmutableSelector{{
				Decoration: "matches",
				Kind:       "doublestar",
				Pattern:    "1.0.0",
			}},
			ScopeSelectors: map[string][]model.ImmutableSelector{
				"repository": {{
					Decoration: "repoMatches",
					Kind:       "doublestar",
					Pattern:    "**",
				}},
			},
		},
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("UpdateImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &errors.ErrNotFound{})

	err := apiClient.UpdateImmuRule(ctx, strconv.Itoa(projectID), params.ImmutableRule, int64(immutableRuleID))

	require.Error(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_UpdateImmuRuleUnauthorized(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.UpdateImmuRuleParams{
		ImmutableRuleID: int64(immutableRuleID),
		ProjectNameOrID: strconv.Itoa(projectID),
		ImmutableRule: &model.ImmutableRule{},
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("UpdateImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &errors.ErrUnauthorized{})

	err := apiClient.UpdateImmuRule(ctx, strconv.Itoa(projectID), params.ImmutableRule, int64(immutableRuleID))

	require.Error(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_DeleteImmuRule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &immutable.DeleteImmuRuleParams{
		ImmutableRuleID: int64(immutableRuleID),
		ProjectNameOrID: strconv.Itoa(projectID),
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("DeleteImmuRule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&immutable.DeleteImmuRuleOK{}, nil)

	err := apiClient.DeleteImmuRule(ctx, strconv.Itoa(projectID), int64(immutableRuleID))

	require.NoError(t, err)

	mockClient.Immutable.AssertExpectations(t)
}

func TestRESTClient_ListImmuRules(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	expectedListImmutableRule := model.ImmutableRule{
		Action:         "immutable",
		ID:             1,
		ScopeSelectors: map[string][]model.ImmutableSelector{},
		TagSelectors: []*model.ImmutableSelector{{
			Decoration: "matches",
			Kind:       "doublestar",
			Pattern:    "**",
		}},
		Template: "immutable_template",
	}

	params := &immutable.ListImmuRulesParams{
		Page:            &apiClient.Options.Page,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: strconv.Itoa(projectID),
		Q:               &apiClient.Options.Query,
		Sort:            &apiClient.Options.Sort,
		Context:         ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Immutable.On("ListImmuRules", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&immutable.ListImmuRulesOK{Payload: []*model.ImmutableRule{&expectedListImmutableRule}}, nil)

	immutableRules, err := apiClient.ListImmuRules(ctx, strconv.Itoa(projectID))

	require.NoError(t, err)

	require.Equal(t, expectedListImmutableRule, *immutableRules[0])

	mockClient.Immutable.AssertExpectations(t)
}
