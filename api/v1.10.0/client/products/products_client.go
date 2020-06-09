// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteProjectsProjectID(params *DeleteProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectsProjectIDOK, error)

	DeleteRegistriesID(params *DeleteRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRegistriesIDOK, error)

	DeleteReplicationPoliciesID(params *DeleteReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReplicationPoliciesIDOK, error)

	DeleteUsersUserID(params *DeleteUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersUserIDOK, error)

	GetHealth(params *GetHealthParams, authInfo runtime.ClientAuthInfoWriter) (*GetHealthOK, error)

	GetProjects(params *GetProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsOK, error)

	GetProjectsProjectID(params *GetProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsProjectIDOK, error)

	GetRegistries(params *GetRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetRegistriesOK, error)

	GetReplicationPolicies(params *GetReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetReplicationPoliciesOK, error)

	GetReplicationPoliciesID(params *GetReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetReplicationPoliciesIDOK, error)

	GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error)

	PostProjects(params *PostProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*PostProjectsCreated, error)

	PostRegistries(params *PostRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*PostRegistriesCreated, error)

	PostReplicationPolicies(params *PostReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*PostReplicationPoliciesCreated, error)

	PostUsers(params *PostUsersParams, authInfo runtime.ClientAuthInfoWriter) (*PostUsersCreated, error)

	PutProjectsProjectID(params *PutProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutProjectsProjectIDOK, error)

	PutRegistriesID(params *PutRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutRegistriesIDOK, error)

	PutReplicationPoliciesID(params *PutReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutReplicationPoliciesIDOK, error)

	PutUsersUserID(params *PutUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutUsersUserIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteProjectsProjectID deletes project by project ID

  This endpoint is aimed to delete project by project ID.

*/
func (a *Client) DeleteProjectsProjectID(params *DeleteProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectsProjectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectsProjectIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteProjectsProjectID",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProjectsProjectIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteProjectsProjectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteProjectsProjectID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRegistriesID deletes specific registry

  This endpoint is for to delete specific registry.

*/
func (a *Client) DeleteRegistriesID(params *DeleteRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRegistriesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRegistriesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRegistriesID",
		Method:             "DELETE",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRegistriesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRegistriesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteRegistriesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteReplicationPoliciesID deletes the replication policy specified by ID

  Delete the replication policy specified by ID.

*/
func (a *Client) DeleteReplicationPoliciesID(params *DeleteReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReplicationPoliciesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReplicationPoliciesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteReplicationPoliciesID",
		Method:             "DELETE",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReplicationPoliciesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteReplicationPoliciesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteReplicationPoliciesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersUserID marks a registered user as be removed

  This endpoint let administrator of Harbor mark a registered user as
be removed.It actually won't be deleted from DB.

*/
func (a *Client) DeleteUsersUserID(params *DeleteUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUserID",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUsersUserIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetHealth healths check API

  The endpoint returns the health stauts of the system.

*/
func (a *Client) GetHealth(params *GetHealthParams, authInfo runtime.ClientAuthInfoWriter) (*GetHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHealth",
		Method:             "GET",
		PathPattern:        "/health",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHealthReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProjects lists projects

  This endpoint returns all projects created by Harbor, and can be filtered by project name.

*/
func (a *Client) GetProjects(params *GetProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProjects",
		Method:             "GET",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProjectsProjectID returns specific project detail information

  This endpoint returns specific project information by project ID.

*/
func (a *Client) GetProjectsProjectID(params *GetProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectsProjectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsProjectIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProjectsProjectID",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectsProjectIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectsProjectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProjectsProjectID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRegistries lists registries

  This endpoint let user list filtered registries by name, if name is nil, list returns all registries.

*/
func (a *Client) GetRegistries(params *GetRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetRegistriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegistriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRegistries",
		Method:             "GET",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegistriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRegistriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRegistries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReplicationPolicies lists replication policies

  This endpoint let user list replication policies

*/
func (a *Client) GetReplicationPolicies(params *GetReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetReplicationPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReplicationPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReplicationPolicies",
		Method:             "GET",
		PathPattern:        "/replication/policies",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReplicationPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReplicationPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetReplicationPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReplicationPoliciesID gets replication policy

  This endpoint let user get replication policy by specific ID.

*/
func (a *Client) GetReplicationPoliciesID(params *GetReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetReplicationPoliciesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReplicationPoliciesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReplicationPoliciesID",
		Method:             "GET",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetReplicationPoliciesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReplicationPoliciesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetReplicationPoliciesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsers gets registered users of harbor

  This endpoint is for user to search registered users, support for filtering results with username.Notice, by now this operation is only for administrator.

*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostProjects creates a new project

  This endpoint is for user to create a new project.

*/
func (a *Client) PostProjects(params *PostProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*PostProjectsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostProjects",
		Method:             "POST",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostProjectsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostRegistries creates a new registry

  This endpoint is for user to create a new registry.

*/
func (a *Client) PostRegistries(params *PostRegistriesParams, authInfo runtime.ClientAuthInfoWriter) (*PostRegistriesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRegistriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRegistries",
		Method:             "POST",
		PathPattern:        "/registries",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRegistriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostRegistriesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostRegistries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostReplicationPolicies creates a replication policy

  This endpoint let user create a replication policy

*/
func (a *Client) PostReplicationPolicies(params *PostReplicationPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*PostReplicationPoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReplicationPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostReplicationPolicies",
		Method:             "POST",
		PathPattern:        "/replication/policies",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostReplicationPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostReplicationPoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostReplicationPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsers creates a new user account

  This endpoint is to create a user if the user does not already exist.

*/
func (a *Client) PostUsers(params *PostUsersParams, authInfo runtime.ClientAuthInfoWriter) (*PostUsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsers",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutProjectsProjectID updates properties for a selected project

  This endpoint is aimed to update the properties of a project.

*/
func (a *Client) PutProjectsProjectID(params *PutProjectsProjectIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutProjectsProjectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutProjectsProjectIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutProjectsProjectID",
		Method:             "PUT",
		PathPattern:        "/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutProjectsProjectIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutProjectsProjectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutProjectsProjectID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutRegistriesID updates a given registry

  This endpoint is for update a given registry.

*/
func (a *Client) PutRegistriesID(params *PutRegistriesIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutRegistriesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRegistriesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRegistriesID",
		Method:             "PUT",
		PathPattern:        "/registries/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutRegistriesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutRegistriesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutRegistriesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutReplicationPoliciesID updates the replication policy

  This endpoint let user update policy.

*/
func (a *Client) PutReplicationPoliciesID(params *PutReplicationPoliciesIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutReplicationPoliciesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutReplicationPoliciesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutReplicationPoliciesID",
		Method:             "PUT",
		PathPattern:        "/replication/policies/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutReplicationPoliciesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutReplicationPoliciesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutReplicationPoliciesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutUsersUserID updates a registered user to change his profile

  This endpoint let a registered user change his profile.

*/
func (a *Client) PutUsersUserID(params *PutUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutUsersUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutUsersUserID",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutUsersUserIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUsersUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutUsersUserID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
