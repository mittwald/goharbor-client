// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
)

// ListRegistryProviderInfosReader is a Reader for the ListRegistryProviderInfos structure.
type ListRegistryProviderInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistryProviderInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistryProviderInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRegistryProviderInfosUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRegistryProviderInfosForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRegistryProviderInfosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRegistryProviderInfosOK creates a ListRegistryProviderInfosOK with default headers values
func NewListRegistryProviderInfosOK() *ListRegistryProviderInfosOK {
	return &ListRegistryProviderInfosOK{}
}

/*ListRegistryProviderInfosOK handles this case with default header values.

Success.
*/
type ListRegistryProviderInfosOK struct {
	Payload map[string]model.RegistryProviderInfo
}

func (o *ListRegistryProviderInfosOK) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderInfosOK) GetPayload() map[string]model.RegistryProviderInfo {
	return o.Payload
}

func (o *ListRegistryProviderInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderInfosUnauthorized creates a ListRegistryProviderInfosUnauthorized with default headers values
func NewListRegistryProviderInfosUnauthorized() *ListRegistryProviderInfosUnauthorized {
	return &ListRegistryProviderInfosUnauthorized{}
}

/*ListRegistryProviderInfosUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRegistryProviderInfosUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderInfosUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderInfosUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderInfosForbidden creates a ListRegistryProviderInfosForbidden with default headers values
func NewListRegistryProviderInfosForbidden() *ListRegistryProviderInfosForbidden {
	return &ListRegistryProviderInfosForbidden{}
}

/*ListRegistryProviderInfosForbidden handles this case with default header values.

Forbidden
*/
type ListRegistryProviderInfosForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderInfosForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderInfosForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderInfosInternalServerError creates a ListRegistryProviderInfosInternalServerError with default headers values
func NewListRegistryProviderInfosInternalServerError() *ListRegistryProviderInfosInternalServerError {
	return &ListRegistryProviderInfosInternalServerError{}
}

/*ListRegistryProviderInfosInternalServerError handles this case with default header values.

Internal server error
*/
type ListRegistryProviderInfosInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderInfosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderInfosInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
