// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// UpdateJobReader is a Reader for the UpdateJob structure.
type UpdateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateJobOK creates a UpdateJobOK with default headers values
func NewUpdateJobOK() *UpdateJobOK {
	return &UpdateJobOK{}
}

/* UpdateJobOK describes a response with status code 200, with default header values.

Job has been updated.
*/
type UpdateJobOK struct {
	Payload *models.JobModel
}

func (o *UpdateJobOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobOK  %+v", 200, o.Payload)
}
func (o *UpdateJobOK) GetPayload() *models.JobModel {
	return o.Payload
}

func (o *UpdateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobBadRequest creates a UpdateJobBadRequest with default headers values
func NewUpdateJobBadRequest() *UpdateJobBadRequest {
	return &UpdateJobBadRequest{}
}

/* UpdateJobBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type UpdateJobBadRequest struct {
	Payload *models.Error
}

func (o *UpdateJobBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateJobBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobUnauthorized creates a UpdateJobUnauthorized with default headers values
func NewUpdateJobUnauthorized() *UpdateJobUnauthorized {
	return &UpdateJobUnauthorized{}
}

/* UpdateJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type UpdateJobUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateJobUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateJobUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobForbidden creates a UpdateJobForbidden with default headers values
func NewUpdateJobForbidden() *UpdateJobForbidden {
	return &UpdateJobForbidden{}
}

/* UpdateJobForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type UpdateJobForbidden struct {
	Payload *models.Error
}

func (o *UpdateJobForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobForbidden  %+v", 403, o.Payload)
}
func (o *UpdateJobForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobNotFound creates a UpdateJobNotFound with default headers values
func NewUpdateJobNotFound() *UpdateJobNotFound {
	return &UpdateJobNotFound{}
}

/* UpdateJobNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type UpdateJobNotFound struct {
	Payload *models.Error
}

func (o *UpdateJobNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobNotFound  %+v", 404, o.Payload)
}
func (o *UpdateJobNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobInternalServerError creates a UpdateJobInternalServerError with default headers values
func NewUpdateJobInternalServerError() *UpdateJobInternalServerError {
	return &UpdateJobInternalServerError{}
}

/* UpdateJobInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type UpdateJobInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateJobInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/jobs/{id}][%d] updateJobInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateJobInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}