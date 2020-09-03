// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchRetentionsIDExecutionsEidReader is a Reader for the PatchRetentionsIDExecutionsEid structure.
type PatchRetentionsIDExecutionsEidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRetentionsIDExecutionsEidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRetentionsIDExecutionsEidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRetentionsIDExecutionsEidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRetentionsIDExecutionsEidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRetentionsIDExecutionsEidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRetentionsIDExecutionsEidOK creates a PatchRetentionsIDExecutionsEidOK with default headers values
func NewPatchRetentionsIDExecutionsEidOK() *PatchRetentionsIDExecutionsEidOK {
	return &PatchRetentionsIDExecutionsEidOK{}
}

/*PatchRetentionsIDExecutionsEidOK handles this case with default header values.

Stop a Retention job successfully.
*/
type PatchRetentionsIDExecutionsEidOK struct {
}

func (o *PatchRetentionsIDExecutionsEidOK) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] patchRetentionsIdExecutionsEidOK ", 200)
}

func (o *PatchRetentionsIDExecutionsEidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRetentionsIDExecutionsEidUnauthorized creates a PatchRetentionsIDExecutionsEidUnauthorized with default headers values
func NewPatchRetentionsIDExecutionsEidUnauthorized() *PatchRetentionsIDExecutionsEidUnauthorized {
	return &PatchRetentionsIDExecutionsEidUnauthorized{}
}

/*PatchRetentionsIDExecutionsEidUnauthorized handles this case with default header values.

User need to log in first.
*/
type PatchRetentionsIDExecutionsEidUnauthorized struct {
}

func (o *PatchRetentionsIDExecutionsEidUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] patchRetentionsIdExecutionsEidUnauthorized ", 401)
}

func (o *PatchRetentionsIDExecutionsEidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRetentionsIDExecutionsEidForbidden creates a PatchRetentionsIDExecutionsEidForbidden with default headers values
func NewPatchRetentionsIDExecutionsEidForbidden() *PatchRetentionsIDExecutionsEidForbidden {
	return &PatchRetentionsIDExecutionsEidForbidden{}
}

/*PatchRetentionsIDExecutionsEidForbidden handles this case with default header values.

User have no permission.
*/
type PatchRetentionsIDExecutionsEidForbidden struct {
}

func (o *PatchRetentionsIDExecutionsEidForbidden) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] patchRetentionsIdExecutionsEidForbidden ", 403)
}

func (o *PatchRetentionsIDExecutionsEidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRetentionsIDExecutionsEidInternalServerError creates a PatchRetentionsIDExecutionsEidInternalServerError with default headers values
func NewPatchRetentionsIDExecutionsEidInternalServerError() *PatchRetentionsIDExecutionsEidInternalServerError {
	return &PatchRetentionsIDExecutionsEidInternalServerError{}
}

/*PatchRetentionsIDExecutionsEidInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PatchRetentionsIDExecutionsEidInternalServerError struct {
}

func (o *PatchRetentionsIDExecutionsEidInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] patchRetentionsIdExecutionsEidInternalServerError ", 500)
}

func (o *PatchRetentionsIDExecutionsEidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PatchRetentionsIDExecutionsEidBody patch retentions ID executions eid body
swagger:model PatchRetentionsIDExecutionsEidBody
*/
type PatchRetentionsIDExecutionsEidBody struct {

	// action
	Action string `json:"action,omitempty"`
}

// Validate validates this patch retentions ID executions eid body
func (o *PatchRetentionsIDExecutionsEidBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchRetentionsIDExecutionsEidBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchRetentionsIDExecutionsEidBody) UnmarshalBinary(b []byte) error {
	var res PatchRetentionsIDExecutionsEidBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
