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

	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// NewPostRegistriesParams creates a new PostRegistriesParams object
// with the default values initialized.
func NewPostRegistriesParams() *PostRegistriesParams {
	var ()
	return &PostRegistriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRegistriesParamsWithTimeout creates a new PostRegistriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRegistriesParamsWithTimeout(timeout time.Duration) *PostRegistriesParams {
	var ()
	return &PostRegistriesParams{

		timeout: timeout,
	}
}

// NewPostRegistriesParamsWithContext creates a new PostRegistriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRegistriesParamsWithContext(ctx context.Context) *PostRegistriesParams {
	var ()
	return &PostRegistriesParams{

		Context: ctx,
	}
}

// NewPostRegistriesParamsWithHTTPClient creates a new PostRegistriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRegistriesParamsWithHTTPClient(client *http.Client) *PostRegistriesParams {
	var ()
	return &PostRegistriesParams{
		HTTPClient: client,
	}
}

/*PostRegistriesParams contains all the parameters to send to the API endpoint
for the post registries operation typically these are written to a http.Request
*/
type PostRegistriesParams struct {

	/*Registry
	  New created registry.

	*/
	Registry *model.Registry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post registries params
func (o *PostRegistriesParams) WithTimeout(timeout time.Duration) *PostRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post registries params
func (o *PostRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post registries params
func (o *PostRegistriesParams) WithContext(ctx context.Context) *PostRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post registries params
func (o *PostRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post registries params
func (o *PostRegistriesParams) WithHTTPClient(client *http.Client) *PostRegistriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post registries params
func (o *PostRegistriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistry adds the registry to the post registries params
func (o *PostRegistriesParams) WithRegistry(registry *model.Registry) *PostRegistriesParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the post registries params
func (o *PostRegistriesParams) SetRegistry(registry *model.Registry) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *PostRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Registry != nil {
		if err := r.SetBodyParam(o.Registry); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
