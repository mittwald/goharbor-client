// Code generated by go-swagger; DO NOT EDIT.

package scanner

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

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// NewSetScannerAsDefaultParams creates a new SetScannerAsDefaultParams object
// with the default values initialized.
func NewSetScannerAsDefaultParams() *SetScannerAsDefaultParams {
	var ()
	return &SetScannerAsDefaultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetScannerAsDefaultParamsWithTimeout creates a new SetScannerAsDefaultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetScannerAsDefaultParamsWithTimeout(timeout time.Duration) *SetScannerAsDefaultParams {
	var ()
	return &SetScannerAsDefaultParams{

		timeout: timeout,
	}
}

// NewSetScannerAsDefaultParamsWithContext creates a new SetScannerAsDefaultParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetScannerAsDefaultParamsWithContext(ctx context.Context) *SetScannerAsDefaultParams {
	var ()
	return &SetScannerAsDefaultParams{

		Context: ctx,
	}
}

// NewSetScannerAsDefaultParamsWithHTTPClient creates a new SetScannerAsDefaultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetScannerAsDefaultParamsWithHTTPClient(client *http.Client) *SetScannerAsDefaultParams {
	var ()
	return &SetScannerAsDefaultParams{
		HTTPClient: client,
	}
}

/*SetScannerAsDefaultParams contains all the parameters to send to the API endpoint
for the set scanner as default operation typically these are written to a http.Request
*/
type SetScannerAsDefaultParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Payload*/
	Payload *model.IsDefault
	/*RegistrationID
	  The scanner registration identifier.

	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithTimeout(timeout time.Duration) *SetScannerAsDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithContext(ctx context.Context) *SetScannerAsDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithHTTPClient(client *http.Client) *SetScannerAsDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithXRequestID(xRequestID *string) *SetScannerAsDefaultParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithPayload(payload *model.IsDefault) *SetScannerAsDefaultParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetPayload(payload *model.IsDefault) {
	o.Payload = payload
}

// WithRegistrationID adds the registrationID to the set scanner as default params
func (o *SetScannerAsDefaultParams) WithRegistrationID(registrationID string) *SetScannerAsDefaultParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the set scanner as default params
func (o *SetScannerAsDefaultParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *SetScannerAsDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}