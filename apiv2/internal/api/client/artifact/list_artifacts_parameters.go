// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewListArtifactsParams creates a new ListArtifactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListArtifactsParams() *ListArtifactsParams {
	return &ListArtifactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListArtifactsParamsWithTimeout creates a new ListArtifactsParams object
// with the ability to set a timeout on a request.
func NewListArtifactsParamsWithTimeout(timeout time.Duration) *ListArtifactsParams {
	return &ListArtifactsParams{
		timeout: timeout,
	}
}

// NewListArtifactsParamsWithContext creates a new ListArtifactsParams object
// with the ability to set a context for a request.
func NewListArtifactsParamsWithContext(ctx context.Context) *ListArtifactsParams {
	return &ListArtifactsParams{
		Context: ctx,
	}
}

// NewListArtifactsParamsWithHTTPClient creates a new ListArtifactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListArtifactsParamsWithHTTPClient(client *http.Client) *ListArtifactsParams {
	return &ListArtifactsParams{
		HTTPClient: client,
	}
}

/* ListArtifactsParams contains all the parameters to send to the API endpoint
   for the list artifacts operation.

   Typically these are written to a http.Request.
*/
type ListArtifactsParams struct {

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

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string

	/* WithImmutableStatus.

	   Specify whether the immutable status is included inside the tags of the returning artifacts. Only works when setting "with_tag=true"
	*/
	WithImmutableStatus *bool

	/* WithLabel.

	   Specify whether the labels are included inside the returning artifacts
	*/
	WithLabel *bool

	/* WithScanOverview.

	   Specify whether the scan overview is included inside the returning artifacts
	*/
	WithScanOverview *bool

	/* WithSignature.

	   Specify whether the signature is included inside the tags of the returning artifacts. Only works when setting "with_tag=true"
	*/
	WithSignature *bool

	/* WithTag.

	   Specify whether the tags are included inside the returning artifacts

	   Default: true
	*/
	WithTag *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListArtifactsParams) WithDefaults() *ListArtifactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list artifacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListArtifactsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)

		withImmutableStatusDefault = bool(false)

		withLabelDefault = bool(false)

		withScanOverviewDefault = bool(false)

		withSignatureDefault = bool(false)

		withTagDefault = bool(true)
	)

	val := ListArtifactsParams{
		Page:                &pageDefault,
		PageSize:            &pageSizeDefault,
		WithImmutableStatus: &withImmutableStatusDefault,
		WithLabel:           &withLabelDefault,
		WithScanOverview:    &withScanOverviewDefault,
		WithSignature:       &withSignatureDefault,
		WithTag:             &withTagDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list artifacts params
func (o *ListArtifactsParams) WithTimeout(timeout time.Duration) *ListArtifactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list artifacts params
func (o *ListArtifactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list artifacts params
func (o *ListArtifactsParams) WithContext(ctx context.Context) *ListArtifactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list artifacts params
func (o *ListArtifactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list artifacts params
func (o *ListArtifactsParams) WithHTTPClient(client *http.Client) *ListArtifactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list artifacts params
func (o *ListArtifactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list artifacts params
func (o *ListArtifactsParams) WithXRequestID(xRequestID *string) *ListArtifactsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list artifacts params
func (o *ListArtifactsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list artifacts params
func (o *ListArtifactsParams) WithPage(page *int64) *ListArtifactsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list artifacts params
func (o *ListArtifactsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list artifacts params
func (o *ListArtifactsParams) WithPageSize(pageSize *int64) *ListArtifactsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list artifacts params
func (o *ListArtifactsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectName adds the projectName to the list artifacts params
func (o *ListArtifactsParams) WithProjectName(projectName string) *ListArtifactsParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the list artifacts params
func (o *ListArtifactsParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithQ adds the q to the list artifacts params
func (o *ListArtifactsParams) WithQ(q *string) *ListArtifactsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list artifacts params
func (o *ListArtifactsParams) SetQ(q *string) {
	o.Q = q
}

// WithRepositoryName adds the repositoryName to the list artifacts params
func (o *ListArtifactsParams) WithRepositoryName(repositoryName string) *ListArtifactsParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the list artifacts params
func (o *ListArtifactsParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WithWithImmutableStatus adds the withImmutableStatus to the list artifacts params
func (o *ListArtifactsParams) WithWithImmutableStatus(withImmutableStatus *bool) *ListArtifactsParams {
	o.SetWithImmutableStatus(withImmutableStatus)
	return o
}

// SetWithImmutableStatus adds the withImmutableStatus to the list artifacts params
func (o *ListArtifactsParams) SetWithImmutableStatus(withImmutableStatus *bool) {
	o.WithImmutableStatus = withImmutableStatus
}

// WithWithLabel adds the withLabel to the list artifacts params
func (o *ListArtifactsParams) WithWithLabel(withLabel *bool) *ListArtifactsParams {
	o.SetWithLabel(withLabel)
	return o
}

// SetWithLabel adds the withLabel to the list artifacts params
func (o *ListArtifactsParams) SetWithLabel(withLabel *bool) {
	o.WithLabel = withLabel
}

// WithWithScanOverview adds the withScanOverview to the list artifacts params
func (o *ListArtifactsParams) WithWithScanOverview(withScanOverview *bool) *ListArtifactsParams {
	o.SetWithScanOverview(withScanOverview)
	return o
}

// SetWithScanOverview adds the withScanOverview to the list artifacts params
func (o *ListArtifactsParams) SetWithScanOverview(withScanOverview *bool) {
	o.WithScanOverview = withScanOverview
}

// WithWithSignature adds the withSignature to the list artifacts params
func (o *ListArtifactsParams) WithWithSignature(withSignature *bool) *ListArtifactsParams {
	o.SetWithSignature(withSignature)
	return o
}

// SetWithSignature adds the withSignature to the list artifacts params
func (o *ListArtifactsParams) SetWithSignature(withSignature *bool) {
	o.WithSignature = withSignature
}

// WithWithTag adds the withTag to the list artifacts params
func (o *ListArtifactsParams) WithWithTag(withTag *bool) *ListArtifactsParams {
	o.SetWithTag(withTag)
	return o
}

// SetWithTag adds the withTag to the list artifacts params
func (o *ListArtifactsParams) SetWithTag(withTag *bool) {
	o.WithTag = withTag
}

// WriteToRequest writes these params to a swagger request
func (o *ListArtifactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if o.WithImmutableStatus != nil {

		// query param with_immutable_status
		var qrWithImmutableStatus bool

		if o.WithImmutableStatus != nil {
			qrWithImmutableStatus = *o.WithImmutableStatus
		}
		qWithImmutableStatus := swag.FormatBool(qrWithImmutableStatus)
		if qWithImmutableStatus != "" {

			if err := r.SetQueryParam("with_immutable_status", qWithImmutableStatus); err != nil {
				return err
			}
		}
	}

	if o.WithLabel != nil {

		// query param with_label
		var qrWithLabel bool

		if o.WithLabel != nil {
			qrWithLabel = *o.WithLabel
		}
		qWithLabel := swag.FormatBool(qrWithLabel)
		if qWithLabel != "" {

			if err := r.SetQueryParam("with_label", qWithLabel); err != nil {
				return err
			}
		}
	}

	if o.WithScanOverview != nil {

		// query param with_scan_overview
		var qrWithScanOverview bool

		if o.WithScanOverview != nil {
			qrWithScanOverview = *o.WithScanOverview
		}
		qWithScanOverview := swag.FormatBool(qrWithScanOverview)
		if qWithScanOverview != "" {

			if err := r.SetQueryParam("with_scan_overview", qWithScanOverview); err != nil {
				return err
			}
		}
	}

	if o.WithSignature != nil {

		// query param with_signature
		var qrWithSignature bool

		if o.WithSignature != nil {
			qrWithSignature = *o.WithSignature
		}
		qWithSignature := swag.FormatBool(qrWithSignature)
		if qWithSignature != "" {

			if err := r.SetQueryParam("with_signature", qWithSignature); err != nil {
				return err
			}
		}
	}

	if o.WithTag != nil {

		// query param with_tag
		var qrWithTag bool

		if o.WithTag != nil {
			qrWithTag = *o.WithTag
		}
		qWithTag := swag.FormatBool(qrWithTag)
		if qWithTag != "" {

			if err := r.SetQueryParam("with_tag", qWithTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
