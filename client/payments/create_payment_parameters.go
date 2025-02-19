// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// NewCreatePaymentParams creates a new CreatePaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePaymentParams() *CreatePaymentParams {
	return &CreatePaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentParamsWithTimeout creates a new CreatePaymentParams object
// with the ability to set a timeout on a request.
func NewCreatePaymentParamsWithTimeout(timeout time.Duration) *CreatePaymentParams {
	return &CreatePaymentParams{
		timeout: timeout,
	}
}

// NewCreatePaymentParamsWithContext creates a new CreatePaymentParams object
// with the ability to set a context for a request.
func NewCreatePaymentParamsWithContext(ctx context.Context) *CreatePaymentParams {
	return &CreatePaymentParams{
		Context: ctx,
	}
}

// NewCreatePaymentParamsWithHTTPClient creates a new CreatePaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePaymentParamsWithHTTPClient(client *http.Client) *CreatePaymentParams {
	return &CreatePaymentParams{
		HTTPClient: client,
	}
}

/*
CreatePaymentParams contains all the parameters to send to the API endpoint

	for the create payment operation.

	Typically these are written to a http.Request.
*/
type CreatePaymentParams struct {

	// CreatePaymentRequest.
	CreatePaymentRequest CreatePaymentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentParams) WithDefaults() *CreatePaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create payment params
func (o *CreatePaymentParams) WithTimeout(timeout time.Duration) *CreatePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment params
func (o *CreatePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment params
func (o *CreatePaymentParams) WithContext(ctx context.Context) *CreatePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment params
func (o *CreatePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment params
func (o *CreatePaymentParams) WithHTTPClient(client *http.Client) *CreatePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment params
func (o *CreatePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatePaymentRequest adds the createPaymentRequest to the create payment params
func (o *CreatePaymentParams) WithCreatePaymentRequest(createPaymentRequest CreatePaymentBody) *CreatePaymentParams {
	o.SetCreatePaymentRequest(createPaymentRequest)
	return o
}

// SetCreatePaymentRequest adds the createPaymentRequest to the create payment params
func (o *CreatePaymentParams) SetCreatePaymentRequest(createPaymentRequest CreatePaymentBody) {
	o.CreatePaymentRequest = createPaymentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreatePaymentRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
