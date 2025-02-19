// Code generated by go-swagger; DO NOT EDIT.

package transaction_batches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new transaction batches API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new transaction batches API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new transaction batches API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for transaction batches API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationJSONCharsetUTF8 sets the Content-Type header to "application/json;charset=utf-8".
func WithContentTypeApplicationJSONCharsetUTF8(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json;charset=utf-8"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationHalJSON sets the Accept header to "application/hal+json".
func WithAcceptApplicationHalJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/hal+json"}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationXML sets the Accept header to "application/xml".
func WithAcceptApplicationXML(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/xml"}
}

// WithAcceptTextCsv sets the Accept header to "text/csv".
func WithAcceptTextCsv(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/csv"}
}

// WithAcceptTextVndCybersourceMapCsv sets the Accept header to "text/vnd.cybersource.map-csv".
func WithAcceptTextVndCybersourceMapCsv(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/vnd.cybersource.map-csv"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTransactionBatchDetails(params *GetTransactionBatchDetailsParams, opts ...ClientOption) (*GetTransactionBatchDetailsOK, error)

	GetTransactionBatchID(params *GetTransactionBatchIDParams, opts ...ClientOption) (*GetTransactionBatchIDOK, error)

	GetTransactionBatches(params *GetTransactionBatchesParams, opts ...ClientOption) (*GetTransactionBatchesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetTransactionBatchDetails gets transaction details for a given batch Id

Provides real-time detailed status information about the transactions that you previously uploaded in the Business Center or processed with the Offline Transaction File Submission service.
*/
func (a *Client) GetTransactionBatchDetails(params *GetTransactionBatchDetailsParams, opts ...ClientOption) (*GetTransactionBatchDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionBatchDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactionBatchDetails",
		Method:             "GET",
		PathPattern:        "/pts/v1/transaction-batch-details/{id}",
		ProducesMediaTypes: []string{"text/csv", "application/xml", "text/vnd.cybersource.map-csv"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionBatchDetailsReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionBatchDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionBatchDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTransactionBatchID gets individual batch file

Provide the search range
*/
func (a *Client) GetTransactionBatchID(params *GetTransactionBatchIDParams, opts ...ClientOption) (*GetTransactionBatchIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionBatchIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactionBatchId",
		Method:             "GET",
		PathPattern:        "/pts/v1/transaction-batches/{id}",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionBatchIDReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionBatchIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionBatchId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTransactionBatches gets a list of batch files

Provide the search range
*/
func (a *Client) GetTransactionBatches(params *GetTransactionBatchesParams, opts ...ClientOption) (*GetTransactionBatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionBatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactionBatches",
		Method:             "GET",
		PathPattern:        "/pts/v1/transaction-batches",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionBatchesReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionBatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionBatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
