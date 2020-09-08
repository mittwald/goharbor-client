// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/apiv2/model"
)

// ListProvidersUnderProjectReader is a Reader for the ListProvidersUnderProject structure.
type ListProvidersUnderProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProvidersUnderProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProvidersUnderProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProvidersUnderProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProvidersUnderProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProvidersUnderProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProvidersUnderProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProvidersUnderProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProvidersUnderProjectOK creates a ListProvidersUnderProjectOK with default headers values
func NewListProvidersUnderProjectOK() *ListProvidersUnderProjectOK {
	return &ListProvidersUnderProjectOK{}
}

/*ListProvidersUnderProjectOK handles this case with default header values.

Success
*/
type ListProvidersUnderProjectOK struct {
	Payload []*model.ProviderUnderProject
}

func (o *ListProvidersUnderProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectOK  %+v", 200, o.Payload)
}

func (o *ListProvidersUnderProjectOK) GetPayload() []*model.ProviderUnderProject {
	return o.Payload
}

func (o *ListProvidersUnderProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnderProjectBadRequest creates a ListProvidersUnderProjectBadRequest with default headers values
func NewListProvidersUnderProjectBadRequest() *ListProvidersUnderProjectBadRequest {
	return &ListProvidersUnderProjectBadRequest{}
}

/*ListProvidersUnderProjectBadRequest handles this case with default header values.

Bad request
*/
type ListProvidersUnderProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnderProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ListProvidersUnderProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnderProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnderProjectUnauthorized creates a ListProvidersUnderProjectUnauthorized with default headers values
func NewListProvidersUnderProjectUnauthorized() *ListProvidersUnderProjectUnauthorized {
	return &ListProvidersUnderProjectUnauthorized{}
}

/*ListProvidersUnderProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type ListProvidersUnderProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnderProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProvidersUnderProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnderProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnderProjectForbidden creates a ListProvidersUnderProjectForbidden with default headers values
func NewListProvidersUnderProjectForbidden() *ListProvidersUnderProjectForbidden {
	return &ListProvidersUnderProjectForbidden{}
}

/*ListProvidersUnderProjectForbidden handles this case with default header values.

Forbidden
*/
type ListProvidersUnderProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnderProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectForbidden  %+v", 403, o.Payload)
}

func (o *ListProvidersUnderProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnderProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnderProjectNotFound creates a ListProvidersUnderProjectNotFound with default headers values
func NewListProvidersUnderProjectNotFound() *ListProvidersUnderProjectNotFound {
	return &ListProvidersUnderProjectNotFound{}
}

/*ListProvidersUnderProjectNotFound handles this case with default header values.

Not found
*/
type ListProvidersUnderProjectNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnderProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectNotFound  %+v", 404, o.Payload)
}

func (o *ListProvidersUnderProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnderProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnderProjectInternalServerError creates a ListProvidersUnderProjectInternalServerError with default headers values
func NewListProvidersUnderProjectInternalServerError() *ListProvidersUnderProjectInternalServerError {
	return &ListProvidersUnderProjectInternalServerError{}
}

/*ListProvidersUnderProjectInternalServerError handles this case with default header values.

Internal server error
*/
type ListProvidersUnderProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnderProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/providers][%d] listProvidersUnderProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProvidersUnderProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnderProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
