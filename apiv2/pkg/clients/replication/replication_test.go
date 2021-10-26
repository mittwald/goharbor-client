//go:build !integyration

package replication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	replicationapi "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/replication"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

const (
	name        string = "example-replication"
	description string = "a test replication"
	ns          string = "test-namespace"
)

var (
	ctx               = context.Background()
	destRegistry      = &modelv2.Registry{ID: 1, Name: "reg1"}
	srcRegistry       = &modelv2.Registry{Name: "reg2"}
	replicateDeletion = true
	override          = true
	enablePolicy      = true
	filters           []*modelv2.ReplicationFilter
	trigger           = &modelv2.ReplicationTrigger{}
	destNamespace     = ns
	replication       = &modelv2.ReplicationPolicy{
		ReplicateDeletion: replicateDeletion,
		Description:       description,
		DestNamespace:     destNamespace,
		DestRegistry:      destRegistry,
		Enabled:           enablePolicy,
		Filters:           filters,
		Name:              name,
		Override:          override,
		SrcRegistry:       srcRegistry,
		Trigger:           trigger,
		ID:                0,
	}
	replExec = &modelv2.ReplicationExecution{
		ID:       1,
		PolicyID: 1,
	}
	startReplExec = &modelv2.StartReplicationExecution{
		PolicyID: 1,
	}
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Replication: mocks.MockReplicationClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_NewReplicationPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	destNamespace := ns
	description := description
	name := name

	createParams := &replicationapi.CreateReplicationPolicyParams{
		Policy:  replication,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("CreateReplicationPolicy", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.CreateReplicationPolicyCreated{}, nil)

	err := apiClient.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	mockClient.Replication.AssertExpectations(t)
	require.NoError(t, err)
}

func TestRESTClient_GetReplicationPolicyByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &replicationapi.GetReplicationPolicyParams{
		ID:      replication.ID,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("GetReplicationPolicy", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.GetReplicationPolicyOK{Payload: &modelv2.ReplicationPolicy{}}, nil)

	_, err := apiClient.GetReplicationPolicyByID(ctx, replication.ID)
	require.NoError(t, err)

	mockClient.Replication.AssertExpectations(t)
}

func TestRESTClient_DeleteReplicationPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &replicationapi.DeleteReplicationPolicyParams{
		ID:      replication.ID,
		Context: ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("DeleteReplicationPolicy", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.DeleteReplicationPolicyOK{}, nil)

	err := apiClient.DeleteReplicationPolicyByID(ctx, replication.ID)

	require.NoError(t, err)
	mockClient.Replication.AssertExpectations(t)
}

func TestRESTClient_UpdateReplicationPolicy(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &replicationapi.UpdateReplicationPolicyParams{
		ID:      replication.ID,
		Policy:  replication,
		Context: ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("UpdateReplicationPolicy", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.UpdateReplicationPolicyOK{}, nil)

	err := apiClient.UpdateReplicationPolicy(ctx, replication, replication.ID)

	require.NoError(t, err)

	mockClient.Replication.AssertExpectations(t)
}

func TestRESTClient_UpdateReplicationPolicy_NilParam(t *testing.T) {
	apiClient, _ := APIandMockClientsForTests()

	err := apiClient.UpdateReplicationPolicy(ctx, nil, replication.ID)

	require.Error(t, err)
	require.ErrorIs(t, &ErrReplicationNotProvided{}, err)
}

func TestRESTClient_ListReplicationExecutions(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &replicationapi.ListReplicationExecutionsParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		PolicyID: &replExec.ID,
		Sort:     &apiClient.Options.Sort,
		Status:   &replExec.Status,
		Trigger:  &replExec.Trigger,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("ListReplicationExecutions", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.ListReplicationExecutionsOK{
			Payload: []*modelv2.ReplicationExecution{replExec},
		}, nil)

	_, err := apiClient.ListReplicationExecutions(ctx, &replExec.ID, &replExec.Status, &replExec.Trigger)
	require.NoError(t, err)

	mockClient.Replication.AssertExpectations(t)
}

func TestRESTClient_TriggerReplicationExecution(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	startParams := &replicationapi.StartReplicationParams{
		Execution: startReplExec,
		Context:   ctx,
	}

	startParams.WithTimeout(apiClient.Options.Timeout)

	re := &modelv2.StartReplicationExecution{
		PolicyID: replExec.PolicyID,
	}

	mockClient.Replication.On("StartReplication", startParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&replicationapi.StartReplicationCreated{}, nil)

	err := apiClient.TriggerReplicationExecution(ctx, re)

	require.NoError(t, err)

	mockClient.Replication.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutionByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &replicationapi.GetReplicationExecutionParams{
		ID:      replExec.ID,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Replication.On("GetReplicationExecution", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&replicationapi.GetReplicationExecutionOK{Payload: &modelv2.ReplicationExecution{ID: replExec.ID}}, nil)

	_, err := apiClient.GetReplicationExecutionByID(ctx, replExec.ID)

	require.NoError(t, err)

	mockClient.Replication.AssertExpectations(t)
}
