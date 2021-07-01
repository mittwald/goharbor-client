// +build !integration

package replication

import (
	"context"
	"net/http"
	"testing"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	replicationapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/replication"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	name        string = "example-replication"
	description string = "a test replication"
	ns          string = "test-namespace"
)

var (
	authInfo = runtimeclient.BasicAuth("foo", "bar")

	destRegistry      = &legacymodel.Registry{ID: 1, Name: "reg1"}
	srcRegistry       = &legacymodel.Registry{Name: "reg2"}
	replicateDeletion = true
	override          = true
	enablePolicy      = true
	filters           []*legacymodel.ReplicationFilter
	trigger           = &legacymodel.ReplicationTrigger{}
	destNamespace     = ns
	replication       = &legacymodel.ReplicationPolicy{
		Deletion:      replicateDeletion,
		Description:   description,
		DestNamespace: destNamespace,
		DestRegistry:  destRegistry,
		Enabled:       enablePolicy,
		Filters:       filters,
		Name:          name,
		Override:      override,
		SrcRegistry:   srcRegistry,
		Trigger:       trigger,
		ID:            0,
	}
	replExec = &modelv2.ReplicationExecution{
		ID:       1,
		PolicyID: 1,
	}
	startReplExec = &modelv2.StartReplicationExecution{
		PolicyID: 1,
	}
)

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildV2ClientWithMocks(r *mocks.MockReplicationClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Artifact:    &mocks.MockArtifactClientService{},
		Auditlog:    &mocks.MockAuditlogClientService{},
		Icon:        &mocks.MockIconClientService{},
		Preheat:     &mocks.MockPreheatClientService{},
		Project:     &mocks.MockProjectClientService{},
		Repository:  &mocks.MockRepositoryClientService{},
		Scan:        &mocks.MockScanClientService{},
		Replication: r,
	}
}

func TestNewClient(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	require.NotNil(t, cl)
	assert.NotNil(t, cl.AuthInfo)
	assert.NotNil(t, cl.V2Client)
	assert.NotNil(t, cl.LegacyClient)
}

func TestRESTClient_NewReplicationPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	destNamespace := ns
	description := description
	name := name

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, r, replication)
}

func TestRESTClient_NewReplicationPolicy_ErrOnPOST(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()
	destNamespace := ns
	description := description
	name := name

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      "unit test error",
		Code:          500,
	})

	r, err := cl.NewReplicationPolicy(
		ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationInternalErrors{}, err)
	}
}

func TestRESTClient_NewReplicationPolicy_ErrOnGET(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      "unauthorized",
		Code:          401,
	})

	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationUnauthorized{}, err)
	}
}

func TestRESTClient_NewReplicationPolicy_ErrReplicationNameAlreadyExists(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	destNamespace := ns
	description := description
	name := name

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &products.PostReplicationPoliciesConflict{})

	_, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy,
		filters,
		trigger, destNamespace, description, name)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNameAlreadyExists{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	r, err := cl.GetReplicationPolicy(ctx, replication.Name)

	assert.NoError(t, err)
	assert.Equal(t, replication, r)

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationPolicyByID(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replication.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &legacymodel.ReplicationPolicy{}}, nil)
	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy,
		filters,
		trigger, destNamespace, description, name)
	assert.NoError(t, err)

	_, err = cl.GetReplicationPolicyByID(ctx, r.ID)
	assert.NoError(t, err)
}

func TestRESTClient_GetReplicationPolicyByID_ErrReplicationIllegalIDFormat(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replication.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusBadRequest})

	_, err := cl.GetReplicationPolicyByID(ctx, replication.ID)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationIllegalIDFormat{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationPolicy_EmptyName(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	r, err := cl.GetReplicationPolicy(context.Background(), "")

	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_GetReplicationPolicy_NotFound(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{}}, nil)

	r, err := cl.GetReplicationPolicy(ctx, replication.Name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplicationPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)
	p.On("DeleteReplicationPoliciesID", &products.DeleteReplicationPoliciesIDParams{
		ID:      replication.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.DeleteReplicationPoliciesIDOK{}, nil)

	err := cl.DeleteReplicationPolicy(ctx, replication)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_DeleteReplicationPolicy_NotFound(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{}}, nil)

	err := cl.DeleteReplicationPolicy(ctx, replication)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplicationPolicy_NilParam(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.DeleteReplicationPolicy(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_DeleteReplicationPolicy_ErrReplicationMismatch(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{{
			ID:   1,
			Name: replication.Name,
		}}}, nil)

	err := cl.DeleteReplicationPolicy(ctx, replication)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationMismatch{}, err)
	}
}

func TestRESTClient_UpdateReplicationPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	p.On("PutReplicationPoliciesID", &products.PutReplicationPoliciesIDParams{
		ID:      replication.ID,
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PutReplicationPoliciesIDOK{}, nil)

	err := cl.UpdateReplicationPolicy(ctx, replication)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_UpdateReplicationPolicy_NilParam(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.UpdateReplicationPolicy(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_UpdateReplicationPolicy_ErrReplicationMismatch(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{{
			ID:   1,
			Name: replication.Name,
		}}}, nil)

	err := cl.UpdateReplicationPolicy(ctx, replication)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationMismatch{}, err)
	}
}

