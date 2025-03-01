// Code generated by go-swagger; DO NOT EDIT.

package decision_manager

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

// NewActionDecisionManagerCaseParams creates a new ActionDecisionManagerCaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionDecisionManagerCaseParams() *ActionDecisionManagerCaseParams {
	return &ActionDecisionManagerCaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionDecisionManagerCaseParamsWithTimeout creates a new ActionDecisionManagerCaseParams object
// with the ability to set a timeout on a request.
func NewActionDecisionManagerCaseParamsWithTimeout(timeout time.Duration) *ActionDecisionManagerCaseParams {
	return &ActionDecisionManagerCaseParams{
		timeout: timeout,
	}
}

// NewActionDecisionManagerCaseParamsWithContext creates a new ActionDecisionManagerCaseParams object
// with the ability to set a context for a request.
func NewActionDecisionManagerCaseParamsWithContext(ctx context.Context) *ActionDecisionManagerCaseParams {
	return &ActionDecisionManagerCaseParams{
		Context: ctx,
	}
}

// NewActionDecisionManagerCaseParamsWithHTTPClient creates a new ActionDecisionManagerCaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionDecisionManagerCaseParamsWithHTTPClient(client *http.Client) *ActionDecisionManagerCaseParams {
	return &ActionDecisionManagerCaseParams{
		HTTPClient: client,
	}
}

/*
ActionDecisionManagerCaseParams contains all the parameters to send to the API endpoint

	for the action decision manager case operation.

	Typically these are written to a http.Request.
*/
type ActionDecisionManagerCaseParams struct {

	// CaseManagementActionsRequest.
	CaseManagementActionsRequest ActionDecisionManagerCaseBody

	/* ID.

	   An unique identification number generated by Cybersource to identify the submitted request.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action decision manager case params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionDecisionManagerCaseParams) WithDefaults() *ActionDecisionManagerCaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action decision manager case params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionDecisionManagerCaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) WithTimeout(timeout time.Duration) *ActionDecisionManagerCaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) WithContext(ctx context.Context) *ActionDecisionManagerCaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) WithHTTPClient(client *http.Client) *ActionDecisionManagerCaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaseManagementActionsRequest adds the caseManagementActionsRequest to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) WithCaseManagementActionsRequest(caseManagementActionsRequest ActionDecisionManagerCaseBody) *ActionDecisionManagerCaseParams {
	o.SetCaseManagementActionsRequest(caseManagementActionsRequest)
	return o
}

// SetCaseManagementActionsRequest adds the caseManagementActionsRequest to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) SetCaseManagementActionsRequest(caseManagementActionsRequest ActionDecisionManagerCaseBody) {
	o.CaseManagementActionsRequest = caseManagementActionsRequest
}

// WithID adds the id to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) WithID(id string) *ActionDecisionManagerCaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the action decision manager case params
func (o *ActionDecisionManagerCaseParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ActionDecisionManagerCaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CaseManagementActionsRequest); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
