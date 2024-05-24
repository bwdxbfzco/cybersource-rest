// Code generated by go-swagger; DO NOT EDIT.

package user_management_search

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

// NewSearchUsersParams creates a new SearchUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUsersParams() *SearchUsersParams {
	return &SearchUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUsersParamsWithTimeout creates a new SearchUsersParams object
// with the ability to set a timeout on a request.
func NewSearchUsersParamsWithTimeout(timeout time.Duration) *SearchUsersParams {
	return &SearchUsersParams{
		timeout: timeout,
	}
}

// NewSearchUsersParamsWithContext creates a new SearchUsersParams object
// with the ability to set a context for a request.
func NewSearchUsersParamsWithContext(ctx context.Context) *SearchUsersParams {
	return &SearchUsersParams{
		Context: ctx,
	}
}

// NewSearchUsersParamsWithHTTPClient creates a new SearchUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUsersParamsWithHTTPClient(client *http.Client) *SearchUsersParams {
	return &SearchUsersParams{
		HTTPClient: client,
	}
}

/*
SearchUsersParams contains all the parameters to send to the API endpoint

	for the search users operation.

	Typically these are written to a http.Request.
*/
type SearchUsersParams struct {

	// SearchRequest.
	SearchRequest SearchUsersBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUsersParams) WithDefaults() *SearchUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search users params
func (o *SearchUsersParams) WithTimeout(timeout time.Duration) *SearchUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search users params
func (o *SearchUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search users params
func (o *SearchUsersParams) WithContext(ctx context.Context) *SearchUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search users params
func (o *SearchUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) WithHTTPClient(client *http.Client) *SearchUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchRequest adds the searchRequest to the search users params
func (o *SearchUsersParams) WithSearchRequest(searchRequest SearchUsersBody) *SearchUsersParams {
	o.SetSearchRequest(searchRequest)
	return o
}

// SetSearchRequest adds the searchRequest to the search users params
func (o *SearchUsersParams) SetSearchRequest(searchRequest SearchUsersBody) {
	o.SearchRequest = searchRequest
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.SearchRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