func TestRESTClient_UpdateReplicationPolicy_NotFound(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{}}, nil)

	err := cl.UpdateReplicationPolicy(ctx, replication)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateReplicationPolicy_ErrReplicationIDNotExists(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &replication.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	p.On("PutReplicationPoliciesID", &products.PutReplicationPoliciesIDParams{
		ID:      replication.ID,
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &products.PutReplicationPoliciesIDNotFound{})

	err := cl.UpdateReplicationPolicy(ctx, replication)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationIDNotExists{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("ListReplicationExecutions",
		mock.Anything,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&replicationapi.ListReplicationExecutionsOK{Payload: []*modelv2.ReplicationExecution{}}, nil)

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	assert.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions_ErrReplicationIllegalIDFormat(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("ListReplicationExecutions",
		mock.Anything,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusBadRequest})

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationIllegalIDFormat{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions_ErrReplicationUnauthorized(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("ListReplicationExecutions",
		mock.Anything,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationUnauthorized{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions_ErrReplicationNoPermission(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("ListReplicationExecutions",
		mock.Anything,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusForbidden})

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNoPermission{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions_ErrReplicationInternalErrors(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("ListReplicationExecutions",
		mock.Anything,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusInternalServerError})

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationInternalErrors{}, err)
	}

	r.AssertExpectations(t)
}

func TestRESTClient_TriggerReplicationExecution(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()
	destRegistry := &legacymodel.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &legacymodel.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true

	var filters []*legacymodel.ReplicationFilter

	trigger := &legacymodel.ReplicationTrigger{}
	destNamespace := ns
	description := description
	name := name

	startReplParams := &replicationapi.StartReplicationParams{
		Execution: startReplExec,
		Context:   ctx,
	}

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy:  replication,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*legacymodel.ReplicationPolicy{replication}}, nil)

	r.On("StartReplication", startReplParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&replicationapi.StartReplicationCreated{}, nil)

	rep, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	assert.NoError(t, err)
	assert.Equal(t, rep, replication)

	err = cl.TriggerReplicationExecution(ctx, startReplExec)

	assert.NoError(t, err)

	p.AssertExpectations(t)
	r.AssertExpectations(t)
}

func TestRESTClient_TriggerReplicationExecution_ErrReplicationExecutionNotProvided(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.TriggerReplicationExecution(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationExecutionNotProvided{}, err)
	}
}

func TestRESTClient_GetReplicationExecutionByID(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("GetReplicationExecution", &replicationapi.GetReplicationExecutionParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&replicationapi.GetReplicationExecutionOK{Payload: &modelv2.ReplicationExecution{ID: replExec.ID}}, nil)

	_, err := cl.GetReplicationExecutionByID(ctx, replExec.ID)

	assert.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutionByID_ErrReplicationIllegalIDFormat(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockReplicationClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	r.On("GetReplicationExecution", &replicationapi.GetReplicationExecutionParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		nil, &runtime.APIError{Code: http.StatusBadRequest})

	_, err := cl.GetReplicationExecutionByID(ctx, replExec.ID)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationIllegalIDFormat{}, err)
	}

	r.AssertExpectations(t)
}

func TestErrReplicationExecutionNotProvided_Error(t *testing.T) {
	var e ErrReplicationExecutionNotProvided

	assert.Equal(t, ErrReplicationExecutionNotProvidedMsg, e.Error())
}

func TestErrReplicationExecutionReplicationIDMismatch_Error(t *testing.T) {
	var e ErrReplicationExecutionReplicationIDMismatch

	assert.Equal(t, ErrReplicationExecutionReplicationIDMismatchMsg, e.Error())
}

func TestErrReplicationIllegalIDFormat_Error(t *testing.T) {
	var e ErrReplicationIllegalIDFormat

	assert.Equal(t, ErrReplicationIllegalIDFormatMsg, e.Error())
}

func TestErrReplicationInternalErrors_Error(t *testing.T) {
	var e ErrReplicationInternalErrors

	assert.Equal(t, ErrReplicationInternalErrorsMsg, e.Error())
}

func TestErrReplicationNameAlreadyExists_Error(t *testing.T) {
	var e ErrReplicationNameAlreadyExists

	assert.Equal(t, ErrReplicationNameAlreadyExistsMsg, e.Error())
}

func TestErrReplicationNoPermission_Error(t *testing.T) {
	var e ErrReplicationNoPermission

	assert.Equal(t, ErrReplicationNoPermissionMsg, e.Error())
}

func TestErrReplicationNotFound_Error(t *testing.T) {
	var e ErrReplicationNotFound

	assert.Equal(t, ErrReplicationNotFoundMsg, e.Error())
}

func TestErrReplicationNotProvided_Error(t *testing.T) {
	var e ErrReplicationNotProvided

	assert.Equal(t, ErrReplicationNotProvidedMsg, e.Error())
}

func TestErrReplicationUnauthorized_Error(t *testing.T) {
	var e ErrReplicationUnauthorized

	assert.Equal(t, ErrReplicationUnauthorizedMsg, e.Error())
}

func TestErrReplicationIDNotExists_Error(t *testing.T) {
	var e ErrReplicationIDNotExists

	assert.Equal(t, ErrReplicationIDNotExistsMsg, e.Error())
}

func TestErrReplicationMismatch_Error(t *testing.T) {
	var e ErrReplicationMismatch

	assert.Equal(t, ErrReplicationMismatchMsg, e.Error())
}
