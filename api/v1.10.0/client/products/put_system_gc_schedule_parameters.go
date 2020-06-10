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

// NewPutSystemGcScheduleParams creates a new PutSystemGcScheduleParams object
// with the default values initialized.
func NewPutSystemGcScheduleParams() *PutSystemGcScheduleParams {
	var ()
	return &PutSystemGcScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemGcScheduleParamsWithTimeout creates a new PutSystemGcScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSystemGcScheduleParamsWithTimeout(timeout time.Duration) *PutSystemGcScheduleParams {
	var ()
	return &PutSystemGcScheduleParams{

		timeout: timeout,
	}
}

// NewPutSystemGcScheduleParamsWithContext creates a new PutSystemGcScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSystemGcScheduleParamsWithContext(ctx context.Context) *PutSystemGcScheduleParams {
	var ()
	return &PutSystemGcScheduleParams{

		Context: ctx,
	}
}

// NewPutSystemGcScheduleParamsWithHTTPClient creates a new PutSystemGcScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSystemGcScheduleParamsWithHTTPClient(client *http.Client) *PutSystemGcScheduleParams {
	var ()
	return &PutSystemGcScheduleParams{
		HTTPClient: client,
	}
}

/*PutSystemGcScheduleParams contains all the parameters to send to the API endpoint
for the put system gc schedule operation typically these are written to a http.Request
*/
type PutSystemGcScheduleParams struct {

	/*Schedule
	  Updates of gc's schedule.

	*/
	Schedule *model.AdminJobSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put system gc schedule params
func (o *PutSystemGcScheduleParams) WithTimeout(timeout time.Duration) *PutSystemGcScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system gc schedule params
func (o *PutSystemGcScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system gc schedule params
func (o *PutSystemGcScheduleParams) WithContext(ctx context.Context) *PutSystemGcScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system gc schedule params
func (o *PutSystemGcScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system gc schedule params
func (o *PutSystemGcScheduleParams) WithHTTPClient(client *http.Client) *PutSystemGcScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system gc schedule params
func (o *PutSystemGcScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchedule adds the schedule to the put system gc schedule params
func (o *PutSystemGcScheduleParams) WithSchedule(schedule *model.AdminJobSchedule) *PutSystemGcScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the put system gc schedule params
func (o *PutSystemGcScheduleParams) SetSchedule(schedule *model.AdminJobSchedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemGcScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
