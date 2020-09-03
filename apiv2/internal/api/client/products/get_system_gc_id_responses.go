// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/apiv2/model"
)

// GetSystemGcIDReader is a Reader for the GetSystemGcID structure.
type GetSystemGcIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemGcIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemGcIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemGcIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSystemGcIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystemGcIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemGcIDOK creates a GetSystemGcIDOK with default headers values
func NewGetSystemGcIDOK() *GetSystemGcIDOK {
	return &GetSystemGcIDOK{}
}

/*GetSystemGcIDOK handles this case with default header values.

Get gc results successfully.
*/
type GetSystemGcIDOK struct {
	Payload *model.GCResult
}

func (o *GetSystemGcIDOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/{id}][%d] getSystemGcIdOK  %+v", 200, o.Payload)
}

func (o *GetSystemGcIDOK) GetPayload() *model.GCResult {
	return o.Payload
}

func (o *GetSystemGcIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GCResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemGcIDUnauthorized creates a GetSystemGcIDUnauthorized with default headers values
func NewGetSystemGcIDUnauthorized() *GetSystemGcIDUnauthorized {
	return &GetSystemGcIDUnauthorized{}
}

/*GetSystemGcIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetSystemGcIDUnauthorized struct {
}

func (o *GetSystemGcIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/{id}][%d] getSystemGcIdUnauthorized ", 401)
}

func (o *GetSystemGcIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemGcIDForbidden creates a GetSystemGcIDForbidden with default headers values
func NewGetSystemGcIDForbidden() *GetSystemGcIDForbidden {
	return &GetSystemGcIDForbidden{}
}

/*GetSystemGcIDForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type GetSystemGcIDForbidden struct {
}

func (o *GetSystemGcIDForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/{id}][%d] getSystemGcIdForbidden ", 403)
}

func (o *GetSystemGcIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemGcIDInternalServerError creates a GetSystemGcIDInternalServerError with default headers values
func NewGetSystemGcIDInternalServerError() *GetSystemGcIDInternalServerError {
	return &GetSystemGcIDInternalServerError{}
}

/*GetSystemGcIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetSystemGcIDInternalServerError struct {
}

func (o *GetSystemGcIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/{id}][%d] getSystemGcIdInternalServerError ", 500)
}

func (o *GetSystemGcIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
