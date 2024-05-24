// Code generated by go-swagger; DO NOT EDIT.

package verification

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

// NewVerifyCustomerAddressParams creates a new VerifyCustomerAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifyCustomerAddressParams() *VerifyCustomerAddressParams {
	return &VerifyCustomerAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyCustomerAddressParamsWithTimeout creates a new VerifyCustomerAddressParams object
// with the ability to set a timeout on a request.
func NewVerifyCustomerAddressParamsWithTimeout(timeout time.Duration) *VerifyCustomerAddressParams {
	return &VerifyCustomerAddressParams{
		timeout: timeout,
	}
}

// NewVerifyCustomerAddressParamsWithContext creates a new VerifyCustomerAddressParams object
// with the ability to set a context for a request.
func NewVerifyCustomerAddressParamsWithContext(ctx context.Context) *VerifyCustomerAddressParams {
	return &VerifyCustomerAddressParams{
		Context: ctx,
	}
}

// NewVerifyCustomerAddressParamsWithHTTPClient creates a new VerifyCustomerAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifyCustomerAddressParamsWithHTTPClient(client *http.Client) *VerifyCustomerAddressParams {
	return &VerifyCustomerAddressParams{
		HTTPClient: client,
	}
}

/*
VerifyCustomerAddressParams contains all the parameters to send to the API endpoint

	for the verify customer address operation.

	Typically these are written to a http.Request.
*/
type VerifyCustomerAddressParams struct {

	// VerifyCustomerAddressRequest.
	VerifyCustomerAddressRequest VerifyCustomerAddressBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify customer address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyCustomerAddressParams) WithDefaults() *VerifyCustomerAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify customer address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifyCustomerAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify customer address params
func (o *VerifyCustomerAddressParams) WithTimeout(timeout time.Duration) *VerifyCustomerAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify customer address params
func (o *VerifyCustomerAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify customer address params
func (o *VerifyCustomerAddressParams) WithContext(ctx context.Context) *VerifyCustomerAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify customer address params
func (o *VerifyCustomerAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify customer address params
func (o *VerifyCustomerAddressParams) WithHTTPClient(client *http.Client) *VerifyCustomerAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify customer address params
func (o *VerifyCustomerAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVerifyCustomerAddressRequest adds the verifyCustomerAddressRequest to the verify customer address params
func (o *VerifyCustomerAddressParams) WithVerifyCustomerAddressRequest(verifyCustomerAddressRequest VerifyCustomerAddressBody) *VerifyCustomerAddressParams {
	o.SetVerifyCustomerAddressRequest(verifyCustomerAddressRequest)
	return o
}

// SetVerifyCustomerAddressRequest adds the verifyCustomerAddressRequest to the verify customer address params
func (o *VerifyCustomerAddressParams) SetVerifyCustomerAddressRequest(verifyCustomerAddressRequest VerifyCustomerAddressBody) {
	o.VerifyCustomerAddressRequest = verifyCustomerAddressRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyCustomerAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.VerifyCustomerAddressRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
