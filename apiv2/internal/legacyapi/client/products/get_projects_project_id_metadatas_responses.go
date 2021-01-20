// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// GetProjectsProjectIDMetadatasReader is a Reader for the GetProjectsProjectIDMetadatas structure.
type GetProjectsProjectIDMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectsProjectIDMetadatasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDMetadatasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDMetadatasOK creates a GetProjectsProjectIDMetadatasOK with default headers values
func NewGetProjectsProjectIDMetadatasOK() *GetProjectsProjectIDMetadatasOK {
	return &GetProjectsProjectIDMetadatasOK{}
}

/* GetProjectsProjectIDMetadatasOK describes a response with status code 200, with default header values.

Get metadata successfully.
*/
type GetProjectsProjectIDMetadatasOK struct {
	Payload *legacy.ProjectMetadata
}

func (o *GetProjectsProjectIDMetadatasOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas][%d] getProjectsProjectIdMetadatasOK  %+v", 200, o.Payload)
}
func (o *GetProjectsProjectIDMetadatasOK) GetPayload() *legacy.ProjectMetadata {
	return o.Payload
}

func (o *GetProjectsProjectIDMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.ProjectMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDMetadatasUnauthorized creates a GetProjectsProjectIDMetadatasUnauthorized with default headers values
func NewGetProjectsProjectIDMetadatasUnauthorized() *GetProjectsProjectIDMetadatasUnauthorized {
	return &GetProjectsProjectIDMetadatasUnauthorized{}
}

/* GetProjectsProjectIDMetadatasUnauthorized describes a response with status code 401, with default header values.

User need to login first.
*/
type GetProjectsProjectIDMetadatasUnauthorized struct {
}

func (o *GetProjectsProjectIDMetadatasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas][%d] getProjectsProjectIdMetadatasUnauthorized ", 401)
}

func (o *GetProjectsProjectIDMetadatasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMetadatasInternalServerError creates a GetProjectsProjectIDMetadatasInternalServerError with default headers values
func NewGetProjectsProjectIDMetadatasInternalServerError() *GetProjectsProjectIDMetadatasInternalServerError {
	return &GetProjectsProjectIDMetadatasInternalServerError{}
}

/* GetProjectsProjectIDMetadatasInternalServerError describes a response with status code 500, with default header values.

Internal server errors.
*/
type GetProjectsProjectIDMetadatasInternalServerError struct {
}

func (o *GetProjectsProjectIDMetadatasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/metadatas][%d] getProjectsProjectIdMetadatasInternalServerError ", 500)
}

func (o *GetProjectsProjectIDMetadatasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
