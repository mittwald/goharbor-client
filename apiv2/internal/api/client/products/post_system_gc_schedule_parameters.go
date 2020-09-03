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

	"github.com/mittwald/goharbor-client/apiv2/model"
)

// NewPostSystemGcScheduleParams creates a new PostSystemGcScheduleParams object
// with the default values initialized.
func NewPostSystemGcScheduleParams() *PostSystemGcScheduleParams {
	var ()
	return &PostSystemGcScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSystemGcScheduleParamsWithTimeout creates a new PostSystemGcScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSystemGcScheduleParamsWithTimeout(timeout time.Duration) *PostSystemGcScheduleParams {
	var ()
	return &PostSystemGcScheduleParams{

		timeout: timeout,
	}
}

// NewPostSystemGcScheduleParamsWithContext creates a new PostSystemGcScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSystemGcScheduleParamsWithContext(ctx context.Context) *PostSystemGcScheduleParams {
	var ()
	return &PostSystemGcScheduleParams{

		Context: ctx,
	}
}

// NewPostSystemGcScheduleParamsWithHTTPClient creates a new PostSystemGcScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSystemGcScheduleParamsWithHTTPClient(client *http.Client) *PostSystemGcScheduleParams {
	var ()
	return &PostSystemGcScheduleParams{
		HTTPClient: client,
	}
}

/*PostSystemGcScheduleParams contains all the parameters to send to the API endpoint
for the post system gc schedule operation typically these are written to a http.Request
*/
type PostSystemGcScheduleParams struct {

	/*Schedule
	  Updates of gc's schedule.

	*/
	Schedule *model.AdminJobSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post system gc schedule params
func (o *PostSystemGcScheduleParams) WithTimeout(timeout time.Duration) *PostSystemGcScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post system gc schedule params
func (o *PostSystemGcScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post system gc schedule params
func (o *PostSystemGcScheduleParams) WithContext(ctx context.Context) *PostSystemGcScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post system gc schedule params
func (o *PostSystemGcScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post system gc schedule params
func (o *PostSystemGcScheduleParams) WithHTTPClient(client *http.Client) *PostSystemGcScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post system gc schedule params
func (o *PostSystemGcScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchedule adds the schedule to the post system gc schedule params
func (o *PostSystemGcScheduleParams) WithSchedule(schedule *model.AdminJobSchedule) *PostSystemGcScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the post system gc schedule params
func (o *PostSystemGcScheduleParams) SetSchedule(schedule *model.AdminJobSchedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PostSystemGcScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
