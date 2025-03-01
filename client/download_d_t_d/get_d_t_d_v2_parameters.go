// Code generated by go-swagger; DO NOT EDIT.

package download_d_t_d

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

// NewGetDTDV2Params creates a new GetDTDV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDTDV2Params() *GetDTDV2Params {
	return &GetDTDV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDTDV2ParamsWithTimeout creates a new GetDTDV2Params object
// with the ability to set a timeout on a request.
func NewGetDTDV2ParamsWithTimeout(timeout time.Duration) *GetDTDV2Params {
	return &GetDTDV2Params{
		timeout: timeout,
	}
}

// NewGetDTDV2ParamsWithContext creates a new GetDTDV2Params object
// with the ability to set a context for a request.
func NewGetDTDV2ParamsWithContext(ctx context.Context) *GetDTDV2Params {
	return &GetDTDV2Params{
		Context: ctx,
	}
}

// NewGetDTDV2ParamsWithHTTPClient creates a new GetDTDV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDTDV2ParamsWithHTTPClient(client *http.Client) *GetDTDV2Params {
	return &GetDTDV2Params{
		HTTPClient: client,
	}
}

/*
GetDTDV2Params contains all the parameters to send to the API endpoint

	for the get d t d v2 operation.

	Typically these are written to a http.Request.
*/
type GetDTDV2Params struct {

	/* ReportDefinitionNameVersion.

	   Name and version of DTD file to download. Some DTDs only have one version. In that case version name is not needed. Some example values are ctdr-1.0, tdr, pbdr-1.1
	*/
	ReportDefinitionNameVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get d t d v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDTDV2Params) WithDefaults() *GetDTDV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get d t d v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDTDV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get d t d v2 params
func (o *GetDTDV2Params) WithTimeout(timeout time.Duration) *GetDTDV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get d t d v2 params
func (o *GetDTDV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get d t d v2 params
func (o *GetDTDV2Params) WithContext(ctx context.Context) *GetDTDV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get d t d v2 params
func (o *GetDTDV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get d t d v2 params
func (o *GetDTDV2Params) WithHTTPClient(client *http.Client) *GetDTDV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get d t d v2 params
func (o *GetDTDV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportDefinitionNameVersion adds the reportDefinitionNameVersion to the get d t d v2 params
func (o *GetDTDV2Params) WithReportDefinitionNameVersion(reportDefinitionNameVersion string) *GetDTDV2Params {
	o.SetReportDefinitionNameVersion(reportDefinitionNameVersion)
	return o
}

// SetReportDefinitionNameVersion adds the reportDefinitionNameVersion to the get d t d v2 params
func (o *GetDTDV2Params) SetReportDefinitionNameVersion(reportDefinitionNameVersion string) {
	o.ReportDefinitionNameVersion = reportDefinitionNameVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetDTDV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reportDefinitionNameVersion
	if err := r.SetPathParam("reportDefinitionNameVersion", o.ReportDefinitionNameVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
