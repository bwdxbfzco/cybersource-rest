// Code generated by go-swagger; DO NOT EDIT.

package key_management

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
	"github.com/go-openapi/swag"
)

// NewSearchKeysParams creates a new SearchKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchKeysParams() *SearchKeysParams {
	return &SearchKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchKeysParamsWithTimeout creates a new SearchKeysParams object
// with the ability to set a timeout on a request.
func NewSearchKeysParamsWithTimeout(timeout time.Duration) *SearchKeysParams {
	return &SearchKeysParams{
		timeout: timeout,
	}
}

// NewSearchKeysParamsWithContext creates a new SearchKeysParams object
// with the ability to set a context for a request.
func NewSearchKeysParamsWithContext(ctx context.Context) *SearchKeysParams {
	return &SearchKeysParams{
		Context: ctx,
	}
}

// NewSearchKeysParamsWithHTTPClient creates a new SearchKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchKeysParamsWithHTTPClient(client *http.Client) *SearchKeysParams {
	return &SearchKeysParams{
		HTTPClient: client,
	}
}

/*
SearchKeysParams contains all the parameters to send to the API endpoint

	for the search keys operation.

	Typically these are written to a http.Request.
*/
type SearchKeysParams struct {

	/* ExpirationEndDate.

	   Expiry Filter End Date. When Expiration Date filter is provided, atleast one more filter needs to be provided

	   Format: date-time
	*/
	ExpirationEndDate *strfmt.DateTime

	/* ExpirationStartDate.

	   Expiry Filter Start Date. When Expiration Date filter is provided, atleast one more filter needs to be provided

	   Format: date-time
	*/
	ExpirationStartDate *strfmt.DateTime

	/* KeyIds.

	   List of Key Ids to search. The maximum size of the Key Ids list is 1
	*/
	KeyIds []string

	/* KeyTypes.

	   Key Type, Possible values -  certificate, password, pgp and scmp_api. When Key Type is provided atleast one more filter needs to be provided
	*/
	KeyTypes []string

	/* Limit.

	   This allows you to specify the total number of records to be returned off the resulting list resultset
	*/
	Limit *int64

	/* Offset.

	   This allows you to specify the page offset from the resulting list resultset you want the records to be returned
	*/
	Offset *int64

	/* OrganizationIds.

	   List of Orgaization Ids to search. The maximum size of the organization Ids list is 1. The maximum length of Organization Id is 30.
	*/
	OrganizationIds []string

	/* Sort.

	   This allows you to specify a comma separated list of fields in the order which the resulting list resultset must be sorted.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchKeysParams) WithDefaults() *SearchKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search keys params
func (o *SearchKeysParams) WithTimeout(timeout time.Duration) *SearchKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search keys params
func (o *SearchKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search keys params
func (o *SearchKeysParams) WithContext(ctx context.Context) *SearchKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search keys params
func (o *SearchKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search keys params
func (o *SearchKeysParams) WithHTTPClient(client *http.Client) *SearchKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search keys params
func (o *SearchKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpirationEndDate adds the expirationEndDate to the search keys params
func (o *SearchKeysParams) WithExpirationEndDate(expirationEndDate *strfmt.DateTime) *SearchKeysParams {
	o.SetExpirationEndDate(expirationEndDate)
	return o
}

// SetExpirationEndDate adds the expirationEndDate to the search keys params
func (o *SearchKeysParams) SetExpirationEndDate(expirationEndDate *strfmt.DateTime) {
	o.ExpirationEndDate = expirationEndDate
}

// WithExpirationStartDate adds the expirationStartDate to the search keys params
func (o *SearchKeysParams) WithExpirationStartDate(expirationStartDate *strfmt.DateTime) *SearchKeysParams {
	o.SetExpirationStartDate(expirationStartDate)
	return o
}

// SetExpirationStartDate adds the expirationStartDate to the search keys params
func (o *SearchKeysParams) SetExpirationStartDate(expirationStartDate *strfmt.DateTime) {
	o.ExpirationStartDate = expirationStartDate
}

// WithKeyIds adds the keyIds to the search keys params
func (o *SearchKeysParams) WithKeyIds(keyIds []string) *SearchKeysParams {
	o.SetKeyIds(keyIds)
	return o
}

// SetKeyIds adds the keyIds to the search keys params
func (o *SearchKeysParams) SetKeyIds(keyIds []string) {
	o.KeyIds = keyIds
}

// WithKeyTypes adds the keyTypes to the search keys params
func (o *SearchKeysParams) WithKeyTypes(keyTypes []string) *SearchKeysParams {
	o.SetKeyTypes(keyTypes)
	return o
}

// SetKeyTypes adds the keyTypes to the search keys params
func (o *SearchKeysParams) SetKeyTypes(keyTypes []string) {
	o.KeyTypes = keyTypes
}

// WithLimit adds the limit to the search keys params
func (o *SearchKeysParams) WithLimit(limit *int64) *SearchKeysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search keys params
func (o *SearchKeysParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search keys params
func (o *SearchKeysParams) WithOffset(offset *int64) *SearchKeysParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search keys params
func (o *SearchKeysParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrganizationIds adds the organizationIds to the search keys params
func (o *SearchKeysParams) WithOrganizationIds(organizationIds []string) *SearchKeysParams {
	o.SetOrganizationIds(organizationIds)
	return o
}

// SetOrganizationIds adds the organizationIds to the search keys params
func (o *SearchKeysParams) SetOrganizationIds(organizationIds []string) {
	o.OrganizationIds = organizationIds
}

// WithSort adds the sort to the search keys params
func (o *SearchKeysParams) WithSort(sort *string) *SearchKeysParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search keys params
func (o *SearchKeysParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExpirationEndDate != nil {

		// query param expirationEndDate
		var qrExpirationEndDate strfmt.DateTime

		if o.ExpirationEndDate != nil {
			qrExpirationEndDate = *o.ExpirationEndDate
		}
		qExpirationEndDate := qrExpirationEndDate.String()
		if qExpirationEndDate != "" {

			if err := r.SetQueryParam("expirationEndDate", qExpirationEndDate); err != nil {
				return err
			}
		}
	}

	if o.ExpirationStartDate != nil {

		// query param expirationStartDate
		var qrExpirationStartDate strfmt.DateTime

		if o.ExpirationStartDate != nil {
			qrExpirationStartDate = *o.ExpirationStartDate
		}
		qExpirationStartDate := qrExpirationStartDate.String()
		if qExpirationStartDate != "" {

			if err := r.SetQueryParam("expirationStartDate", qExpirationStartDate); err != nil {
				return err
			}
		}
	}

	if o.KeyIds != nil {

		// binding items for keyIds
		joinedKeyIds := o.bindParamKeyIds(reg)

		// query array param keyIds
		if err := r.SetQueryParam("keyIds", joinedKeyIds...); err != nil {
			return err
		}
	}

	if o.KeyTypes != nil {

		// binding items for keyTypes
		joinedKeyTypes := o.bindParamKeyTypes(reg)

		// query array param keyTypes
		if err := r.SetQueryParam("keyTypes", joinedKeyTypes...); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrganizationIds != nil {

		// binding items for organizationIds
		joinedOrganizationIds := o.bindParamOrganizationIds(reg)

		// query array param organizationIds
		if err := r.SetQueryParam("organizationIds", joinedOrganizationIds...); err != nil {
			return err
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchKeys binds the parameter keyIds
func (o *SearchKeysParams) bindParamKeyIds(formats strfmt.Registry) []string {
	keyIdsIR := o.KeyIds

	var keyIdsIC []string
	for _, keyIdsIIR := range keyIdsIR { // explode []string

		keyIdsIIV := keyIdsIIR // string as string
		keyIdsIC = append(keyIdsIC, keyIdsIIV)
	}

	// items.CollectionFormat: ""
	keyIdsIS := swag.JoinByFormat(keyIdsIC, "")

	return keyIdsIS
}

// bindParamSearchKeys binds the parameter keyTypes
func (o *SearchKeysParams) bindParamKeyTypes(formats strfmt.Registry) []string {
	keyTypesIR := o.KeyTypes

	var keyTypesIC []string
	for _, keyTypesIIR := range keyTypesIR { // explode []string

		keyTypesIIV := keyTypesIIR // string as string
		keyTypesIC = append(keyTypesIC, keyTypesIIV)
	}

	// items.CollectionFormat: ""
	keyTypesIS := swag.JoinByFormat(keyTypesIC, "")

	return keyTypesIS
}

// bindParamSearchKeys binds the parameter organizationIds
func (o *SearchKeysParams) bindParamOrganizationIds(formats strfmt.Registry) []string {
	organizationIdsIR := o.OrganizationIds

	var organizationIdsIC []string
	for _, organizationIdsIIR := range organizationIdsIR { // explode []string

		organizationIdsIIV := organizationIdsIIR // string as string
		organizationIdsIC = append(organizationIdsIC, organizationIdsIIV)
	}

	// items.CollectionFormat: ""
	organizationIdsIS := swag.JoinByFormat(organizationIdsIC, "")

	return organizationIdsIS
}
