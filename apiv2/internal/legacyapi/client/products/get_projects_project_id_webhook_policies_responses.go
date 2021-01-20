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

// GetProjectsProjectIDWebhookPoliciesReader is a Reader for the GetProjectsProjectIDWebhookPolicies structure.
type GetProjectsProjectIDWebhookPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDWebhookPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDWebhookPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDWebhookPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDWebhookPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDWebhookPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDWebhookPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDWebhookPoliciesOK creates a GetProjectsProjectIDWebhookPoliciesOK with default headers values
func NewGetProjectsProjectIDWebhookPoliciesOK() *GetProjectsProjectIDWebhookPoliciesOK {
	return &GetProjectsProjectIDWebhookPoliciesOK{}
}

/* GetProjectsProjectIDWebhookPoliciesOK describes a response with status code 200, with default header values.

List project webhook policies successfully.
*/
type GetProjectsProjectIDWebhookPoliciesOK struct {
	Payload []*legacy.WebhookPolicy
}

func (o *GetProjectsProjectIDWebhookPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies][%d] getProjectsProjectIdWebhookPoliciesOK  %+v", 200, o.Payload)
}
func (o *GetProjectsProjectIDWebhookPoliciesOK) GetPayload() []*legacy.WebhookPolicy {
	return o.Payload
}

func (o *GetProjectsProjectIDWebhookPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesBadRequest creates a GetProjectsProjectIDWebhookPoliciesBadRequest with default headers values
func NewGetProjectsProjectIDWebhookPoliciesBadRequest() *GetProjectsProjectIDWebhookPoliciesBadRequest {
	return &GetProjectsProjectIDWebhookPoliciesBadRequest{}
}

/* GetProjectsProjectIDWebhookPoliciesBadRequest describes a response with status code 400, with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDWebhookPoliciesBadRequest struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies][%d] getProjectsProjectIdWebhookPoliciesBadRequest ", 400)
}

func (o *GetProjectsProjectIDWebhookPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesUnauthorized creates a GetProjectsProjectIDWebhookPoliciesUnauthorized with default headers values
func NewGetProjectsProjectIDWebhookPoliciesUnauthorized() *GetProjectsProjectIDWebhookPoliciesUnauthorized {
	return &GetProjectsProjectIDWebhookPoliciesUnauthorized{}
}

/* GetProjectsProjectIDWebhookPoliciesUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetProjectsProjectIDWebhookPoliciesUnauthorized struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies][%d] getProjectsProjectIdWebhookPoliciesUnauthorized ", 401)
}

func (o *GetProjectsProjectIDWebhookPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesForbidden creates a GetProjectsProjectIDWebhookPoliciesForbidden with default headers values
func NewGetProjectsProjectIDWebhookPoliciesForbidden() *GetProjectsProjectIDWebhookPoliciesForbidden {
	return &GetProjectsProjectIDWebhookPoliciesForbidden{}
}

/* GetProjectsProjectIDWebhookPoliciesForbidden describes a response with status code 403, with default header values.

User have no permission to list webhook policies of the project.
*/
type GetProjectsProjectIDWebhookPoliciesForbidden struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies][%d] getProjectsProjectIdWebhookPoliciesForbidden ", 403)
}

func (o *GetProjectsProjectIDWebhookPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesInternalServerError creates a GetProjectsProjectIDWebhookPoliciesInternalServerError with default headers values
func NewGetProjectsProjectIDWebhookPoliciesInternalServerError() *GetProjectsProjectIDWebhookPoliciesInternalServerError {
	return &GetProjectsProjectIDWebhookPoliciesInternalServerError{}
}

/* GetProjectsProjectIDWebhookPoliciesInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDWebhookPoliciesInternalServerError struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies][%d] getProjectsProjectIdWebhookPoliciesInternalServerError ", 500)
}

func (o *GetProjectsProjectIDWebhookPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
