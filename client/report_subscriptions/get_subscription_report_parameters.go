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

// NewGetSubscriptionReportParams creates a new GetSubscriptionReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSubscriptionReportParams() *GetSubscriptionReportParams {
	return &GetSubscriptionReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriptionReportParamsWithTimeout creates a new GetSubscriptionReportParams object
// with the ability to set a timeout on a request.
func NewGetSubscriptionReportParamsWithTimeout(timeout time.Duration) *GetSubscriptionReportParams {
	return &GetSubscriptionReportParams{
		timeout: timeout,
	}
}

// NewGetSubscriptionReportParamsWithContext creates a new GetSubscriptionReportParams object
// with the ability to set a context for a request.
func NewGetSubscriptionReportParamsWithContext(ctx context.Context) *GetSubscriptionReportParams {
	return &GetSubscriptionReportParams{
		Context: ctx,
	}
}

// NewGetSubscriptionReportParamsWithHTTPClient creates a new GetSubscriptionReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSubscriptionReportParamsWithHTTPClient(client *http.Client) *GetSubscriptionReportParams {
	return &GetSubscriptionReportParams{
		HTTPClient: client,
	}
}

/*
GetSubscriptionReportParams contains all the parameters to send to the API endpoint

	for the get subscription report operation.

	Typically these are written to a http.Request.
*/
type GetSubscriptionReportParams struct {

	/* OrganizationID.

	   Valid Organization Id
	*/
	OrganizationID *string

	/* ReportName.

	   Name of the Report to Retrieve
	*/
	ReportName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get subscription report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubscriptionReportParams) WithDefaults() *GetSubscriptionReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get subscription report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubscriptionReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get subscription report params
func (o *GetSubscriptionReportParams) WithTimeout(timeout time.Duration) *GetSubscriptionReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscription report params
func (o *GetSubscriptionReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscription report params
func (o *GetSubscriptionReportParams) WithContext(ctx context.Context) *GetSubscriptionReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscription report params
func (o *GetSubscriptionReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscription report params
func (o *GetSubscriptionReportParams) WithHTTPClient(client *http.Client) *GetSubscriptionReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscription report params
func (o *GetSubscriptionReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get subscription report params
func (o *GetSubscriptionReportParams) WithOrganizationID(organizationID *string) *GetSubscriptionReportParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get subscription report params
func (o *GetSubscriptionReportParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithReportName adds the reportName to the get subscription report params
func (o *GetSubscriptionReportParams) WithReportName(reportName string) *GetSubscriptionReportParams {
	o.SetReportName(reportName)
	return o
}

// SetReportName adds the reportName to the get subscription report params
func (o *GetSubscriptionReportParams) SetReportName(reportName string) {
	o.ReportName = reportName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
