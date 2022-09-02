package webhook

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/webhook"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
)

var (
	exampleProjectID = 1
	ctx              = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Webhook: mocks.MockWebhookClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_ListProjectWebhookPolicies(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	expectedWebhookPolicies := []*modelv2.WebhookPolicy{
		{
			ID:        42,
			Name:      "example-policy",
			ProjectID: int64(exampleProjectID),
		},
	}

	listParams := &webhook.ListWebhookPoliciesOfProjectParams{
		Page:            &apiClient.Options.Page,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: util.ProjectIDAsString(1),
		Q:               &apiClient.Options.Query,
		Sort:            &apiClient.Options.Sort,
		Context:         ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Webhook.On("ListWebhookPoliciesOfProject", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&webhook.ListWebhookPoliciesOfProjectOK{Payload: expectedWebhookPolicies}, nil)

	webhookPolicies, err := apiClient.ListProjectWebhookPolicies(ctx, exampleProjectID)

	require.NoError(t, err)

	require.Equal(t, expectedWebhookPolicies, webhookPolicies)

	mockClient.Webhook.AssertExpectations(t)
}

func TestRESTClient_AddProjectWebhookPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	newPolicy := &modelv2.WebhookPolicy{
		Enabled: true,
		Name:    "my-policy",
		Targets: []*modelv2.WebhookTargetObject{{
			Address: "http://example-webhook.com",
		}},
		EventTypes: []string{
			"SCANNING_FAILED",
			"SCANNING_COMPLETED",
		},
	}

	updateParams := &webhook.CreateWebhookPolicyOfProjectParams{
		ProjectNameOrID: fmt.Sprintf("%d", exampleProjectID),
		Policy:          newPolicy,
		Context:         ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Webhook.On("CreateWebhookPolicyOfProject", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&webhook.CreateWebhookPolicyOfProjectCreated{}, nil)

	err := apiClient.AddProjectWebhookPolicy(ctx, exampleProjectID, newPolicy)

	require.NoError(t, err)

	mockClient.Webhook.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectWebhookPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	newPolicy := &modelv2.WebhookPolicy{
		Enabled: true,
		Name:    "my-policy",
		Targets: []*modelv2.WebhookTargetObject{{
			Address: "http://example-webhook.com",
		}},
		EventTypes: []string{
			"SCANNING_FAILED",
			"SCANNING_COMPLETED",
		},
	}

	updateParams := &webhook.UpdateWebhookPolicyOfProjectParams{
		ProjectNameOrID: fmt.Sprintf("%d", exampleProjectID),
		Policy:          newPolicy,
		Context:         ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Webhook.On("UpdateWebhookPolicyOfProject", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&webhook.UpdateWebhookPolicyOfProjectOK{}, nil)

	err := apiClient.UpdateProjectWebhookPolicy(ctx, exampleProjectID, newPolicy)

	require.NoError(t, err)

	mockClient.Webhook.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectWebhookPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	const examplePolicyID = 42

	deleteParams := &webhook.DeleteWebhookPolicyOfProjectParams{
		ProjectNameOrID: fmt.Sprintf("%d", exampleProjectID),
		WebhookPolicyID: examplePolicyID,
		Context:         ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Webhook.On("DeleteWebhookPolicyOfProject", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&webhook.DeleteWebhookPolicyOfProjectOK{}, nil)

	err := apiClient.DeleteProjectWebhookPolicy(ctx, exampleProjectID, examplePolicyID)

	require.NoError(t, err)

	mockClient.Webhook.AssertExpectations(t)
}
