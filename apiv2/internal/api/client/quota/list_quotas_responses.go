// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
)

// ListQuotasReader is a Reader for the ListQuotas structure.
type ListQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListQuotasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListQuotasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListQuotasOK creates a ListQuotasOK with default headers values
func NewListQuotasOK() *ListQuotasOK {
	return &ListQuotasOK{}
}

/*ListQuotasOK handles this case with default header values.

Successfully retrieved the quotas.
*/
type ListQuotasOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of access logs
	 */
	XTotalCount int64

	Payload []*model.Quota
}

func (o *ListQuotasOK) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasOK  %+v", 200, o.Payload)
}

func (o *ListQuotasOK) GetPayload() []*model.Quota {
	return o.Payload
}

func (o *ListQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotasUnauthorized creates a ListQuotasUnauthorized with default headers values
func NewListQuotasUnauthorized() *ListQuotasUnauthorized {
	return &ListQuotasUnauthorized{}
}

/*ListQuotasUnauthorized handles this case with default header values.

Unauthorized
*/
type ListQuotasUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListQuotasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasUnauthorized  %+v", 401, o.Payload)
}

func (o *ListQuotasUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListQuotasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotasForbidden creates a ListQuotasForbidden with default headers values
func NewListQuotasForbidden() *ListQuotasForbidden {
	return &ListQuotasForbidden{}
}

/*ListQuotasForbidden handles this case with default header values.

Forbidden
*/
type ListQuotasForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListQuotasForbidden) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasForbidden  %+v", 403, o.Payload)
}

func (o *ListQuotasForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListQuotasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotasInternalServerError creates a ListQuotasInternalServerError with default headers values
func NewListQuotasInternalServerError() *ListQuotasInternalServerError {
	return &ListQuotasInternalServerError{}
}

/*ListQuotasInternalServerError handles this case with default header values.

Internal server error
*/
type ListQuotasInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasInternalServerError  %+v", 500, o.Payload)
}

func (o *ListQuotasInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
