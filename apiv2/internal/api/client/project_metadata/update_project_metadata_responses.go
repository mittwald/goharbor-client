// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
)

// UpdateProjectMetadataReader is a Reader for the UpdateProjectMetadata structure.
type UpdateProjectMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateProjectMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectMetadataOK creates a UpdateProjectMetadataOK with default headers values
func NewUpdateProjectMetadataOK() *UpdateProjectMetadataOK {
	return &UpdateProjectMetadataOK{}
}

/*UpdateProjectMetadataOK handles this case with default header values.

Success
*/
type UpdateProjectMetadataOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateProjectMetadataOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataOK ", 200)
}

func (o *UpdateProjectMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateProjectMetadataBadRequest creates a UpdateProjectMetadataBadRequest with default headers values
func NewUpdateProjectMetadataBadRequest() *UpdateProjectMetadataBadRequest {
	return &UpdateProjectMetadataBadRequest{}
}

/*UpdateProjectMetadataBadRequest handles this case with default header values.

Bad request
*/
type UpdateProjectMetadataBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMetadataBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMetadataUnauthorized creates a UpdateProjectMetadataUnauthorized with default headers values
func NewUpdateProjectMetadataUnauthorized() *UpdateProjectMetadataUnauthorized {
	return &UpdateProjectMetadataUnauthorized{}
}

/*UpdateProjectMetadataUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateProjectMetadataUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMetadataUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMetadataForbidden creates a UpdateProjectMetadataForbidden with default headers values
func NewUpdateProjectMetadataForbidden() *UpdateProjectMetadataForbidden {
	return &UpdateProjectMetadataForbidden{}
}

/*UpdateProjectMetadataForbidden handles this case with default header values.

Forbidden
*/
type UpdateProjectMetadataForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMetadataForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMetadataNotFound creates a UpdateProjectMetadataNotFound with default headers values
func NewUpdateProjectMetadataNotFound() *UpdateProjectMetadataNotFound {
	return &UpdateProjectMetadataNotFound{}
}

/*UpdateProjectMetadataNotFound handles this case with default header values.

Not found
*/
type UpdateProjectMetadataNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMetadataNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMetadataConflict creates a UpdateProjectMetadataConflict with default headers values
func NewUpdateProjectMetadataConflict() *UpdateProjectMetadataConflict {
	return &UpdateProjectMetadataConflict{}
}

/*UpdateProjectMetadataConflict handles this case with default header values.

Conflict
*/
type UpdateProjectMetadataConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataConflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectMetadataConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMetadataInternalServerError creates a UpdateProjectMetadataInternalServerError with default headers values
func NewUpdateProjectMetadataInternalServerError() *UpdateProjectMetadataInternalServerError {
	return &UpdateProjectMetadataInternalServerError{}
}

/*UpdateProjectMetadataInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateProjectMetadataInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/metadatas/{meta_name}][%d] updateProjectMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMetadataInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
