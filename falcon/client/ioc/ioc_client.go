// Code generated by go-swagger; DO NOT EDIT.

package ioc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ioc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ioc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IndicatorCombinedV1(params *IndicatorCombinedV1Params, opts ...ClientOption) (*IndicatorCombinedV1OK, error)

	IndicatorCreateV1(params *IndicatorCreateV1Params, opts ...ClientOption) (*IndicatorCreateV1Created, error)

	IndicatorDeleteV1(params *IndicatorDeleteV1Params, opts ...ClientOption) (*IndicatorDeleteV1OK, error)

	IndicatorGetV1(params *IndicatorGetV1Params, opts ...ClientOption) (*IndicatorGetV1OK, error)

	IndicatorSearchV1(params *IndicatorSearchV1Params, opts ...ClientOption) (*IndicatorSearchV1OK, error)

	IndicatorUpdateV1(params *IndicatorUpdateV1Params, opts ...ClientOption) (*IndicatorUpdateV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IndicatorCombinedV1 gets combined for indicators
*/
func (a *Client) IndicatorCombinedV1(params *IndicatorCombinedV1Params, opts ...ClientOption) (*IndicatorCombinedV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorCombinedV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.combined.v1",
		Method:             "GET",
		PathPattern:        "/iocs/combined/indicator/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorCombinedV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorCombinedV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IndicatorCombinedV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IndicatorCreateV1 creates indicators
*/
func (a *Client) IndicatorCreateV1(params *IndicatorCreateV1Params, opts ...ClientOption) (*IndicatorCreateV1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorCreateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.create.v1",
		Method:             "POST",
		PathPattern:        "/iocs/entities/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorCreateV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorCreateV1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for indicator.create.v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IndicatorDeleteV1 deletes indicators by ids
*/
func (a *Client) IndicatorDeleteV1(params *IndicatorDeleteV1Params, opts ...ClientOption) (*IndicatorDeleteV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorDeleteV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.delete.v1",
		Method:             "DELETE",
		PathPattern:        "/iocs/entities/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorDeleteV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorDeleteV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IndicatorDeleteV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IndicatorGetV1 gets indicators by ids
*/
func (a *Client) IndicatorGetV1(params *IndicatorGetV1Params, opts ...ClientOption) (*IndicatorGetV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorGetV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.get.v1",
		Method:             "GET",
		PathPattern:        "/iocs/entities/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorGetV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorGetV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IndicatorGetV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IndicatorSearchV1 searches for indicators
*/
func (a *Client) IndicatorSearchV1(params *IndicatorSearchV1Params, opts ...ClientOption) (*IndicatorSearchV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorSearchV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.search.v1",
		Method:             "GET",
		PathPattern:        "/iocs/queries/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorSearchV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorSearchV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IndicatorSearchV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IndicatorUpdateV1 updates indicators
*/
func (a *Client) IndicatorUpdateV1(params *IndicatorUpdateV1Params, opts ...ClientOption) (*IndicatorUpdateV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndicatorUpdateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "indicator.update.v1",
		Method:             "PATCH",
		PathPattern:        "/iocs/entities/indicators/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndicatorUpdateV1Reader{formats: a.formats},
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
	success, ok := result.(*IndicatorUpdateV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IndicatorUpdateV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
