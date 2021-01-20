// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v3/apiv2/model"
)

// GetReportLogReader is a Reader for the GetReportLog structure.
type GetReportLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReportLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportLogOK creates a GetReportLogOK with default headers values
func NewGetReportLogOK() *GetReportLogOK {
	return &GetReportLogOK{}
}

/* GetReportLogOK describes a response with status code 200, with default header values.

Successfully get scan log file
*/
type GetReportLogOK struct {
	Payload string
}

func (o *GetReportLogOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log][%d] getReportLogOK  %+v", 200, o.Payload)
}
func (o *GetReportLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetReportLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportLogUnauthorized creates a GetReportLogUnauthorized with default headers values
func NewGetReportLogUnauthorized() *GetReportLogUnauthorized {
	return &GetReportLogUnauthorized{}
}

/* GetReportLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReportLogUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReportLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log][%d] getReportLogUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReportLogUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReportLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportLogForbidden creates a GetReportLogForbidden with default headers values
func NewGetReportLogForbidden() *GetReportLogForbidden {
	return &GetReportLogForbidden{}
}

/* GetReportLogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReportLogForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReportLogForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log][%d] getReportLogForbidden  %+v", 403, o.Payload)
}
func (o *GetReportLogForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReportLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportLogNotFound creates a GetReportLogNotFound with default headers values
func NewGetReportLogNotFound() *GetReportLogNotFound {
	return &GetReportLogNotFound{}
}

/* GetReportLogNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetReportLogNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReportLogNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log][%d] getReportLogNotFound  %+v", 404, o.Payload)
}
func (o *GetReportLogNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReportLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportLogInternalServerError creates a GetReportLogInternalServerError with default headers values
func NewGetReportLogInternalServerError() *GetReportLogInternalServerError {
	return &GetReportLogInternalServerError{}
}

/* GetReportLogInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetReportLogInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReportLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log][%d] getReportLogInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReportLogInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReportLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
