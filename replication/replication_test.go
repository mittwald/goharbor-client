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

func TestRESTClient_NewReplication(t *testing.T) {
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{}
	destNamespace := "testnamespace"
	description := "a test replication"
	name := "testreplication"
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

	r, err := cl.NewReplication(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, r, returnedReplication)
}

func TestRESTClient_NewReplication_ErrOnPOST(t *testing.T) {
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{}
	destNamespace := "testnamespace"
	description := "a test replication"
	name := "testreplication"

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

	r, err := cl.NewReplication(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationInternalErrors{}, err)
	}
}

func TestRESTClient_NewReplication_ErrOnGET(t *testing.T) {
	ctx := context.Background()
	destRegistry := &model.Registry{ID: 1, Name: "reg1"}
	srcRegistry := &model.Registry{Name: "reg2"}
	replicateDeletion := true
	override := true
	enablePolicy := true
	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{}
	destNamespace := "testnamespace"
	description := "a test replication"
	name := "testreplication"

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

	r, err := cl.NewReplication(ctx, destRegistry, srcRegistry, replicateDeletion, override, enablePolicy, filters,
		trigger, destNamespace, description, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationUnauthorized{}, err)
	}
}

func TestRESTClient_GetReplication(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{repl}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplication(ctx, repl.Name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, repl, r)
}

func TestRESTClient_GetReplication_EmptyName(t *testing.T) {
	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplication(context.Background(), "")

	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_GetReplication_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.GetReplication(ctx, repl.Name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplication(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
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

	err := cl.DeleteReplication(ctx, repl)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_DeleteReplication_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
	}
	ctx := context.Background()
	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.DeleteReplication(ctx, repl)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestRESTClient_DeleteReplication_NilParam(t *testing.T) {
	ctx := context.Background()

	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.DeleteReplication(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_UpdateReplication(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
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
	err := cl.UpdateReplication(ctx, repl)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_UpdateReplication_NilParam(t *testing.T) {
	ctx := context.Background()

	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.UpdateReplication(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotProvided{}, err)
	}
}

func TestRESTClient_UpdateReplication_NotFound(t *testing.T) {
	repl := &model.ReplicationPolicy{
		Deletion:      true,
		Description:   "a replication policy",
		DestNamespace: "testnamespace",
		DestRegistry: &model.Registry{
			Description: "a test registry",
			ID:          11,
			Name:        "testregistry",
		},
		ID:   1,
		Name: "testreplication",
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetReplicationPolicies", &products.GetReplicationPoliciesParams{
		Name:    &repl.Name,
		Context: ctx,
	}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetReplicationPoliciesOK{Payload: []*model.ReplicationPolicy{}}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	err := cl.UpdateReplication(ctx, repl)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}
