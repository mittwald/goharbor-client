// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectsProjectIDReader is a Reader for the DeleteProjectsProjectID structure.
type DeleteProjectsProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectsProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectsProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectsProjectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectsProjectIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectsProjectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteProjectsProjectIDPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectsProjectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectsProjectIDOK creates a DeleteProjectsProjectIDOK with default headers values
func NewDeleteProjectsProjectIDOK() *DeleteProjectsProjectIDOK {
	return &DeleteProjectsProjectIDOK{}
}

/* DeleteProjectsProjectIDOK describes a response with status code 200, with default header values.

Project is deleted successfully.
*/
type DeleteProjectsProjectIDOK struct {
}

func (o *DeleteProjectsProjectIDOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdOK ", 200)
}

func (o *DeleteProjectsProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDBadRequest creates a DeleteProjectsProjectIDBadRequest with default headers values
func NewDeleteProjectsProjectIDBadRequest() *DeleteProjectsProjectIDBadRequest {
	return &DeleteProjectsProjectIDBadRequest{}
}

/* DeleteProjectsProjectIDBadRequest describes a response with status code 400, with default header values.

Invalid project id.
*/
type DeleteProjectsProjectIDBadRequest struct {
}

func (o *DeleteProjectsProjectIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdBadRequest ", 400)
}

func (o *DeleteProjectsProjectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDForbidden creates a DeleteProjectsProjectIDForbidden with default headers values
func NewDeleteProjectsProjectIDForbidden() *DeleteProjectsProjectIDForbidden {
	return &DeleteProjectsProjectIDForbidden{}
}

/* DeleteProjectsProjectIDForbidden describes a response with status code 403, with default header values.

User need to log in first.
*/
type DeleteProjectsProjectIDForbidden struct {
}

func (o *DeleteProjectsProjectIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdForbidden ", 403)
}

func (o *DeleteProjectsProjectIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDNotFound creates a DeleteProjectsProjectIDNotFound with default headers values
func NewDeleteProjectsProjectIDNotFound() *DeleteProjectsProjectIDNotFound {
	return &DeleteProjectsProjectIDNotFound{}
}

/* DeleteProjectsProjectIDNotFound describes a response with status code 404, with default header values.

Project does not exist.
*/
type DeleteProjectsProjectIDNotFound struct {
}

func (o *DeleteProjectsProjectIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdNotFound ", 404)
}

func (o *DeleteProjectsProjectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDPreconditionFailed creates a DeleteProjectsProjectIDPreconditionFailed with default headers values
func NewDeleteProjectsProjectIDPreconditionFailed() *DeleteProjectsProjectIDPreconditionFailed {
	return &DeleteProjectsProjectIDPreconditionFailed{}
}

/* DeleteProjectsProjectIDPreconditionFailed describes a response with status code 412, with default header values.

Project contains policies, can not be deleted.
*/
type DeleteProjectsProjectIDPreconditionFailed struct {
}

func (o *DeleteProjectsProjectIDPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdPreconditionFailed ", 412)
}

func (o *DeleteProjectsProjectIDPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectsProjectIDInternalServerError creates a DeleteProjectsProjectIDInternalServerError with default headers values
func NewDeleteProjectsProjectIDInternalServerError() *DeleteProjectsProjectIDInternalServerError {
	return &DeleteProjectsProjectIDInternalServerError{}
}

/* DeleteProjectsProjectIDInternalServerError describes a response with status code 500, with default header values.

Internal errors.
*/
type DeleteProjectsProjectIDInternalServerError struct {
}

func (o *DeleteProjectsProjectIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectsProjectIdInternalServerError ", 500)
}

func (o *DeleteProjectsProjectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
