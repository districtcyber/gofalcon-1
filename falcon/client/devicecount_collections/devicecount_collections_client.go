// Code generated by go-swagger; DO NOT EDIT.

package devicecount_collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new devicecount collections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devicecount collections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AggregateDeviceCountCollection(params *AggregateDeviceCountCollectionParams, opts ...ClientOption) (*AggregateDeviceCountCollectionOK, error)

	GetDeviceCountCollectionQueriesByFilter(params *GetDeviceCountCollectionQueriesByFilterParams, opts ...ClientOption) (*GetDeviceCountCollectionQueriesByFilterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AggregateDeviceCountCollection retrieves aggregate host devices count based on the matched filter
*/
func (a *Client) AggregateDeviceCountCollection(params *AggregateDeviceCountCollectionParams, opts ...ClientOption) (*AggregateDeviceCountCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAggregateDeviceCountCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AggregateDeviceCountCollection",
		Method:             "POST",
		PathPattern:        "/falcon-complete-dashboards/aggregates/devicecount-collections/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AggregateDeviceCountCollectionReader{formats: a.formats},
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
	success, ok := result.(*AggregateDeviceCountCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AggregateDeviceCountCollectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceCountCollectionQueriesByFilter retrieves device count collection ids that match the provided f q l filter criteria with scrolling enabled
*/
func (a *Client) GetDeviceCountCollectionQueriesByFilter(params *GetDeviceCountCollectionQueriesByFilterParams, opts ...ClientOption) (*GetDeviceCountCollectionQueriesByFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCountCollectionQueriesByFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceCountCollectionQueriesByFilter",
		Method:             "GET",
		PathPattern:        "/falcon-complete-dashboards/queries/devicecount-collections/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCountCollectionQueriesByFilterReader{formats: a.formats},
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
	success, ok := result.(*GetDeviceCountCollectionQueriesByFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceCountCollectionQueriesByFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
