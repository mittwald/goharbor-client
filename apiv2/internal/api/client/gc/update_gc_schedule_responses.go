// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
)

// UpdateGCScheduleReader is a Reader for the UpdateGCSchedule structure.
type UpdateGCScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGCScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGCScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGCScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateGCScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGCScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateGCScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGCScheduleOK creates a UpdateGCScheduleOK with default headers values
func NewUpdateGCScheduleOK() *UpdateGCScheduleOK {
	return &UpdateGCScheduleOK{}
}

/*UpdateGCScheduleOK handles this case with default header values.

Updated gc's schedule successfully.
*/
type UpdateGCScheduleOK struct {
}

func (o *UpdateGCScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /system/gc/schedule][%d] updateGcScheduleOK ", 200)
}

func (o *UpdateGCScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGCScheduleBadRequest creates a UpdateGCScheduleBadRequest with default headers values
func NewUpdateGCScheduleBadRequest() *UpdateGCScheduleBadRequest {
	return &UpdateGCScheduleBadRequest{}
}

/*UpdateGCScheduleBadRequest handles this case with default header values.

Bad request
*/
type UpdateGCScheduleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateGCScheduleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system/gc/schedule][%d] updateGcScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGCScheduleBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateGCScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGCScheduleUnauthorized creates a UpdateGCScheduleUnauthorized with default headers values
func NewUpdateGCScheduleUnauthorized() *UpdateGCScheduleUnauthorized {
	return &UpdateGCScheduleUnauthorized{}
}

/*UpdateGCScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateGCScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateGCScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/gc/schedule][%d] updateGcScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateGCScheduleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateGCScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGCScheduleForbidden creates a UpdateGCScheduleForbidden with default headers values
func NewUpdateGCScheduleForbidden() *UpdateGCScheduleForbidden {
	return &UpdateGCScheduleForbidden{}
}

/*UpdateGCScheduleForbidden handles this case with default header values.

Forbidden
*/
type UpdateGCScheduleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateGCScheduleForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/gc/schedule][%d] updateGcScheduleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateGCScheduleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateGCScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGCScheduleInternalServerError creates a UpdateGCScheduleInternalServerError with default headers values
func NewUpdateGCScheduleInternalServerError() *UpdateGCScheduleInternalServerError {
	return &UpdateGCScheduleInternalServerError{}
}

/*UpdateGCScheduleInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateGCScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateGCScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/gc/schedule][%d] updateGcScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateGCScheduleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateGCScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
