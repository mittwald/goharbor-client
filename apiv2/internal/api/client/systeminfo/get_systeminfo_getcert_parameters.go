// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

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
)

// NewGetSysteminfoGetcertParams creates a new GetSysteminfoGetcertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSysteminfoGetcertParams() *GetSysteminfoGetcertParams {
	return &GetSysteminfoGetcertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSysteminfoGetcertParamsWithTimeout creates a new GetSysteminfoGetcertParams object
// with the ability to set a timeout on a request.
func NewGetSysteminfoGetcertParamsWithTimeout(timeout time.Duration) *GetSysteminfoGetcertParams {
	return &GetSysteminfoGetcertParams{
		timeout: timeout,
	}
}

// NewGetSysteminfoGetcertParamsWithContext creates a new GetSysteminfoGetcertParams object
// with the ability to set a context for a request.
func NewGetSysteminfoGetcertParamsWithContext(ctx context.Context) *GetSysteminfoGetcertParams {
	return &GetSysteminfoGetcertParams{
		Context: ctx,
	}
}

// NewGetSysteminfoGetcertParamsWithHTTPClient creates a new GetSysteminfoGetcertParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSysteminfoGetcertParamsWithHTTPClient(client *http.Client) *GetSysteminfoGetcertParams {
	return &GetSysteminfoGetcertParams{
		HTTPClient: client,
	}
}

/* GetSysteminfoGetcertParams contains all the parameters to send to the API endpoint
   for the get systeminfo getcert operation.

   Typically these are written to a http.Request.
*/
type GetSysteminfoGetcertParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get systeminfo getcert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSysteminfoGetcertParams) WithDefaults() *GetSysteminfoGetcertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get systeminfo getcert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSysteminfoGetcertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) WithTimeout(timeout time.Duration) *GetSysteminfoGetcertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) WithContext(ctx context.Context) *GetSysteminfoGetcertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) WithHTTPClient(client *http.Client) *GetSysteminfoGetcertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get systeminfo getcert params
func (o *GetSysteminfoGetcertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSysteminfoGetcertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
