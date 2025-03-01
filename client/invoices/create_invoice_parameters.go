// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateInvoiceParams creates a new CreateInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInvoiceParams() *CreateInvoiceParams {
	return &CreateInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoiceParamsWithTimeout creates a new CreateInvoiceParams object
// with the ability to set a timeout on a request.
func NewCreateInvoiceParamsWithTimeout(timeout time.Duration) *CreateInvoiceParams {
	return &CreateInvoiceParams{
		timeout: timeout,
	}
}

// NewCreateInvoiceParamsWithContext creates a new CreateInvoiceParams object
// with the ability to set a context for a request.
func NewCreateInvoiceParamsWithContext(ctx context.Context) *CreateInvoiceParams {
	return &CreateInvoiceParams{
		Context: ctx,
	}
}

// NewCreateInvoiceParamsWithHTTPClient creates a new CreateInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInvoiceParamsWithHTTPClient(client *http.Client) *CreateInvoiceParams {
	return &CreateInvoiceParams{
		HTTPClient: client,
	}
}

/*
CreateInvoiceParams contains all the parameters to send to the API endpoint

	for the create invoice operation.

	Typically these are written to a http.Request.
*/
type CreateInvoiceParams struct {

	// CreateInvoiceRequest.
	CreateInvoiceRequest CreateInvoiceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceParams) WithDefaults() *CreateInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create invoice params
func (o *CreateInvoiceParams) WithTimeout(timeout time.Duration) *CreateInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice params
func (o *CreateInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice params
func (o *CreateInvoiceParams) WithContext(ctx context.Context) *CreateInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice params
func (o *CreateInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice params
func (o *CreateInvoiceParams) WithHTTPClient(client *http.Client) *CreateInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice params
func (o *CreateInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateInvoiceRequest adds the createInvoiceRequest to the create invoice params
func (o *CreateInvoiceParams) WithCreateInvoiceRequest(createInvoiceRequest CreateInvoiceBody) *CreateInvoiceParams {
	o.SetCreateInvoiceRequest(createInvoiceRequest)
	return o
}

// SetCreateInvoiceRequest adds the createInvoiceRequest to the create invoice params
func (o *CreateInvoiceParams) SetCreateInvoiceRequest(createInvoiceRequest CreateInvoiceBody) {
	o.CreateInvoiceRequest = createInvoiceRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateInvoiceRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
