// Code generated by go-swagger; DO NOT EDIT.

package report_subscriptions

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

// NewDeleteSubscriptionParams creates a new DeleteSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSubscriptionParams() *DeleteSubscriptionParams {
	return &DeleteSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubscriptionParamsWithTimeout creates a new DeleteSubscriptionParams object
// with the ability to set a timeout on a request.
func NewDeleteSubscriptionParamsWithTimeout(timeout time.Duration) *DeleteSubscriptionParams {
	return &DeleteSubscriptionParams{
		timeout: timeout,
	}
}

// NewDeleteSubscriptionParamsWithContext creates a new DeleteSubscriptionParams object
// with the ability to set a context for a request.
func NewDeleteSubscriptionParamsWithContext(ctx context.Context) *DeleteSubscriptionParams {
	return &DeleteSubscriptionParams{
		Context: ctx,
	}
}

// NewDeleteSubscriptionParamsWithHTTPClient creates a new DeleteSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSubscriptionParamsWithHTTPClient(client *http.Client) *DeleteSubscriptionParams {
	return &DeleteSubscriptionParams{
		HTTPClient: client,
	}
}

/*
DeleteSubscriptionParams contains all the parameters to send to the API endpoint

	for the delete subscription operation.

	Typically these are written to a http.Request.
*/
type DeleteSubscriptionParams struct {

	/* OrganizationID.

	   Valid Organization Id
	*/
	OrganizationID *string

	/* ReportName.

	   Name of the Report to Delete
	*/
	ReportName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubscriptionParams) WithDefaults() *DeleteSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete subscription params
func (o *DeleteSubscriptionParams) WithTimeout(timeout time.Duration) *DeleteSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subscription params
func (o *DeleteSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subscription params
func (o *DeleteSubscriptionParams) WithContext(ctx context.Context) *DeleteSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subscription params
func (o *DeleteSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subscription params
func (o *DeleteSubscriptionParams) WithHTTPClient(client *http.Client) *DeleteSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subscription params
func (o *DeleteSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the delete subscription params
func (o *DeleteSubscriptionParams) WithOrganizationID(organizationID *string) *DeleteSubscriptionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete subscription params
func (o *DeleteSubscriptionParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithReportName adds the reportName to the delete subscription params
func (o *DeleteSubscriptionParams) WithReportName(reportName string) *DeleteSubscriptionParams {
	o.SetReportName(reportName)
	return o
}

// SetReportName adds the reportName to the delete subscription params
func (o *DeleteSubscriptionParams) SetReportName(reportName string) {
	o.ReportName = reportName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}
	}

	// path param reportName
	if err := r.SetPathParam("reportName", o.ReportName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
