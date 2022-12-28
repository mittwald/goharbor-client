// Code generated by go-swagger; DO NOT EDIT.

package jobservice

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

// NewListJobQueuesParams creates a new ListJobQueuesParams object
// with the default values initialized.
func NewListJobQueuesParams() *ListJobQueuesParams {
	var ()
	return &ListJobQueuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListJobQueuesParamsWithTimeout creates a new ListJobQueuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListJobQueuesParamsWithTimeout(timeout time.Duration) *ListJobQueuesParams {
	var ()
	return &ListJobQueuesParams{

		timeout: timeout,
	}
}

// NewListJobQueuesParamsWithContext creates a new ListJobQueuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListJobQueuesParamsWithContext(ctx context.Context) *ListJobQueuesParams {
	var ()
	return &ListJobQueuesParams{

		Context: ctx,
	}
}

// NewListJobQueuesParamsWithHTTPClient creates a new ListJobQueuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListJobQueuesParamsWithHTTPClient(client *http.Client) *ListJobQueuesParams {
	var ()
	return &ListJobQueuesParams{
		HTTPClient: client,
	}
}

/*ListJobQueuesParams contains all the parameters to send to the API endpoint
for the list job queues operation typically these are written to a http.Request
*/
type ListJobQueuesParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list job queues params
func (o *ListJobQueuesParams) WithTimeout(timeout time.Duration) *ListJobQueuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list job queues params
func (o *ListJobQueuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list job queues params
func (o *ListJobQueuesParams) WithContext(ctx context.Context) *ListJobQueuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list job queues params
func (o *ListJobQueuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list job queues params
func (o *ListJobQueuesParams) WithHTTPClient(client *http.Client) *ListJobQueuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list job queues params
func (o *ListJobQueuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list job queues params
func (o *ListJobQueuesParams) WithXRequestID(xRequestID *string) *ListJobQueuesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list job queues params
func (o *ListJobQueuesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *ListJobQueuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}