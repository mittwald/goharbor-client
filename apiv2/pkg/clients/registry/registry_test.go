//go:build !integration

package registry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/registry"
	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	unittesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/testwill/goharbor-client/v5/apiv2/mocks"

	runtimeclient "github.com/go-openapi/runtime/client"
)

const name string = "example-registry"

var (
	ctx      = context.Background()
	authInfo = runtimeclient.BasicAuth("foo", "bar")
	reg      = &modelv2.Registry{
		Credential: &modelv2.RegistryCredential{
			AccessKey:    "",
			AccessSecret: "",
			Type:         "",
		},
		ID:       10,
		Insecure: false,
		Name:     name,
		Type:     "harbor",
		URL:      "http://foo.bar",
	}
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_NewRegistry(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &registry.CreateRegistryParams{
		Registry: reg,
		Context:  ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Registry.On("CreateRegistry", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&registry.CreateRegistryCreated{}, nil)

	err := apiClient.NewRegistry(ctx, reg)

	require.NoError(t, err)

	mockClient.Registry.AssertExpectations(t)
}

func TestRESTClient_GetRegistryByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &registry.GetRegistryParams{
		ID:      reg.ID,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Registry.On("GetRegistry", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&registry.GetRegistryOK{Payload: reg}, nil)

	r, err := apiClient.GetRegistryByID(ctx, reg.ID)
	require.NoError(t, err)
	require.NotNil(t, r)

	mockClient.Registry.AssertExpectations(t)
}

func TestRESTClient_GetRegistryByName(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	name := "name=" + reg.Name

	listParams := &registry.ListRegistriesParams{
		PageSize: &apiClient.Options.PageSize,
		Q:        &name,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Registry.On("ListRegistries", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&registry.ListRegistriesOK{Payload: []*modelv2.Registry{reg}}, nil)

	r, err := apiClient.GetRegistryByName(ctx, reg.Name)
	require.NoError(t, err)
	require.NotNil(t, r)

	mockClient.Registry.AssertExpectations(t)
}

func TestRESTClient_UpdateRegistry(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &registry.UpdateRegistryParams{
		ID:       reg.ID,
		Registry: &modelv2.RegistryUpdate{},
		Context:  ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Registry.On("UpdateRegistry", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&registry.UpdateRegistryOK{}, nil)

	err := apiClient.UpdateRegistry(ctx, &modelv2.RegistryUpdate{}, reg.ID)

	require.NoError(t, err)

	mockClient.Registry.AssertExpectations(t)
}

func TestRESTClient_DeleteRegistryByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &registry.DeleteRegistryParams{
		ID:      reg.ID,
		Context: ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Registry.On("DeleteRegistry", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&registry.DeleteRegistryOK{}, nil)

	err := apiClient.DeleteRegistryByID(ctx, reg.ID)
	require.NoError(t, err)

	mockClient.Registry.AssertExpectations(t)
}
