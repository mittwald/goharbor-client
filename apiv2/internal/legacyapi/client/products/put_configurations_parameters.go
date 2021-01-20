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

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// NewPutConfigurationsParams creates a new PutConfigurationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutConfigurationsParams() *PutConfigurationsParams {
	return &PutConfigurationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutConfigurationsParamsWithTimeout creates a new PutConfigurationsParams object
// with the ability to set a timeout on a request.
func NewPutConfigurationsParamsWithTimeout(timeout time.Duration) *PutConfigurationsParams {
	return &PutConfigurationsParams{
		timeout: timeout,
	}
}

// NewPutConfigurationsParamsWithContext creates a new PutConfigurationsParams object
// with the ability to set a context for a request.
func NewPutConfigurationsParamsWithContext(ctx context.Context) *PutConfigurationsParams {
	return &PutConfigurationsParams{
		Context: ctx,
	}
}

// NewPutConfigurationsParamsWithHTTPClient creates a new PutConfigurationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutConfigurationsParamsWithHTTPClient(client *http.Client) *PutConfigurationsParams {
	return &PutConfigurationsParams{
		HTTPClient: client,
	}
}

/* PutConfigurationsParams contains all the parameters to send to the API endpoint
   for the put configurations operation.

   Typically these are written to a http.Request.
*/
type PutConfigurationsParams struct {

	/* Configurations.

	   The configuration map can contain a subset of the attributes of the schema, which are to be updated.
	*/
	Configurations *legacy.Configurations

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put configurations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConfigurationsParams) WithDefaults() *PutConfigurationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put configurations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConfigurationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put configurations params
func (o *PutConfigurationsParams) WithTimeout(timeout time.Duration) *PutConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put configurations params
func (o *PutConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put configurations params
func (o *PutConfigurationsParams) WithContext(ctx context.Context) *PutConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put configurations params
func (o *PutConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put configurations params
func (o *PutConfigurationsParams) WithHTTPClient(client *http.Client) *PutConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put configurations params
func (o *PutConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurations adds the configurations to the put configurations params
func (o *PutConfigurationsParams) WithConfigurations(configurations *legacy.Configurations) *PutConfigurationsParams {
	o.SetConfigurations(configurations)
	return o
}

// SetConfigurations adds the configurations to the put configurations params
func (o *PutConfigurationsParams) SetConfigurations(configurations *legacy.Configurations) {
	o.Configurations = configurations
}

// WriteToRequest writes these params to a swagger request
func (o *PutConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Configurations != nil {
		if err := r.SetBodyParam(o.Configurations); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
