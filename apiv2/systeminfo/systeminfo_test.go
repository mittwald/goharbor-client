//go:build !integration
// +build !integration

package systeminfo

//
// var authInfo = runtimeclient.BasicAuth("foo", "bar")
//
// func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *legacyclient.Harbor {
// 	return &legacyclient.Harbor{
// 		Products: service,
// 	}
// }
//
// func TestRESTClient_Health(t *testing.T) {
// 	p := &mocks.MockProductsClientService{}
//
// 	legacyClient := BuildLegacyClientWithMock(p)
//
// 	cl := NewClient(legacyClient, nil, authInfo)
//
// 	ctx := context.Background()
//
// 	healthParams := &products.GetHealthParams{
// 		Context: ctx,
// 	}
//
// 	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&products.GetHealthOK{Payload: &legacymodel.OverallHealthStatus{}}, nil)
//
// 	_, err := cl.Health(ctx)
//
// 	assert.NoError(t, err)
//
// 	p.AssertExpectations(t)
// }
//
// func TestRESTClient_Health_Error(t *testing.T) {
// 	p := &mocks.MockProductsClientService{}
//
// 	legacyClient := BuildLegacyClientWithMock(p)
//
// 	cl := NewClient(legacyClient, nil, authInfo)
//
// 	ctx := context.Background()
//
// 	healthParams := &products.GetHealthParams{
// 		Context: ctx,
// 	}
//
// 	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&products.GetHealthOK{Payload: &legacymodel.OverallHealthStatus{}},
// 			errors.New("err"))
//
// 	_, err := cl.Health(ctx)
//
// 	if assert.Error(t, err) {
// 		assert.Equal(t, errors.New("err"), err)
// 	}
//
// 	p.AssertExpectations(t)
// }
