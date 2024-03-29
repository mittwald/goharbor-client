// Code generated by go-swagger; DO NOT EDIT.

package schedule

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

// NewListSchedulesParams creates a new ListSchedulesParams object
// with the default values initialized.
func NewListSchedulesParams() *ListSchedulesParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)
	return &ListSchedulesParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListSchedulesParamsWithTimeout creates a new ListSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSchedulesParamsWithTimeout(timeout time.Duration) *ListSchedulesParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)
	return &ListSchedulesParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListSchedulesParamsWithContext creates a new ListSchedulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSchedulesParamsWithContext(ctx context.Context) *ListSchedulesParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)
	return &ListSchedulesParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewListSchedulesParamsWithHTTPClient creates a new ListSchedulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSchedulesParamsWithHTTPClient(client *http.Client) *ListSchedulesParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)
	return &ListSchedulesParams{
		Page:       &pageDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*ListSchedulesParams contains all the parameters to send to the API endpoint
for the list schedules operation typically these are written to a http.Request
*/
type ListSchedulesParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Page
	  The page number

	*/
	Page *int64
	/*PageSize
	  The size of per page

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list schedules params
func (o *ListSchedulesParams) WithTimeout(timeout time.Duration) *ListSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list schedules params
func (o *ListSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list schedules params
func (o *ListSchedulesParams) WithContext(ctx context.Context) *ListSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list schedules params
func (o *ListSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list schedules params
func (o *ListSchedulesParams) WithHTTPClient(client *http.Client) *ListSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list schedules params
func (o *ListSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list schedules params
func (o *ListSchedulesParams) WithXRequestID(xRequestID *string) *ListSchedulesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list schedules params
func (o *ListSchedulesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list schedules params
func (o *ListSchedulesParams) WithPage(page *int64) *ListSchedulesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list schedules params
func (o *ListSchedulesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list schedules params
func (o *ListSchedulesParams) WithPageSize(pageSize *int64) *ListSchedulesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list schedules params
func (o *ListSchedulesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
