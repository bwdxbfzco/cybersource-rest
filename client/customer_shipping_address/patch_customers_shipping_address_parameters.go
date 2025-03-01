// Code generated by go-swagger; DO NOT EDIT.

package customer_shipping_address

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

// NewPatchCustomersShippingAddressParams creates a new PatchCustomersShippingAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCustomersShippingAddressParams() *PatchCustomersShippingAddressParams {
	return &PatchCustomersShippingAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCustomersShippingAddressParamsWithTimeout creates a new PatchCustomersShippingAddressParams object
// with the ability to set a timeout on a request.
func NewPatchCustomersShippingAddressParamsWithTimeout(timeout time.Duration) *PatchCustomersShippingAddressParams {
	return &PatchCustomersShippingAddressParams{
		timeout: timeout,
	}
}

// NewPatchCustomersShippingAddressParamsWithContext creates a new PatchCustomersShippingAddressParams object
// with the ability to set a context for a request.
func NewPatchCustomersShippingAddressParamsWithContext(ctx context.Context) *PatchCustomersShippingAddressParams {
	return &PatchCustomersShippingAddressParams{
		Context: ctx,
	}
}

// NewPatchCustomersShippingAddressParamsWithHTTPClient creates a new PatchCustomersShippingAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCustomersShippingAddressParamsWithHTTPClient(client *http.Client) *PatchCustomersShippingAddressParams {
	return &PatchCustomersShippingAddressParams{
		HTTPClient: client,
	}
}

/*
PatchCustomersShippingAddressParams contains all the parameters to send to the API endpoint

	for the patch customers shipping address operation.

	Typically these are written to a http.Request.
*/
type PatchCustomersShippingAddressParams struct {

	/* CustomerID.

	   The Id of a Customer.
	*/
	CustomerID string

	/* IfMatch.

	   Contains an ETag value from a GET request to make the request conditional.
	*/
	IfMatch *string

	// PatchCustomerShippingAddressRequest.
	PatchCustomerShippingAddressRequest PatchCustomersShippingAddressBody

	/* ProfileID.

	   The Id of a profile containing user specific TMS configuration.
	*/
	ProfileID *string

	/* ShippingAddressID.

	   The Id of a shipping address.
	*/
	ShippingAddressID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch customers shipping address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCustomersShippingAddressParams) WithDefaults() *PatchCustomersShippingAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch customers shipping address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCustomersShippingAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithTimeout(timeout time.Duration) *PatchCustomersShippingAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithContext(ctx context.Context) *PatchCustomersShippingAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithHTTPClient(client *http.Client) *PatchCustomersShippingAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithCustomerID(customerID string) *PatchCustomersShippingAddressParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithIfMatch adds the ifMatch to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithIfMatch(ifMatch *string) *PatchCustomersShippingAddressParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithPatchCustomerShippingAddressRequest adds the patchCustomerShippingAddressRequest to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithPatchCustomerShippingAddressRequest(patchCustomerShippingAddressRequest PatchCustomersShippingAddressBody) *PatchCustomersShippingAddressParams {
	o.SetPatchCustomerShippingAddressRequest(patchCustomerShippingAddressRequest)
	return o
}

// SetPatchCustomerShippingAddressRequest adds the patchCustomerShippingAddressRequest to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetPatchCustomerShippingAddressRequest(patchCustomerShippingAddressRequest PatchCustomersShippingAddressBody) {
	o.PatchCustomerShippingAddressRequest = patchCustomerShippingAddressRequest
}

// WithProfileID adds the profileID to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithProfileID(profileID *string) *PatchCustomersShippingAddressParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WithShippingAddressID adds the shippingAddressID to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) WithShippingAddressID(shippingAddressID string) *PatchCustomersShippingAddressParams {
	o.SetShippingAddressID(shippingAddressID)
	return o
}

// SetShippingAddressID adds the shippingAddressId to the patch customers shipping address params
func (o *PatchCustomersShippingAddressParams) SetShippingAddressID(shippingAddressID string) {
	o.ShippingAddressID = shippingAddressID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCustomersShippingAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.PatchCustomerShippingAddressRequest); err != nil {
		return err
	}

	if o.ProfileID != nil {

		// header param profile-id
		if err := r.SetHeaderParam("profile-id", *o.ProfileID); err != nil {
			return err
		}
	}

	// path param shippingAddressId
	if err := r.SetPathParam("shippingAddressId", o.ShippingAddressID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
