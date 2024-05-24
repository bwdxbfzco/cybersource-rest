// Code generated by go-swagger; DO NOT EDIT.

package e_m_v_tag_details

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

// NewParseEmvTagsParams creates a new ParseEmvTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewParseEmvTagsParams() *ParseEmvTagsParams {
	return &ParseEmvTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewParseEmvTagsParamsWithTimeout creates a new ParseEmvTagsParams object
// with the ability to set a timeout on a request.
func NewParseEmvTagsParamsWithTimeout(timeout time.Duration) *ParseEmvTagsParams {
	return &ParseEmvTagsParams{
		timeout: timeout,
	}
}

// NewParseEmvTagsParamsWithContext creates a new ParseEmvTagsParams object
// with the ability to set a context for a request.
func NewParseEmvTagsParamsWithContext(ctx context.Context) *ParseEmvTagsParams {
	return &ParseEmvTagsParams{
		Context: ctx,
	}
}

// NewParseEmvTagsParamsWithHTTPClient creates a new ParseEmvTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewParseEmvTagsParamsWithHTTPClient(client *http.Client) *ParseEmvTagsParams {
	return &ParseEmvTagsParams{
		HTTPClient: client,
	}
}

/*
ParseEmvTagsParams contains all the parameters to send to the API endpoint

	for the parse emv tags operation.

	Typically these are written to a http.Request.
*/
type ParseEmvTagsParams struct {

	// Body.
	Body ParseEmvTagsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the parse emv tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParseEmvTagsParams) WithDefaults() *ParseEmvTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the parse emv tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParseEmvTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the parse emv tags params
func (o *ParseEmvTagsParams) WithTimeout(timeout time.Duration) *ParseEmvTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the parse emv tags params
func (o *ParseEmvTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the parse emv tags params
func (o *ParseEmvTagsParams) WithContext(ctx context.Context) *ParseEmvTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the parse emv tags params
func (o *ParseEmvTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the parse emv tags params
func (o *ParseEmvTagsParams) WithHTTPClient(client *http.Client) *ParseEmvTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the parse emv tags params
func (o *ParseEmvTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the parse emv tags params
func (o *ParseEmvTagsParams) WithBody(body ParseEmvTagsBody) *ParseEmvTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the parse emv tags params
func (o *ParseEmvTagsParams) SetBody(body ParseEmvTagsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ParseEmvTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
