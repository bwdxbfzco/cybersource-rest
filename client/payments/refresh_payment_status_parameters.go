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

// NewRefreshPaymentStatusParams creates a new RefreshPaymentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefreshPaymentStatusParams() *RefreshPaymentStatusParams {
	return &RefreshPaymentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshPaymentStatusParamsWithTimeout creates a new RefreshPaymentStatusParams object
// with the ability to set a timeout on a request.
func NewRefreshPaymentStatusParamsWithTimeout(timeout time.Duration) *RefreshPaymentStatusParams {
	return &RefreshPaymentStatusParams{
		timeout: timeout,
	}
}

// NewRefreshPaymentStatusParamsWithContext creates a new RefreshPaymentStatusParams object
// with the ability to set a context for a request.
func NewRefreshPaymentStatusParamsWithContext(ctx context.Context) *RefreshPaymentStatusParams {
	return &RefreshPaymentStatusParams{
		Context: ctx,
	}
}

// NewRefreshPaymentStatusParamsWithHTTPClient creates a new RefreshPaymentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefreshPaymentStatusParamsWithHTTPClient(client *http.Client) *RefreshPaymentStatusParams {
	return &RefreshPaymentStatusParams{
		HTTPClient: client,
	}
}

/*
RefreshPaymentStatusParams contains all the parameters to send to the API endpoint

	for the refresh payment status operation.

	Typically these are written to a http.Request.
*/
type RefreshPaymentStatusParams struct {

	/* ID.

	   The payment id whose status needs to be checked and updated.
	*/
	ID string

	// RefreshPaymentStatusRequest.
	RefreshPaymentStatusRequest RefreshPaymentStatusBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refresh payment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshPaymentStatusParams) WithDefaults() *RefreshPaymentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refresh payment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshPaymentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the refresh payment status params
func (o *RefreshPaymentStatusParams) WithTimeout(timeout time.Duration) *RefreshPaymentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh payment status params
func (o *RefreshPaymentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh payment status params
func (o *RefreshPaymentStatusParams) WithContext(ctx context.Context) *RefreshPaymentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh payment status params
func (o *RefreshPaymentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh payment status params
func (o *RefreshPaymentStatusParams) WithHTTPClient(client *http.Client) *RefreshPaymentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh payment status params
func (o *RefreshPaymentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the refresh payment status params
func (o *RefreshPaymentStatusParams) WithID(id string) *RefreshPaymentStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the refresh payment status params
func (o *RefreshPaymentStatusParams) SetID(id string) {
	o.ID = id
}

// WithRefreshPaymentStatusRequest adds the refreshPaymentStatusRequest to the refresh payment status params
func (o *RefreshPaymentStatusParams) WithRefreshPaymentStatusRequest(refreshPaymentStatusRequest RefreshPaymentStatusBody) *RefreshPaymentStatusParams {
	o.SetRefreshPaymentStatusRequest(refreshPaymentStatusRequest)
	return o
}

// SetRefreshPaymentStatusRequest adds the refreshPaymentStatusRequest to the refresh payment status params
func (o *RefreshPaymentStatusParams) SetRefreshPaymentStatusRequest(refreshPaymentStatusRequest RefreshPaymentStatusBody) {
	o.RefreshPaymentStatusRequest = refreshPaymentStatusRequest
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshPaymentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.RefreshPaymentStatusRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
