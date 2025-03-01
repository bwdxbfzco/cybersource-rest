// Code generated by go-swagger; DO NOT EDIT.

package taxes

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

// NewVoidTaxParams creates a new VoidTaxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVoidTaxParams() *VoidTaxParams {
	return &VoidTaxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVoidTaxParamsWithTimeout creates a new VoidTaxParams object
// with the ability to set a timeout on a request.
func NewVoidTaxParamsWithTimeout(timeout time.Duration) *VoidTaxParams {
	return &VoidTaxParams{
		timeout: timeout,
	}
}

// NewVoidTaxParamsWithContext creates a new VoidTaxParams object
// with the ability to set a context for a request.
func NewVoidTaxParamsWithContext(ctx context.Context) *VoidTaxParams {
	return &VoidTaxParams{
		Context: ctx,
	}
}

// NewVoidTaxParamsWithHTTPClient creates a new VoidTaxParams object
// with the ability to set a custom HTTPClient for a request.
func NewVoidTaxParamsWithHTTPClient(client *http.Client) *VoidTaxParams {
	return &VoidTaxParams{
		HTTPClient: client,
	}
}

/*
VoidTaxParams contains all the parameters to send to the API endpoint

	for the void tax operation.

	Typically these are written to a http.Request.
*/
type VoidTaxParams struct {

	/* ID.

	   The tax ID returned from a previous request.
	*/
	ID string

	// VoidTaxRequest.
	VoidTaxRequest VoidTaxBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the void tax params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoidTaxParams) WithDefaults() *VoidTaxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the void tax params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoidTaxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the void tax params
func (o *VoidTaxParams) WithTimeout(timeout time.Duration) *VoidTaxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void tax params
func (o *VoidTaxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void tax params
func (o *VoidTaxParams) WithContext(ctx context.Context) *VoidTaxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void tax params
func (o *VoidTaxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void tax params
func (o *VoidTaxParams) WithHTTPClient(client *http.Client) *VoidTaxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void tax params
func (o *VoidTaxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the void tax params
func (o *VoidTaxParams) WithID(id string) *VoidTaxParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void tax params
func (o *VoidTaxParams) SetID(id string) {
	o.ID = id
}

// WithVoidTaxRequest adds the voidTaxRequest to the void tax params
func (o *VoidTaxParams) WithVoidTaxRequest(voidTaxRequest VoidTaxBody) *VoidTaxParams {
	o.SetVoidTaxRequest(voidTaxRequest)
	return o
}

// SetVoidTaxRequest adds the voidTaxRequest to the void tax params
func (o *VoidTaxParams) SetVoidTaxRequest(voidTaxRequest VoidTaxBody) {
	o.VoidTaxRequest = voidTaxRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VoidTaxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.VoidTaxRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
