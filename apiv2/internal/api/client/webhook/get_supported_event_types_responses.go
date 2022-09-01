// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
)

// GetSupportedEventTypesReader is a Reader for the GetSupportedEventTypes structure.
type GetSupportedEventTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupportedEventTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupportedEventTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSupportedEventTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSupportedEventTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSupportedEventTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSupportedEventTypesOK creates a GetSupportedEventTypesOK with default headers values
func NewGetSupportedEventTypesOK() *GetSupportedEventTypesOK {
	return &GetSupportedEventTypesOK{}
}

/*GetSupportedEventTypesOK handles this case with default header values.

Success
*/
type GetSupportedEventTypesOK struct {
	Payload *model.SupportedWebhookEventTypes
}

func (o *GetSupportedEventTypesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesOK  %+v", 200, o.Payload)
}

func (o *GetSupportedEventTypesOK) GetPayload() *model.SupportedWebhookEventTypes {
	return o.Payload
}

func (o *GetSupportedEventTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SupportedWebhookEventTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedEventTypesUnauthorized creates a GetSupportedEventTypesUnauthorized with default headers values
func NewGetSupportedEventTypesUnauthorized() *GetSupportedEventTypesUnauthorized {
	return &GetSupportedEventTypesUnauthorized{}
}

/*GetSupportedEventTypesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSupportedEventTypesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSupportedEventTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSupportedEventTypesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedEventTypesForbidden creates a GetSupportedEventTypesForbidden with default headers values
func NewGetSupportedEventTypesForbidden() *GetSupportedEventTypesForbidden {
	return &GetSupportedEventTypesForbidden{}
}

/*GetSupportedEventTypesForbidden handles this case with default header values.

Forbidden
*/
type GetSupportedEventTypesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSupportedEventTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetSupportedEventTypesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedEventTypesInternalServerError creates a GetSupportedEventTypesInternalServerError with default headers values
func NewGetSupportedEventTypesInternalServerError() *GetSupportedEventTypesInternalServerError {
	return &GetSupportedEventTypesInternalServerError{}
}

/*GetSupportedEventTypesInternalServerError handles this case with default header values.

Internal server error
*/
type GetSupportedEventTypesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSupportedEventTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupportedEventTypesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
