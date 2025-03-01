// Code generated by go-swagger; DO NOT EDIT.

package void

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

// NewVoidPaymentParams creates a new VoidPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVoidPaymentParams() *VoidPaymentParams {
	return &VoidPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVoidPaymentParamsWithTimeout creates a new VoidPaymentParams object
// with the ability to set a timeout on a request.
func NewVoidPaymentParamsWithTimeout(timeout time.Duration) *VoidPaymentParams {
	return &VoidPaymentParams{
		timeout: timeout,
	}
}

// NewVoidPaymentParamsWithContext creates a new VoidPaymentParams object
// with the ability to set a context for a request.
func NewVoidPaymentParamsWithContext(ctx context.Context) *VoidPaymentParams {
	return &VoidPaymentParams{
		Context: ctx,
	}
}

// NewVoidPaymentParamsWithHTTPClient creates a new VoidPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewVoidPaymentParamsWithHTTPClient(client *http.Client) *VoidPaymentParams {
	return &VoidPaymentParams{
		HTTPClient: client,
	}
}

/*
VoidPaymentParams contains all the parameters to send to the API endpoint

	for the void payment operation.

	Typically these are written to a http.Request.
*/
type VoidPaymentParams struct {

	/* ID.

	   The payment ID returned from a previous payment request.
	*/
	ID string

	// VoidPaymentRequest.
	VoidPaymentRequest VoidPaymentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the void payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoidPaymentParams) WithDefaults() *VoidPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the void payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoidPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) WithTimeout(timeout time.Duration) *VoidPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void payment params
func (o *VoidPaymentParams) WithContext(ctx context.Context) *VoidPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void payment params
func (o *VoidPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) WithHTTPClient(client *http.Client) *VoidPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the void payment params
func (o *VoidPaymentParams) WithID(id string) *VoidPaymentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void payment params
func (o *VoidPaymentParams) SetID(id string) {
	o.ID = id
}

// WithVoidPaymentRequest adds the voidPaymentRequest to the void payment params
func (o *VoidPaymentParams) WithVoidPaymentRequest(voidPaymentRequest VoidPaymentBody) *VoidPaymentParams {
	o.SetVoidPaymentRequest(voidPaymentRequest)
	return o
}

// SetVoidPaymentRequest adds the voidPaymentRequest to the void payment params
func (o *VoidPaymentParams) SetVoidPaymentRequest(voidPaymentRequest VoidPaymentBody) {
	o.VoidPaymentRequest = voidPaymentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VoidPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.VoidPaymentRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
