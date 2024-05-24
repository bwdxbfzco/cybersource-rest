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

// NewUpdatePlanParams creates a new UpdatePlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePlanParams() *UpdatePlanParams {
	return &UpdatePlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePlanParamsWithTimeout creates a new UpdatePlanParams object
// with the ability to set a timeout on a request.
func NewUpdatePlanParamsWithTimeout(timeout time.Duration) *UpdatePlanParams {
	return &UpdatePlanParams{
		timeout: timeout,
	}
}

// NewUpdatePlanParamsWithContext creates a new UpdatePlanParams object
// with the ability to set a context for a request.
func NewUpdatePlanParamsWithContext(ctx context.Context) *UpdatePlanParams {
	return &UpdatePlanParams{
		Context: ctx,
	}
}

// NewUpdatePlanParamsWithHTTPClient creates a new UpdatePlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePlanParamsWithHTTPClient(client *http.Client) *UpdatePlanParams {
	return &UpdatePlanParams{
		HTTPClient: client,
	}
}

/*
UpdatePlanParams contains all the parameters to send to the API endpoint

	for the update plan operation.

	Typically these are written to a http.Request.
*/
type UpdatePlanParams struct {

	/* ID.

	   Plan Id
	*/
	ID string

	// UpdatePlanRequest.
	UpdatePlanRequest UpdatePlanBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePlanParams) WithDefaults() *UpdatePlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update plan params
func (o *UpdatePlanParams) WithTimeout(timeout time.Duration) *UpdatePlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update plan params
func (o *UpdatePlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update plan params
func (o *UpdatePlanParams) WithContext(ctx context.Context) *UpdatePlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update plan params
func (o *UpdatePlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update plan params
func (o *UpdatePlanParams) WithHTTPClient(client *http.Client) *UpdatePlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update plan params
func (o *UpdatePlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update plan params
func (o *UpdatePlanParams) WithID(id string) *UpdatePlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update plan params
func (o *UpdatePlanParams) SetID(id string) {
	o.ID = id
}

// WithUpdatePlanRequest adds the updatePlanRequest to the update plan params
func (o *UpdatePlanParams) WithUpdatePlanRequest(updatePlanRequest UpdatePlanBody) *UpdatePlanParams {
	o.SetUpdatePlanRequest(updatePlanRequest)
	return o
}

// SetUpdatePlanRequest adds the updatePlanRequest to the update plan params
func (o *UpdatePlanParams) SetUpdatePlanRequest(updatePlanRequest UpdatePlanBody) {
	o.UpdatePlanRequest = updatePlanRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdatePlanRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
