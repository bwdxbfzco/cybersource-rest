// Code generated by go-swagger; DO NOT EDIT.

package plans

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

// NewGetPlanCodeParams creates a new GetPlanCodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlanCodeParams() *GetPlanCodeParams {
	return &GetPlanCodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlanCodeParamsWithTimeout creates a new GetPlanCodeParams object
// with the ability to set a timeout on a request.
func NewGetPlanCodeParamsWithTimeout(timeout time.Duration) *GetPlanCodeParams {
	return &GetPlanCodeParams{
		timeout: timeout,
	}
}

// NewGetPlanCodeParamsWithContext creates a new GetPlanCodeParams object
// with the ability to set a context for a request.
func NewGetPlanCodeParamsWithContext(ctx context.Context) *GetPlanCodeParams {
	return &GetPlanCodeParams{
		Context: ctx,
	}
}

// NewGetPlanCodeParamsWithHTTPClient creates a new GetPlanCodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlanCodeParamsWithHTTPClient(client *http.Client) *GetPlanCodeParams {
	return &GetPlanCodeParams{
		HTTPClient: client,
	}
}

/*
GetPlanCodeParams contains all the parameters to send to the API endpoint

	for the get plan code operation.

	Typically these are written to a http.Request.
*/
type GetPlanCodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plan code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanCodeParams) WithDefaults() *GetPlanCodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plan code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanCodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get plan code params
func (o *GetPlanCodeParams) WithTimeout(timeout time.Duration) *GetPlanCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plan code params
func (o *GetPlanCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plan code params
func (o *GetPlanCodeParams) WithContext(ctx context.Context) *GetPlanCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plan code params
func (o *GetPlanCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plan code params
func (o *GetPlanCodeParams) WithHTTPClient(client *http.Client) *GetPlanCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plan code params
func (o *GetPlanCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlanCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
