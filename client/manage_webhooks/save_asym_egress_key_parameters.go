// Code generated by go-swagger; DO NOT EDIT.

package manage_webhooks

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

// NewSaveAsymEgressKeyParams creates a new SaveAsymEgressKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveAsymEgressKeyParams() *SaveAsymEgressKeyParams {
	return &SaveAsymEgressKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveAsymEgressKeyParamsWithTimeout creates a new SaveAsymEgressKeyParams object
// with the ability to set a timeout on a request.
func NewSaveAsymEgressKeyParamsWithTimeout(timeout time.Duration) *SaveAsymEgressKeyParams {
	return &SaveAsymEgressKeyParams{
		timeout: timeout,
	}
}

// NewSaveAsymEgressKeyParamsWithContext creates a new SaveAsymEgressKeyParams object
// with the ability to set a context for a request.
func NewSaveAsymEgressKeyParamsWithContext(ctx context.Context) *SaveAsymEgressKeyParams {
	return &SaveAsymEgressKeyParams{
		Context: ctx,
	}
}

// NewSaveAsymEgressKeyParamsWithHTTPClient creates a new SaveAsymEgressKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveAsymEgressKeyParamsWithHTTPClient(client *http.Client) *SaveAsymEgressKeyParams {
	return &SaveAsymEgressKeyParams{
		HTTPClient: client,
	}
}

/*
SaveAsymEgressKeyParams contains all the parameters to send to the API endpoint

	for the save asym egress key operation.

	Typically these are written to a http.Request.
*/
type SaveAsymEgressKeyParams struct {

	/* SaveAsymEgressKey.

	   Provide egress Asymmetric key information to save (create or store)
	*/
	SaveAsymEgressKey SaveAsymEgressKeyBody

	/* VcCorrelationID.

	   A globally unique id associated with your request
	*/
	VcCorrelationID *string

	/* VcPermissions.

	   Encoded user permissions returned by the CGK, for the entity user who initiated the boarding
	*/
	VcPermissions string

	/* VcSenderOrganizationID.

	   Sender organization id
	*/
	VcSenderOrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save asym egress key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveAsymEgressKeyParams) WithDefaults() *SaveAsymEgressKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save asym egress key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveAsymEgressKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithTimeout(timeout time.Duration) *SaveAsymEgressKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithContext(ctx context.Context) *SaveAsymEgressKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithHTTPClient(client *http.Client) *SaveAsymEgressKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSaveAsymEgressKey adds the saveAsymEgressKey to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithSaveAsymEgressKey(saveAsymEgressKey SaveAsymEgressKeyBody) *SaveAsymEgressKeyParams {
	o.SetSaveAsymEgressKey(saveAsymEgressKey)
	return o
}

// SetSaveAsymEgressKey adds the saveAsymEgressKey to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetSaveAsymEgressKey(saveAsymEgressKey SaveAsymEgressKeyBody) {
	o.SaveAsymEgressKey = saveAsymEgressKey
}

// WithVcCorrelationID adds the vcCorrelationID to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithVcCorrelationID(vcCorrelationID *string) *SaveAsymEgressKeyParams {
	o.SetVcCorrelationID(vcCorrelationID)
	return o
}

// SetVcCorrelationID adds the vCCorrelationId to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetVcCorrelationID(vcCorrelationID *string) {
	o.VcCorrelationID = vcCorrelationID
}

// WithVcPermissions adds the vcPermissions to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithVcPermissions(vcPermissions string) *SaveAsymEgressKeyParams {
	o.SetVcPermissions(vcPermissions)
	return o
}

// SetVcPermissions adds the vCPermissions to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetVcPermissions(vcPermissions string) {
	o.VcPermissions = vcPermissions
}

// WithVcSenderOrganizationID adds the vcSenderOrganizationID to the save asym egress key params
func (o *SaveAsymEgressKeyParams) WithVcSenderOrganizationID(vcSenderOrganizationID string) *SaveAsymEgressKeyParams {
	o.SetVcSenderOrganizationID(vcSenderOrganizationID)
	return o
}

// SetVcSenderOrganizationID adds the vCSenderOrganizationId to the save asym egress key params
func (o *SaveAsymEgressKeyParams) SetVcSenderOrganizationID(vcSenderOrganizationID string) {
	o.VcSenderOrganizationID = vcSenderOrganizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SaveAsymEgressKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.SaveAsymEgressKey); err != nil {
		return err
	}

	if o.VcCorrelationID != nil {

		// header param v-c-correlation-id
		if err := r.SetHeaderParam("v-c-correlation-id", *o.VcCorrelationID); err != nil {
			return err
		}
	}

	// header param v-c-permissions
	if err := r.SetHeaderParam("v-c-permissions", o.VcPermissions); err != nil {
		return err
	}

	// header param v-c-sender-organization-id
	if err := r.SetHeaderParam("v-c-sender-organization-id", o.VcSenderOrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
