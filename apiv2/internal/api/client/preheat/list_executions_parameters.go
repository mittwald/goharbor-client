// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// NewListExecutionsParams creates a new ListExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListExecutionsParams() *ListExecutionsParams {
	return &ListExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListExecutionsParamsWithTimeout creates a new ListExecutionsParams object
// with the ability to set a timeout on a request.
func NewListExecutionsParamsWithTimeout(timeout time.Duration) *ListExecutionsParams {
	return &ListExecutionsParams{
		timeout: timeout,
	}
}

// NewListExecutionsParamsWithContext creates a new ListExecutionsParams object
// with the ability to set a context for a request.
func NewListExecutionsParamsWithContext(ctx context.Context) *ListExecutionsParams {
	return &ListExecutionsParams{
		Context: ctx,
	}
}

// NewListExecutionsParamsWithHTTPClient creates a new ListExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListExecutionsParamsWithHTTPClient(client *http.Client) *ListExecutionsParams {
	return &ListExecutionsParams{
		HTTPClient: client,
	}
}

/* ListExecutionsParams contains all the parameters to send to the API endpoint
   for the list executions operation.

   Typically these are written to a http.Request.
*/
type ListExecutionsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* PreheatPolicyName.

	   Preheat Policy Name
	*/
	PreheatPolicyName string

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListExecutionsParams) WithDefaults() *ListExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListExecutionsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListExecutionsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list executions params
func (o *ListExecutionsParams) WithTimeout(timeout time.Duration) *ListExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list executions params
func (o *ListExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list executions params
func (o *ListExecutionsParams) WithContext(ctx context.Context) *ListExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list executions params
func (o *ListExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list executions params
func (o *ListExecutionsParams) WithHTTPClient(client *http.Client) *ListExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list executions params
func (o *ListExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list executions params
func (o *ListExecutionsParams) WithXRequestID(xRequestID *string) *ListExecutionsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list executions params
func (o *ListExecutionsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list executions params
func (o *ListExecutionsParams) WithPage(page *int64) *ListExecutionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list executions params
func (o *ListExecutionsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list executions params
func (o *ListExecutionsParams) WithPageSize(pageSize *int64) *ListExecutionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list executions params
func (o *ListExecutionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPreheatPolicyName adds the preheatPolicyName to the list executions params
func (o *ListExecutionsParams) WithPreheatPolicyName(preheatPolicyName string) *ListExecutionsParams {
	o.SetPreheatPolicyName(preheatPolicyName)
	return o
}

// SetPreheatPolicyName adds the preheatPolicyName to the list executions params
func (o *ListExecutionsParams) SetPreheatPolicyName(preheatPolicyName string) {
	o.PreheatPolicyName = preheatPolicyName
}

// WithProjectName adds the projectName to the list executions params
func (o *ListExecutionsParams) WithProjectName(projectName string) *ListExecutionsParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the list executions params
func (o *ListExecutionsParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithQ adds the q to the list executions params
func (o *ListExecutionsParams) WithQ(q *string) *ListExecutionsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list executions params
func (o *ListExecutionsParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ListExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param preheat_policy_name
	if err := r.SetPathParam("preheat_policy_name", o.PreheatPolicyName); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
