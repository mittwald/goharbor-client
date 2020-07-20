// +build !integration

package replication

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
	"github.com/mittwald/goharbor-client/mocks"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	name        string = "testreplication"
	description string = "a test replication"
	ns          string = "test-namespace"
)

var authInfo = runtimeclient.BasicAuth("foo", "bar")

func TestNewClient(t *testing.T) {
	swaggerClient := client.New(runtimeclient.New("foobar:30002", "/api",
		[]string{"http"}), strfmt.Default)
	authInfo := runtimeclient.BasicAuth("foo", "bar")
	c := NewClient(swaggerClient, authInfo)

	require.NotNil(t, c)
	assert.NotNil(t, c.AuthInfo)
	assert.NotNil(t, c.Client)

	assert.Equal(t, swaggerClient, c.Client)
}

func TestRESTClient_NewReplicationPolicy(t *testing.T) {
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	trigger := &model.ReplicationTrigger{}

	var filters []*model.ReplicationFilter

	destNamespace := ns
	description := description
	name := name
	returnedReplication := &model.ReplicationPolicy{
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
	}

	p := &mocks.MockClientService{}
	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{returnedReplication}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, r, returnedReplication)
}

func TestRESTClient_NewReplicationPolicy_ErrOnPOST(t *testing.T) {
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	trigger := &model.ReplicationTrigger{}

	var filters []*model.ReplicationFilter

	destNamespace := ns
	description := description
	name := name

	p := &mocks.MockClientService{}
	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      "unit test error",
		Code:          500,
	})

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

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
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	trigger := &model.ReplicationTrigger{}

	var filters []*model.ReplicationFilter

	destNamespace := ns
	description := description
	name := name

	p := &mocks.MockClientService{}
	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      "unauthorized",
		Code:          401,
	})

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationUnauthorized{}, err)
	}
}

func TestRESTClient_GetReplicationPolicy(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{repl}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplicationPolicy(ctx, repl.Name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, repl, r)
}

func TestRESTClient_GetReplicationPolicyByID(t *testing.T) {
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true

	var filters []*model.ReplicationFilter

	trigger := &model.ReplicationTrigger{}
	destNamespace := ns
	description := description
	name := name
	replication := &model.ReplicationPolicy{
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
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{replication}}, nil)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replication.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)
	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy,
		filters,
		trigger, destNamespace, description, name)
	assert.NoError(t, err)

	_, err = cl.GetReplicationPolicyByID(ctx, r.ID)
	assert.NoError(t, err)
}

func TestRESTClient_GetReplicationPolicy_EmptyName(t *testing.T) {
	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplicationPolicy(context.Background(), "")

	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_GetReplicationPolicy_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplicationPolicy(ctx, repl.Name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplicationPolicy(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{repl}}, nil)
	p.On("DeleteReplicationPoliciesID", &products.DeleteReplicationPoliciesIDParams{
		ID:      repl.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.DeleteReplicationPoliciesIDOK{}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.DeleteReplicationPolicy(ctx, repl)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_DeleteReplicationPolicy_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()
	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.DeleteReplicationPolicy(ctx, repl)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplicationPolicy_NilParam(t *testing.T) {
	ctx := context.Background()

	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.DeleteReplicationPolicy(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_UpdateReplicationPolicy(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{repl}}, nil)

	p.On("PutReplicationPoliciesID", &products.PutReplicationPoliciesIDParams{
		ID:      repl.ID,
		Policy:  repl,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PutReplicationPoliciesIDOK{}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	err := cl.UpdateReplicationPolicy(ctx, repl)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_UpdateReplicationPolicy_NilParam(t *testing.T) {
	ctx := context.Background()

	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.UpdateReplicationPolicy(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_UpdateReplicationPolicy_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: ns,
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: name,
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.UpdateReplicationPolicy(ctx, repl)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutions_ReplicationIDNotFound(t *testing.T) {
	replExec := &model.ReplicationExecution{
		ID:       1,
		PolicyID: 1,
	}

	ctx := context.Background()

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      1,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)

	_, err := cl.GetReplicationExecutions(ctx, replExec)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationExecutionReplicationIDNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_TriggerReplicationExecution_ReplicationIDNotFound(t *testing.T) {
	replExec := &model.ReplicationExecution{
		ID:       1,
		PolicyID: 1,
	}

	ctx := context.Background()

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)

	err := cl.TriggerReplicationExecution(ctx, replExec)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationExecutionReplicationIDNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_TriggerReplicationExecution(t *testing.T) {
	replExec := &model.ReplicationExecution{
		ID:       0,
		PolicyID: 0,
	}

	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true

	var filters []*model.ReplicationFilter

	trigger := &model.ReplicationTrigger{}
	destNamespace := ns
	description := description
	name := name
	returnedReplication := &model.ReplicationPolicy{
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
	}

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{returnedReplication}}, nil)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)

	p.On("PostReplicationExecutions", &products.PostReplicationExecutionsParams{
		Execution: replExec,
		Context:   ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationExecutionsCreated{}, nil)

	r, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	assert.NoError(t, err)
	assert.Equal(t, r, returnedReplication)

	err = cl.TriggerReplicationExecution(ctx, replExec)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutionsByID_ReplicationNotFound(t *testing.T) {
	replExec := &model.ReplicationExecution{
		ID:       1,
		PolicyID: 1,
	}

	ctx := context.Background()

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)

	_, err := cl.GetReplicationPolicyByID(ctx, replExec.ID)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetReplicationExecutionsByID(t *testing.T) {
	replExec := &model.ReplicationExecution{
		ID:       0,
		PolicyID: 0,
	}

	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true

	var filters []*model.ReplicationFilter

	trigger := &model.ReplicationTrigger{}
	destNamespace := ns
	description := description
	name := name
	replication := &model.ReplicationPolicy{
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
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("PostReplicationPolicies", &products.PostReplicationPoliciesParams{
		Policy: &model.ReplicationPolicy{
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
		},
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PostReplicationPoliciesCreated{}, nil)

	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{replication}}, nil)

	p.On("GetReplicationPoliciesID", &products.GetReplicationPoliciesIDParams{
		ID:      replExec.ID,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesIDOK{Payload: &model.ReplicationPolicy{}}, nil)

	_, err := cl.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy,
		filters,
		trigger, destNamespace, description, name)
	assert.NoError(t, err)

	_, err = cl.GetReplicationPolicyByID(ctx, replExec.ID)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestErrReplicationExecutionNotProvided_Error(t *testing.T) {
	var e ErrReplicationExecutionNotProvided

	assert.Equal(t, ErrReplicationExecutionNotProvidedMsg, e.Error())
}

func TestErrReplicationExecutionReplicationIDMismatch_Error(t *testing.T) {
	var e ErrReplicationExecutionReplicationIDMismatch

	assert.Equal(t, ErrReplicationExecutionReplicationIDMismatchMsg, e.Error())
}

func TestErrReplicationExecutionReplicationIDNotFound_Error(t *testing.T) {
	var e ErrReplicationExecutionReplicationIDNotFound

	assert.Equal(t, ErrReplicationExecutionReplicationIDNotFoundMsg, e.Error())
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
