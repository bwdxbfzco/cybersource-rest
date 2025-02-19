// Code generated by go-swagger; DO NOT EDIT.

package invoice_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoice settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new invoice settings API client with basic auth credentials.
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

// New creates a new invoice settings API client with a bearer token for authentication.
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
Client for invoice settings API
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

// WithAcceptApplicationHalJSONCharsetUTF8 sets the Accept header to "application/hal+json;charset=utf-8".
func WithAcceptApplicationHalJSONCharsetUTF8(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/hal+json;charset=utf-8"}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationJSONCharsetUTF8 sets the Accept header to "application/json;charset=utf-8".
func WithAcceptApplicationJSONCharsetUTF8(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json;charset=utf-8"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetInvoiceSettings(params *GetInvoiceSettingsParams, opts ...ClientOption) (*GetInvoiceSettingsOK, error)

	UpdateInvoiceSettings(params *UpdateInvoiceSettingsParams, opts ...ClientOption) (*UpdateInvoiceSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInvoiceSettings gets invoice settings

Allows you to retrieve the invoice settings for the payment page.
*/
func (a *Client) GetInvoiceSettings(params *GetInvoiceSettingsParams, opts ...ClientOption) (*GetInvoiceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoiceSettings",
		Method:             "GET",
		PathPattern:        "/invoicing/v2/invoiceSettings",
		ProducesMediaTypes: []string{"application/json", "application/hal+json", "application/json;charset=utf-8", "application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetInvoiceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInvoiceSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateInvoiceSettings updates invoice settings

Allows you to customize the payment page, the checkout experience, email communication and payer authentication. You can customize the invoice to match your brand with your business name, logo and brand colors, and a VAT Tax number. You can choose to capture the payers shipping details, phone number and email during the checkout process. You can add a custom message to all invoice emails and enable or disable payer authentication for invoice payments.
*/
func (a *Client) UpdateInvoiceSettings(params *UpdateInvoiceSettingsParams, opts ...ClientOption) (*UpdateInvoiceSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInvoiceSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInvoiceSettings",
		Method:             "PUT",
		PathPattern:        "/invoicing/v2/invoiceSettings",
		ProducesMediaTypes: []string{"application/json", "application/hal+json", "application/json;charset=utf-8", "application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInvoiceSettingsReader{formats: a.formats},
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
	success, ok := result.(*UpdateInvoiceSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateInvoiceSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
