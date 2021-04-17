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

// GetAllTrafficRulesReader is a Reader for the GetAllTrafficRules structure.
type GetAllTrafficRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTrafficRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTrafficRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllTrafficRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllTrafficRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllTrafficRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllTrafficRulesOK creates a GetAllTrafficRulesOK with default headers values
func NewGetAllTrafficRulesOK() *GetAllTrafficRulesOK {
	return &GetAllTrafficRulesOK{}
}

/* GetAllTrafficRulesOK describes a response with status code 200, with default header values.

OK
*/
type GetAllTrafficRulesOK struct {
	Payload *models.GlobalNetworkTrafficRulesModel
}

func (o *GetAllTrafficRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/trafficRules][%d] getAllTrafficRulesOK  %+v", 200, o.Payload)
}
func (o *GetAllTrafficRulesOK) GetPayload() *models.GlobalNetworkTrafficRulesModel {
	return o.Payload
}

func (o *GetAllTrafficRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GlobalNetworkTrafficRulesModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTrafficRulesUnauthorized creates a GetAllTrafficRulesUnauthorized with default headers values
func NewGetAllTrafficRulesUnauthorized() *GetAllTrafficRulesUnauthorized {
	return &GetAllTrafficRulesUnauthorized{}
}

/* GetAllTrafficRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetAllTrafficRulesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAllTrafficRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/trafficRules][%d] getAllTrafficRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAllTrafficRulesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllTrafficRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTrafficRulesForbidden creates a GetAllTrafficRulesForbidden with default headers values
func NewGetAllTrafficRulesForbidden() *GetAllTrafficRulesForbidden {
	return &GetAllTrafficRulesForbidden{}
}

/* GetAllTrafficRulesForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetAllTrafficRulesForbidden struct {
	Payload *models.Error
}

func (o *GetAllTrafficRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/trafficRules][%d] getAllTrafficRulesForbidden  %+v", 403, o.Payload)
}
func (o *GetAllTrafficRulesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllTrafficRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTrafficRulesInternalServerError creates a GetAllTrafficRulesInternalServerError with default headers values
func NewGetAllTrafficRulesInternalServerError() *GetAllTrafficRulesInternalServerError {
	return &GetAllTrafficRulesInternalServerError{}
}

/* GetAllTrafficRulesInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetAllTrafficRulesInternalServerError struct {
	Payload *models.Error
}

func (o *GetAllTrafficRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/trafficRules][%d] getAllTrafficRulesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAllTrafficRulesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllTrafficRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}