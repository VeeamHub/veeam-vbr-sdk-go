// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// UpdateRepositoryReader is a Reader for the UpdateRepository structure.
type UpdateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryCreated creates a UpdateRepositoryCreated with default headers values
func NewUpdateRepositoryCreated() *UpdateRepositoryCreated {
	return &UpdateRepositoryCreated{}
}

/* UpdateRepositoryCreated describes a response with status code 201, with default header values.

Infrastructure session has been created to edit the repository. To check the progress, track the session `state`.
*/
type UpdateRepositoryCreated struct {
	Payload *models.SessionModel
}

func (o *UpdateRepositoryCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryCreated  %+v", 201, o.Payload)
}
func (o *UpdateRepositoryCreated) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *UpdateRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryBadRequest creates a UpdateRepositoryBadRequest with default headers values
func NewUpdateRepositoryBadRequest() *UpdateRepositoryBadRequest {
	return &UpdateRepositoryBadRequest{}
}

/* UpdateRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type UpdateRepositoryBadRequest struct {
	Payload *models.Error
}

func (o *UpdateRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRepositoryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryUnauthorized creates a UpdateRepositoryUnauthorized with default headers values
func NewUpdateRepositoryUnauthorized() *UpdateRepositoryUnauthorized {
	return &UpdateRepositoryUnauthorized{}
}

/* UpdateRepositoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type UpdateRepositoryUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRepositoryUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryForbidden creates a UpdateRepositoryForbidden with default headers values
func NewUpdateRepositoryForbidden() *UpdateRepositoryForbidden {
	return &UpdateRepositoryForbidden{}
}

/* UpdateRepositoryForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type UpdateRepositoryForbidden struct {
	Payload *models.Error
}

func (o *UpdateRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRepositoryForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryNotFound creates a UpdateRepositoryNotFound with default headers values
func NewUpdateRepositoryNotFound() *UpdateRepositoryNotFound {
	return &UpdateRepositoryNotFound{}
}

/* UpdateRepositoryNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type UpdateRepositoryNotFound struct {
	Payload *models.Error
}

func (o *UpdateRepositoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRepositoryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryInternalServerError creates a UpdateRepositoryInternalServerError with default headers values
func NewUpdateRepositoryInternalServerError() *UpdateRepositoryInternalServerError {
	return &UpdateRepositoryInternalServerError{}
}

/* UpdateRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type UpdateRepositoryInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/backupInfrastructure/repositories/{id}][%d] updateRepositoryInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRepositoryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}