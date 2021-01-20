// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetProjectsProjectIDImmutabletagrulesParams creates a new GetProjectsProjectIDImmutabletagrulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsProjectIDImmutabletagrulesParams() *GetProjectsProjectIDImmutabletagrulesParams {
	return &GetProjectsProjectIDImmutabletagrulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDImmutabletagrulesParamsWithTimeout creates a new GetProjectsProjectIDImmutabletagrulesParams object
// with the ability to set a timeout on a request.
func NewGetProjectsProjectIDImmutabletagrulesParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDImmutabletagrulesParams {
	return &GetProjectsProjectIDImmutabletagrulesParams{
		timeout: timeout,
	}
}

// NewGetProjectsProjectIDImmutabletagrulesParamsWithContext creates a new GetProjectsProjectIDImmutabletagrulesParams object
// with the ability to set a context for a request.
func NewGetProjectsProjectIDImmutabletagrulesParamsWithContext(ctx context.Context) *GetProjectsProjectIDImmutabletagrulesParams {
	return &GetProjectsProjectIDImmutabletagrulesParams{
		Context: ctx,
	}
}

// NewGetProjectsProjectIDImmutabletagrulesParamsWithHTTPClient creates a new GetProjectsProjectIDImmutabletagrulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsProjectIDImmutabletagrulesParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDImmutabletagrulesParams {
	return &GetProjectsProjectIDImmutabletagrulesParams{
		HTTPClient: client,
	}
}

/* GetProjectsProjectIDImmutabletagrulesParams contains all the parameters to send to the API endpoint
   for the get projects project ID immutabletagrules operation.

   Typically these are written to a http.Request.
*/
type GetProjectsProjectIDImmutabletagrulesParams struct {

	/* ProjectID.

	   Relevant project ID.

	   Format: int64
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects project ID immutabletagrules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDImmutabletagrulesParams) WithDefaults() *GetProjectsProjectIDImmutabletagrulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects project ID immutabletagrules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDImmutabletagrulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDImmutabletagrulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) WithContext(ctx context.Context) *GetProjectsProjectIDImmutabletagrulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDImmutabletagrulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) WithProjectID(projectID int64) *GetProjectsProjectIDImmutabletagrulesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID immutabletagrules params
func (o *GetProjectsProjectIDImmutabletagrulesParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDImmutabletagrulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
