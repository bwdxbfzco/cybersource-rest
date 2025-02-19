// Code generated by go-swagger; DO NOT EDIT.

package invoice_settings

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

// NewGetInvoiceSettingsParams creates a new GetInvoiceSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInvoiceSettingsParams() *GetInvoiceSettingsParams {
	return &GetInvoiceSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceSettingsParamsWithTimeout creates a new GetInvoiceSettingsParams object
// with the ability to set a timeout on a request.
func NewGetInvoiceSettingsParamsWithTimeout(timeout time.Duration) *GetInvoiceSettingsParams {
	return &GetInvoiceSettingsParams{
		timeout: timeout,
	}
}

// NewGetInvoiceSettingsParamsWithContext creates a new GetInvoiceSettingsParams object
// with the ability to set a context for a request.
func NewGetInvoiceSettingsParamsWithContext(ctx context.Context) *GetInvoiceSettingsParams {
	return &GetInvoiceSettingsParams{
		Context: ctx,
	}
}

// NewGetInvoiceSettingsParamsWithHTTPClient creates a new GetInvoiceSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInvoiceSettingsParamsWithHTTPClient(client *http.Client) *GetInvoiceSettingsParams {
	return &GetInvoiceSettingsParams{
		HTTPClient: client,
	}
}

/*
GetInvoiceSettingsParams contains all the parameters to send to the API endpoint

	for the get invoice settings operation.

	Typically these are written to a http.Request.
*/
type GetInvoiceSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get invoice settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceSettingsParams) WithDefaults() *GetInvoiceSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get invoice settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get invoice settings params
func (o *GetInvoiceSettingsParams) WithTimeout(timeout time.Duration) *GetInvoiceSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice settings params
func (o *GetInvoiceSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice settings params
func (o *GetInvoiceSettingsParams) WithContext(ctx context.Context) *GetInvoiceSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice settings params
func (o *GetInvoiceSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice settings params
func (o *GetInvoiceSettingsParams) WithHTTPClient(client *http.Client) *GetInvoiceSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice settings params
func (o *GetInvoiceSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
