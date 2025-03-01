// Code generated by go-swagger; DO NOT EDIT.

package batches

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

// NewGetBatchStatusParams creates a new GetBatchStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBatchStatusParams() *GetBatchStatusParams {
	return &GetBatchStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchStatusParamsWithTimeout creates a new GetBatchStatusParams object
// with the ability to set a timeout on a request.
func NewGetBatchStatusParamsWithTimeout(timeout time.Duration) *GetBatchStatusParams {
	return &GetBatchStatusParams{
		timeout: timeout,
	}
}

// NewGetBatchStatusParamsWithContext creates a new GetBatchStatusParams object
// with the ability to set a context for a request.
func NewGetBatchStatusParamsWithContext(ctx context.Context) *GetBatchStatusParams {
	return &GetBatchStatusParams{
		Context: ctx,
	}
}

// NewGetBatchStatusParamsWithHTTPClient creates a new GetBatchStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBatchStatusParamsWithHTTPClient(client *http.Client) *GetBatchStatusParams {
	return &GetBatchStatusParams{
		HTTPClient: client,
	}
}

/*
GetBatchStatusParams contains all the parameters to send to the API endpoint

	for the get batch status operation.

	Typically these are written to a http.Request.
*/
type GetBatchStatusParams struct {

	/* BatchID.

	   Unique identification number assigned to the submitted request.
	*/
	BatchID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get batch status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchStatusParams) WithDefaults() *GetBatchStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get batch status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get batch status params
func (o *GetBatchStatusParams) WithTimeout(timeout time.Duration) *GetBatchStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batch status params
func (o *GetBatchStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batch status params
func (o *GetBatchStatusParams) WithContext(ctx context.Context) *GetBatchStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batch status params
func (o *GetBatchStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batch status params
func (o *GetBatchStatusParams) WithHTTPClient(client *http.Client) *GetBatchStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batch status params
func (o *GetBatchStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchID adds the batchID to the get batch status params
func (o *GetBatchStatusParams) WithBatchID(batchID string) *GetBatchStatusParams {
	o.SetBatchID(batchID)
	return o
}

// SetBatchID adds the batchId to the get batch status params
func (o *GetBatchStatusParams) SetBatchID(batchID string) {
	o.BatchID = batchID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param batchId
	if err := r.SetPathParam("batchId", o.BatchID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
