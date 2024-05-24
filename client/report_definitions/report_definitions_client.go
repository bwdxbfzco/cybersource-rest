// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new report definitions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new report definitions API client with basic auth credentials.
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

// New creates a new report definitions API client with a bearer token for authentication.
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
Client for report definitions API
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

// ClientService is the interface for Client methods
type ClientService interface {
	GetResourceInfoByReportDefinition(params *GetResourceInfoByReportDefinitionParams, opts ...ClientOption) (*GetResourceInfoByReportDefinitionOK, error)

	GetResourceV2Info(params *GetResourceV2InfoParams, opts ...ClientOption) (*GetResourceV2InfoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetResourceInfoByReportDefinition gets report definition

View the attributes of an individual report type. For a list of values for reportDefinitionName, see the [Reporting Developer Guide](https://www.cybersource.com/developers/documentation/reporting_and_reconciliation/)
*/
func (a *Client) GetResourceInfoByReportDefinition(params *GetResourceInfoByReportDefinitionParams, opts ...ClientOption) (*GetResourceInfoByReportDefinitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceInfoByReportDefinitionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceInfoByReportDefinition",
		Method:             "GET",
		PathPattern:        "/reporting/v3/report-definitions/{reportDefinitionName}",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceInfoByReportDefinitionReader{formats: a.formats},
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
	success, ok := result.(*GetResourceInfoByReportDefinitionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceInfoByReportDefinition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResourceV2Info gets reporting resource information

View a list of supported reports and their attributes before subscribing to them.
*/
func (a *Client) GetResourceV2Info(params *GetResourceV2InfoParams, opts ...ClientOption) (*GetResourceV2InfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceV2InfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceV2Info",
		Method:             "GET",
		PathPattern:        "/reporting/v3/report-definitions",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceV2InfoReader{formats: a.formats},
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
	success, ok := result.(*GetResourceV2InfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceV2Info: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
