package configurations

/*func TestAPIGetConfiguration(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	resp, err := c.GetConfigurationsInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestAPIUpdateConfiguration(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	resp, err := c.GetConfigurationsInfo(ctx)
	require.NoError(t, err)

	resp.AuthMode.Value = "oidc_auth"
	authMode, _ := resp.AuthMode.MarshalBinary()
	str := string(authMode)
	updateParams := &model.Configurations{
		AuthMode: &str,
	}
	updateErr := c.UpdateConfigurationsInfo(ctx, updateParams)
	require.NoError(t, updateErr)
}*/
