// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutQuotasIDReader is a Reader for the PutQuotasID structure.
type PutQuotasIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutQuotasIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutQuotasIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutQuotasIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutQuotasIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutQuotasIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutQuotasIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutQuotasIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutQuotasIDOK creates a PutQuotasIDOK with default headers values
func NewPutQuotasIDOK() *PutQuotasIDOK {
	return &PutQuotasIDOK{}
}

/* PutQuotasIDOK describes a response with status code 200, with default header values.

Updated quota hard limits successfully.
*/
type PutQuotasIDOK struct {
}

func (o *PutQuotasIDOK) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdOK ", 200)
}

func (o *PutQuotasIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutQuotasIDBadRequest creates a PutQuotasIDBadRequest with default headers values
func NewPutQuotasIDBadRequest() *PutQuotasIDBadRequest {
	return &PutQuotasIDBadRequest{}
}

/* PutQuotasIDBadRequest describes a response with status code 400, with default header values.

Illegal format of quota update request.
*/
type PutQuotasIDBadRequest struct {
}

func (o *PutQuotasIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdBadRequest ", 400)
}

func (o *PutQuotasIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutQuotasIDUnauthorized creates a PutQuotasIDUnauthorized with default headers values
func NewPutQuotasIDUnauthorized() *PutQuotasIDUnauthorized {
	return &PutQuotasIDUnauthorized{}
}

/* PutQuotasIDUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type PutQuotasIDUnauthorized struct {
}

func (o *PutQuotasIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdUnauthorized ", 401)
}

func (o *PutQuotasIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutQuotasIDForbidden creates a PutQuotasIDForbidden with default headers values
func NewPutQuotasIDForbidden() *PutQuotasIDForbidden {
	return &PutQuotasIDForbidden{}
}

/* PutQuotasIDForbidden describes a response with status code 403, with default header values.

User does not have permission to the quota.
*/
type PutQuotasIDForbidden struct {
}

func (o *PutQuotasIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdForbidden ", 403)
}

func (o *PutQuotasIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutQuotasIDNotFound creates a PutQuotasIDNotFound with default headers values
func NewPutQuotasIDNotFound() *PutQuotasIDNotFound {
	return &PutQuotasIDNotFound{}
}

/* PutQuotasIDNotFound describes a response with status code 404, with default header values.

Quota ID does not exist.
*/
type PutQuotasIDNotFound struct {
}

func (o *PutQuotasIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdNotFound ", 404)
}

func (o *PutQuotasIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutQuotasIDInternalServerError creates a PutQuotasIDInternalServerError with default headers values
func NewPutQuotasIDInternalServerError() *PutQuotasIDInternalServerError {
	return &PutQuotasIDInternalServerError{}
}

/* PutQuotasIDInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type PutQuotasIDInternalServerError struct {
}

func (o *PutQuotasIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /quotas/{id}][%d] putQuotasIdInternalServerError ", 500)
}

func (o *PutQuotasIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
