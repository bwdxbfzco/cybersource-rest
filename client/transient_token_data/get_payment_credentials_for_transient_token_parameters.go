// Code generated by go-swagger; DO NOT EDIT.

package transient_token_data

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

// NewGetPaymentCredentialsForTransientTokenParams creates a new GetPaymentCredentialsForTransientTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPaymentCredentialsForTransientTokenParams() *GetPaymentCredentialsForTransientTokenParams {
	return &GetPaymentCredentialsForTransientTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentCredentialsForTransientTokenParamsWithTimeout creates a new GetPaymentCredentialsForTransientTokenParams object
// with the ability to set a timeout on a request.
func NewGetPaymentCredentialsForTransientTokenParamsWithTimeout(timeout time.Duration) *GetPaymentCredentialsForTransientTokenParams {
	return &GetPaymentCredentialsForTransientTokenParams{
		timeout: timeout,
	}
}

// NewGetPaymentCredentialsForTransientTokenParamsWithContext creates a new GetPaymentCredentialsForTransientTokenParams object
// with the ability to set a context for a request.
func NewGetPaymentCredentialsForTransientTokenParamsWithContext(ctx context.Context) *GetPaymentCredentialsForTransientTokenParams {
	return &GetPaymentCredentialsForTransientTokenParams{
		Context: ctx,
	}
}

// NewGetPaymentCredentialsForTransientTokenParamsWithHTTPClient creates a new GetPaymentCredentialsForTransientTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPaymentCredentialsForTransientTokenParamsWithHTTPClient(client *http.Client) *GetPaymentCredentialsForTransientTokenParams {
	return &GetPaymentCredentialsForTransientTokenParams{
		HTTPClient: client,
	}
}

/*
GetPaymentCredentialsForTransientTokenParams contains all the parameters to send to the API endpoint

	for the get payment credentials for transient token operation.

	Typically these are written to a http.Request.
*/
type GetPaymentCredentialsForTransientTokenParams struct {

	/* Jti.

	   The jti field contained within the Transient token returned from a successful Unified Checkout transaction

	*/
	Jti string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get payment credentials for transient token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentCredentialsForTransientTokenParams) WithDefaults() *GetPaymentCredentialsForTransientTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payment credentials for transient token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentCredentialsForTransientTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) WithTimeout(timeout time.Duration) *GetPaymentCredentialsForTransientTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) WithContext(ctx context.Context) *GetPaymentCredentialsForTransientTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) WithHTTPClient(client *http.Client) *GetPaymentCredentialsForTransientTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJti adds the jti to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) WithJti(jti string) *GetPaymentCredentialsForTransientTokenParams {
	o.SetJti(jti)
	return o
}

// SetJti adds the jti to the get payment credentials for transient token params
func (o *GetPaymentCredentialsForTransientTokenParams) SetJti(jti string) {
	o.Jti = jti
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentCredentialsForTransientTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jti
	if err := r.SetPathParam("jti", o.Jti); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
