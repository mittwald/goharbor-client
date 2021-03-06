// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUsersUserIDReader is a Reader for the DeleteUsersUserID structure.
type DeleteUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUsersUserIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsersUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsersUserIDOK creates a DeleteUsersUserIDOK with default headers values
func NewDeleteUsersUserIDOK() *DeleteUsersUserIDOK {
	return &DeleteUsersUserIDOK{}
}

/* DeleteUsersUserIDOK describes a response with status code 200, with default header values.

Marked user as be removed successfully.
*/
type DeleteUsersUserIDOK struct {
}

func (o *DeleteUsersUserIDOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdOK ", 200)
}

func (o *DeleteUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersUserIDBadRequest creates a DeleteUsersUserIDBadRequest with default headers values
func NewDeleteUsersUserIDBadRequest() *DeleteUsersUserIDBadRequest {
	return &DeleteUsersUserIDBadRequest{}
}

/* DeleteUsersUserIDBadRequest describes a response with status code 400, with default header values.

Invalid user ID.
*/
type DeleteUsersUserIDBadRequest struct {
}

func (o *DeleteUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdBadRequest ", 400)
}

func (o *DeleteUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersUserIDUnauthorized creates a DeleteUsersUserIDUnauthorized with default headers values
func NewDeleteUsersUserIDUnauthorized() *DeleteUsersUserIDUnauthorized {
	return &DeleteUsersUserIDUnauthorized{}
}

/* DeleteUsersUserIDUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type DeleteUsersUserIDUnauthorized struct {
}

func (o *DeleteUsersUserIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdUnauthorized ", 401)
}

func (o *DeleteUsersUserIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersUserIDForbidden creates a DeleteUsersUserIDForbidden with default headers values
func NewDeleteUsersUserIDForbidden() *DeleteUsersUserIDForbidden {
	return &DeleteUsersUserIDForbidden{}
}

/* DeleteUsersUserIDForbidden describes a response with status code 403, with default header values.

User does not have permission of admin role.
*/
type DeleteUsersUserIDForbidden struct {
}

func (o *DeleteUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdForbidden ", 403)
}

func (o *DeleteUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersUserIDNotFound creates a DeleteUsersUserIDNotFound with default headers values
func NewDeleteUsersUserIDNotFound() *DeleteUsersUserIDNotFound {
	return &DeleteUsersUserIDNotFound{}
}

/* DeleteUsersUserIDNotFound describes a response with status code 404, with default header values.

User ID does not exist.
*/
type DeleteUsersUserIDNotFound struct {
}

func (o *DeleteUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdNotFound ", 404)
}

func (o *DeleteUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersUserIDInternalServerError creates a DeleteUsersUserIDInternalServerError with default headers values
func NewDeleteUsersUserIDInternalServerError() *DeleteUsersUserIDInternalServerError {
	return &DeleteUsersUserIDInternalServerError{}
}

/* DeleteUsersUserIDInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type DeleteUsersUserIDInternalServerError struct {
}

func (o *DeleteUsersUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}][%d] deleteUsersUserIdInternalServerError ", 500)
}

func (o *DeleteUsersUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
