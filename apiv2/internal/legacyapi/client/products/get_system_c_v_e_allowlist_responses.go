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

// GetSystemCVEAllowlistReader is a Reader for the GetSystemCVEAllowlist structure.
type GetSystemCVEAllowlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemCVEAllowlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemCVEAllowlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemCVEAllowlistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystemCVEAllowlistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemCVEAllowlistOK creates a GetSystemCVEAllowlistOK with default headers values
func NewGetSystemCVEAllowlistOK() *GetSystemCVEAllowlistOK {
	return &GetSystemCVEAllowlistOK{}
}

/* GetSystemCVEAllowlistOK describes a response with status code 200, with default header values.

Successfully retrieved the CVE allowlist.
*/
type GetSystemCVEAllowlistOK struct {
	Payload *legacy.CVEAllowlist
}

func (o *GetSystemCVEAllowlistOK) Error() string {
	return fmt.Sprintf("[GET /system/CVEAllowlist][%d] getSystemCVEAllowlistOK  %+v", 200, o.Payload)
}
func (o *GetSystemCVEAllowlistOK) GetPayload() *legacy.CVEAllowlist {
	return o.Payload
}

func (o *GetSystemCVEAllowlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.CVEAllowlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemCVEAllowlistUnauthorized creates a GetSystemCVEAllowlistUnauthorized with default headers values
func NewGetSystemCVEAllowlistUnauthorized() *GetSystemCVEAllowlistUnauthorized {
	return &GetSystemCVEAllowlistUnauthorized{}
}

/* GetSystemCVEAllowlistUnauthorized describes a response with status code 401, with default header values.

User is not authenticated.
*/
type GetSystemCVEAllowlistUnauthorized struct {
}

func (o *GetSystemCVEAllowlistUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/CVEAllowlist][%d] getSystemCVEAllowlistUnauthorized ", 401)
}

func (o *GetSystemCVEAllowlistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemCVEAllowlistInternalServerError creates a GetSystemCVEAllowlistInternalServerError with default headers values
func NewGetSystemCVEAllowlistInternalServerError() *GetSystemCVEAllowlistInternalServerError {
	return &GetSystemCVEAllowlistInternalServerError{}
}

/* GetSystemCVEAllowlistInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetSystemCVEAllowlistInternalServerError struct {
}

func (o *GetSystemCVEAllowlistInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/CVEAllowlist][%d] getSystemCVEAllowlistInternalServerError ", 500)
}

func (o *GetSystemCVEAllowlistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
