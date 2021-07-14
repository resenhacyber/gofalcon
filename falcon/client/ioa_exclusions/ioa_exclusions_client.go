// Code generated by go-swagger; DO NOT EDIT.

package ioa_exclusions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ioa exclusions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ioa exclusions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateIOAExclusionsV1(params *CreateIOAExclusionsV1Params, opts ...ClientOption) (*CreateIOAExclusionsV1OK, error)

	DeleteIOAExclusionsV1(params *DeleteIOAExclusionsV1Params, opts ...ClientOption) (*DeleteIOAExclusionsV1OK, error)

	GetIOAExclusionsV1(params *GetIOAExclusionsV1Params, opts ...ClientOption) (*GetIOAExclusionsV1OK, error)

	QueryIOAExclusionsV1(params *QueryIOAExclusionsV1Params, opts ...ClientOption) (*QueryIOAExclusionsV1OK, error)

	UpdateIOAExclusionsV1(params *UpdateIOAExclusionsV1Params, opts ...ClientOption) (*UpdateIOAExclusionsV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateIOAExclusionsV1 creates the i o a exclusions
*/
func (a *Client) CreateIOAExclusionsV1(params *CreateIOAExclusionsV1Params, opts ...ClientOption) (*CreateIOAExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIOAExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createIOAExclusionsV1",
		Method:             "POST",
		PathPattern:        "/policy/entities/ioa-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIOAExclusionsV1Reader{formats: a.formats},
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
	success, ok := result.(*CreateIOAExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateIOAExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteIOAExclusionsV1 deletes the i o a exclusions by id
*/
func (a *Client) DeleteIOAExclusionsV1(params *DeleteIOAExclusionsV1Params, opts ...ClientOption) (*DeleteIOAExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIOAExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteIOAExclusionsV1",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/ioa-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIOAExclusionsV1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteIOAExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIOAExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetIOAExclusionsV1 gets a set of i o a exclusions by specifying their i ds
*/
func (a *Client) GetIOAExclusionsV1(params *GetIOAExclusionsV1Params, opts ...ClientOption) (*GetIOAExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIOAExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIOAExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/entities/ioa-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIOAExclusionsV1Reader{formats: a.formats},
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
	success, ok := result.(*GetIOAExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIOAExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryIOAExclusionsV1 searches for i o a exclusions
*/
func (a *Client) QueryIOAExclusionsV1(params *QueryIOAExclusionsV1Params, opts ...ClientOption) (*QueryIOAExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryIOAExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryIOAExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/queries/ioa-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryIOAExclusionsV1Reader{formats: a.formats},
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
	success, ok := result.(*QueryIOAExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryIOAExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateIOAExclusionsV1 updates the i o a exclusions
*/
func (a *Client) UpdateIOAExclusionsV1(params *UpdateIOAExclusionsV1Params, opts ...ClientOption) (*UpdateIOAExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIOAExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateIOAExclusionsV1",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/ioa-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIOAExclusionsV1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateIOAExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateIOAExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
