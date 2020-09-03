// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostProjectsProjectIDWebhookPoliciesTestReader is a Reader for the PostProjectsProjectIDWebhookPoliciesTest structure.
type PostProjectsProjectIDWebhookPoliciesTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectsProjectIDWebhookPoliciesTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProjectsProjectIDWebhookPoliciesTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectsProjectIDWebhookPoliciesTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProjectsProjectIDWebhookPoliciesTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostProjectsProjectIDWebhookPoliciesTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProjectsProjectIDWebhookPoliciesTestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostProjectsProjectIDWebhookPoliciesTestOK creates a PostProjectsProjectIDWebhookPoliciesTestOK with default headers values
func NewPostProjectsProjectIDWebhookPoliciesTestOK() *PostProjectsProjectIDWebhookPoliciesTestOK {
	return &PostProjectsProjectIDWebhookPoliciesTestOK{}
}

/*PostProjectsProjectIDWebhookPoliciesTestOK handles this case with default header values.

Test webhook connection successfully.
*/
type PostProjectsProjectIDWebhookPoliciesTestOK struct {
}

func (o *PostProjectsProjectIDWebhookPoliciesTestOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/webhook/policies/test][%d] postProjectsProjectIdWebhookPoliciesTestOK ", 200)
}

func (o *PostProjectsProjectIDWebhookPoliciesTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDWebhookPoliciesTestBadRequest creates a PostProjectsProjectIDWebhookPoliciesTestBadRequest with default headers values
func NewPostProjectsProjectIDWebhookPoliciesTestBadRequest() *PostProjectsProjectIDWebhookPoliciesTestBadRequest {
	return &PostProjectsProjectIDWebhookPoliciesTestBadRequest{}
}

/*PostProjectsProjectIDWebhookPoliciesTestBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type PostProjectsProjectIDWebhookPoliciesTestBadRequest struct {
}

func (o *PostProjectsProjectIDWebhookPoliciesTestBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/webhook/policies/test][%d] postProjectsProjectIdWebhookPoliciesTestBadRequest ", 400)
}

func (o *PostProjectsProjectIDWebhookPoliciesTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDWebhookPoliciesTestUnauthorized creates a PostProjectsProjectIDWebhookPoliciesTestUnauthorized with default headers values
func NewPostProjectsProjectIDWebhookPoliciesTestUnauthorized() *PostProjectsProjectIDWebhookPoliciesTestUnauthorized {
	return &PostProjectsProjectIDWebhookPoliciesTestUnauthorized{}
}

/*PostProjectsProjectIDWebhookPoliciesTestUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostProjectsProjectIDWebhookPoliciesTestUnauthorized struct {
}

func (o *PostProjectsProjectIDWebhookPoliciesTestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/webhook/policies/test][%d] postProjectsProjectIdWebhookPoliciesTestUnauthorized ", 401)
}

func (o *PostProjectsProjectIDWebhookPoliciesTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDWebhookPoliciesTestForbidden creates a PostProjectsProjectIDWebhookPoliciesTestForbidden with default headers values
func NewPostProjectsProjectIDWebhookPoliciesTestForbidden() *PostProjectsProjectIDWebhookPoliciesTestForbidden {
	return &PostProjectsProjectIDWebhookPoliciesTestForbidden{}
}

/*PostProjectsProjectIDWebhookPoliciesTestForbidden handles this case with default header values.

User have no permission to get webhook policy of the project.
*/
type PostProjectsProjectIDWebhookPoliciesTestForbidden struct {
}

func (o *PostProjectsProjectIDWebhookPoliciesTestForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/webhook/policies/test][%d] postProjectsProjectIdWebhookPoliciesTestForbidden ", 403)
}

func (o *PostProjectsProjectIDWebhookPoliciesTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDWebhookPoliciesTestInternalServerError creates a PostProjectsProjectIDWebhookPoliciesTestInternalServerError with default headers values
func NewPostProjectsProjectIDWebhookPoliciesTestInternalServerError() *PostProjectsProjectIDWebhookPoliciesTestInternalServerError {
	return &PostProjectsProjectIDWebhookPoliciesTestInternalServerError{}
}

/*PostProjectsProjectIDWebhookPoliciesTestInternalServerError handles this case with default header values.

Internal server errors.
*/
type PostProjectsProjectIDWebhookPoliciesTestInternalServerError struct {
}

func (o *PostProjectsProjectIDWebhookPoliciesTestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/webhook/policies/test][%d] postProjectsProjectIdWebhookPoliciesTestInternalServerError ", 500)
}

func (o *PostProjectsProjectIDWebhookPoliciesTestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
