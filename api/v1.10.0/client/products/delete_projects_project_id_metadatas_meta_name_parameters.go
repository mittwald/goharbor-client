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
	"github.com/go-openapi/swag"
)

// NewDeleteProjectsProjectIDMetadatasMetaNameParams creates a new DeleteProjectsProjectIDMetadatasMetaNameParams object
// with the default values initialized.
func NewDeleteProjectsProjectIDMetadatasMetaNameParams() *DeleteProjectsProjectIDMetadatasMetaNameParams {
	var ()
	return &DeleteProjectsProjectIDMetadatasMetaNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithTimeout creates a new DeleteProjectsProjectIDMetadatasMetaNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithTimeout(timeout time.Duration) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	var ()
	return &DeleteProjectsProjectIDMetadatasMetaNameParams{

		timeout: timeout,
	}
}

// NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithContext creates a new DeleteProjectsProjectIDMetadatasMetaNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithContext(ctx context.Context) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	var ()
	return &DeleteProjectsProjectIDMetadatasMetaNameParams{

		Context: ctx,
	}
}

// NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithHTTPClient creates a new DeleteProjectsProjectIDMetadatasMetaNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProjectsProjectIDMetadatasMetaNameParamsWithHTTPClient(client *http.Client) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	var ()
	return &DeleteProjectsProjectIDMetadatasMetaNameParams{
		HTTPClient: client,
	}
}

/*DeleteProjectsProjectIDMetadatasMetaNameParams contains all the parameters to send to the API endpoint
for the delete projects project ID metadatas meta name operation typically these are written to a http.Request
*/
type DeleteProjectsProjectIDMetadatasMetaNameParams struct {

	/*MetaName
	  The name of metadat.

	*/
	MetaName string
	/*ProjectID
	  The ID of project.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WithTimeout(timeout time.Duration) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WithContext(ctx context.Context) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WithHTTPClient(client *http.Client) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetaName adds the metaName to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WithMetaName(metaName string) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	o.SetMetaName(metaName)
	return o
}

// SetMetaName adds the metaName to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) SetMetaName(metaName string) {
	o.MetaName = metaName
}

// WithProjectID adds the projectID to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WithProjectID(projectID int64) *DeleteProjectsProjectIDMetadatasMetaNameParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete projects project ID metadatas meta name params
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectsProjectIDMetadatasMetaNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param meta_name
	if err := r.SetPathParam("meta_name", o.MetaName); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
