// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutReplicationPoliciesIDReader is a Reader for the PutReplicationPoliciesID structure.
type PutReplicationPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutReplicationPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutReplicationPoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutReplicationPoliciesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutReplicationPoliciesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutReplicationPoliciesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutReplicationPoliciesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutReplicationPoliciesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutReplicationPoliciesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutReplicationPoliciesIDOK creates a PutReplicationPoliciesIDOK with default headers values
func NewPutReplicationPoliciesIDOK() *PutReplicationPoliciesIDOK {
	return &PutReplicationPoliciesIDOK{}
}

/*PutReplicationPoliciesIDOK handles this case with default header values.

Success
*/
type PutReplicationPoliciesIDOK struct {
}

func (o *PutReplicationPoliciesIDOK) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdOK ", 200)
}

func (o *PutReplicationPoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDBadRequest creates a PutReplicationPoliciesIDBadRequest with default headers values
func NewPutReplicationPoliciesIDBadRequest() *PutReplicationPoliciesIDBadRequest {
	return &PutReplicationPoliciesIDBadRequest{}
}

/*PutReplicationPoliciesIDBadRequest handles this case with default header values.

Bad Request
*/
type PutReplicationPoliciesIDBadRequest struct {
}

func (o *PutReplicationPoliciesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdBadRequest ", 400)
}

func (o *PutReplicationPoliciesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDUnauthorized creates a PutReplicationPoliciesIDUnauthorized with default headers values
func NewPutReplicationPoliciesIDUnauthorized() *PutReplicationPoliciesIDUnauthorized {
	return &PutReplicationPoliciesIDUnauthorized{}
}

/*PutReplicationPoliciesIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PutReplicationPoliciesIDUnauthorized struct {
}

func (o *PutReplicationPoliciesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdUnauthorized ", 401)
}

func (o *PutReplicationPoliciesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDForbidden creates a PutReplicationPoliciesIDForbidden with default headers values
func NewPutReplicationPoliciesIDForbidden() *PutReplicationPoliciesIDForbidden {
	return &PutReplicationPoliciesIDForbidden{}
}

/*PutReplicationPoliciesIDForbidden handles this case with default header values.

Forbidden
*/
type PutReplicationPoliciesIDForbidden struct {
}

func (o *PutReplicationPoliciesIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdForbidden ", 403)
}

func (o *PutReplicationPoliciesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDNotFound creates a PutReplicationPoliciesIDNotFound with default headers values
func NewPutReplicationPoliciesIDNotFound() *PutReplicationPoliciesIDNotFound {
	return &PutReplicationPoliciesIDNotFound{}
}

/*PutReplicationPoliciesIDNotFound handles this case with default header values.

Not Found
*/
type PutReplicationPoliciesIDNotFound struct {
}

func (o *PutReplicationPoliciesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdNotFound ", 404)
}

func (o *PutReplicationPoliciesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDConflict creates a PutReplicationPoliciesIDConflict with default headers values
func NewPutReplicationPoliciesIDConflict() *PutReplicationPoliciesIDConflict {
	return &PutReplicationPoliciesIDConflict{}
}

/*PutReplicationPoliciesIDConflict handles this case with default header values.

Conflict
*/
type PutReplicationPoliciesIDConflict struct {
}

func (o *PutReplicationPoliciesIDConflict) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdConflict ", 409)
}

func (o *PutReplicationPoliciesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutReplicationPoliciesIDInternalServerError creates a PutReplicationPoliciesIDInternalServerError with default headers values
func NewPutReplicationPoliciesIDInternalServerError() *PutReplicationPoliciesIDInternalServerError {
	return &PutReplicationPoliciesIDInternalServerError{}
}

/*PutReplicationPoliciesIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutReplicationPoliciesIDInternalServerError struct {
}

func (o *PutReplicationPoliciesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] putReplicationPoliciesIdInternalServerError ", 500)
}

func (o *PutReplicationPoliciesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
