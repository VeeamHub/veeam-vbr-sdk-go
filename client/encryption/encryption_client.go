// Code generated by go-swagger; DO NOT EDIT.

package encryption

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new encryption API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for encryption API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEncryptionPassword(params *CreateEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEncryptionPasswordCreated, error)

	DeleteEncryptionPassword(params *DeleteEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEncryptionPasswordNoContent, error)

	GetAllEncryptionPasswords(params *GetAllEncryptionPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllEncryptionPasswordsOK, error)

	GetEncryptionPassword(params *GetEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEncryptionPasswordOK, error)

	UpdateEncryptionPassword(params *UpdateEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEncryptionPasswordOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEncryptionPassword adds encryption password

  The HTTP POST request to the `/api/v1/encryptionPasswords` path allows you to add an encryption password.
*/
func (a *Client) CreateEncryptionPassword(params *CreateEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEncryptionPasswordCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEncryptionPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateEncryptionPassword",
		Method:             "POST",
		PathPattern:        "/api/v1/encryptionPasswords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEncryptionPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateEncryptionPasswordCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateEncryptionPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEncryptionPassword removes encryption password

  The HTTP DELETE request to the `/api/v1/encryptionPasswords/{id}` path allows you to remove an encryption password that has the specified `id`.
*/
func (a *Client) DeleteEncryptionPassword(params *DeleteEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteEncryptionPasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEncryptionPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEncryptionPassword",
		Method:             "DELETE",
		PathPattern:        "/api/v1/encryptionPasswords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEncryptionPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEncryptionPasswordNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEncryptionPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllEncryptionPasswords gets all encryption passwords

  The HTTP GET request to the `/api/v1/encryptionPasswords` path allows you to get an array of all passwords that are used for data encryption.
*/
func (a *Client) GetAllEncryptionPasswords(params *GetAllEncryptionPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllEncryptionPasswordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllEncryptionPasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAllEncryptionPasswords",
		Method:             "GET",
		PathPattern:        "/api/v1/encryptionPasswords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllEncryptionPasswordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllEncryptionPasswordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllEncryptionPasswords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEncryptionPassword gets encryption password

  The HTTP GET request to the `/api/v1/encryptionPasswords/{id}` path allows you to get an encryption password that has the specified `id`.
*/
func (a *Client) GetEncryptionPassword(params *GetEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEncryptionPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEncryptionPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEncryptionPassword",
		Method:             "GET",
		PathPattern:        "/api/v1/encryptionPasswords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEncryptionPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEncryptionPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEncryptionPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEncryptionPassword edits encryption password

  The HTTP PUT request to the `/api/v1/encryptionPasswords/{id}` path allows you to edit an encryption password that has the specified `id`.
*/
func (a *Client) UpdateEncryptionPassword(params *UpdateEncryptionPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateEncryptionPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEncryptionPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateEncryptionPassword",
		Method:             "PUT",
		PathPattern:        "/api/v1/encryptionPasswords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEncryptionPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateEncryptionPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateEncryptionPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}