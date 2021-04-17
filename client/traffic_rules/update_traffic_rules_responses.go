// Code generated by go-swagger; DO NOT EDIT.

package traffic_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// UpdateTrafficRulesReader is a Reader for the UpdateTrafficRules structure.
type UpdateTrafficRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTrafficRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateTrafficRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTrafficRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTrafficRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTrafficRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTrafficRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTrafficRulesCreated creates a UpdateTrafficRulesCreated with default headers values
func NewUpdateTrafficRulesCreated() *UpdateTrafficRulesCreated {
	return &UpdateTrafficRulesCreated{}
}

/* UpdateTrafficRulesCreated describes a response with status code 201, with default header values.

Infrastructure session has been created to edit the traffic rules. To check the progress, track the session `state`.
*/
type UpdateTrafficRulesCreated struct {
	Payload *models.GlobalNetworkTrafficRulesModel
}

func (o *UpdateTrafficRulesCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/trafficRules][%d] updateTrafficRulesCreated  %+v", 201, o.Payload)
}
func (o *UpdateTrafficRulesCreated) GetPayload() *models.GlobalNetworkTrafficRulesModel {
	return o.Payload
}

func (o *UpdateTrafficRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GlobalNetworkTrafficRulesModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrafficRulesBadRequest creates a UpdateTrafficRulesBadRequest with default headers values
func NewUpdateTrafficRulesBadRequest() *UpdateTrafficRulesBadRequest {
	return &UpdateTrafficRulesBadRequest{}
}

/* UpdateTrafficRulesBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type UpdateTrafficRulesBadRequest struct {
	Payload *models.Error
}

func (o *UpdateTrafficRulesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/trafficRules][%d] updateTrafficRulesBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateTrafficRulesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTrafficRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrafficRulesUnauthorized creates a UpdateTrafficRulesUnauthorized with default headers values
func NewUpdateTrafficRulesUnauthorized() *UpdateTrafficRulesUnauthorized {
	return &UpdateTrafficRulesUnauthorized{}
}

/* UpdateTrafficRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type UpdateTrafficRulesUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateTrafficRulesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/trafficRules][%d] updateTrafficRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateTrafficRulesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTrafficRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrafficRulesForbidden creates a UpdateTrafficRulesForbidden with default headers values
func NewUpdateTrafficRulesForbidden() *UpdateTrafficRulesForbidden {
	return &UpdateTrafficRulesForbidden{}
}

/* UpdateTrafficRulesForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type UpdateTrafficRulesForbidden struct {
	Payload *models.Error
}

func (o *UpdateTrafficRulesForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/trafficRules][%d] updateTrafficRulesForbidden  %+v", 403, o.Payload)
}
func (o *UpdateTrafficRulesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTrafficRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrafficRulesInternalServerError creates a UpdateTrafficRulesInternalServerError with default headers values
func NewUpdateTrafficRulesInternalServerError() *UpdateTrafficRulesInternalServerError {
	return &UpdateTrafficRulesInternalServerError{}
}

/* UpdateTrafficRulesInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type UpdateTrafficRulesInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateTrafficRulesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/trafficRules][%d] updateTrafficRulesInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateTrafficRulesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTrafficRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}