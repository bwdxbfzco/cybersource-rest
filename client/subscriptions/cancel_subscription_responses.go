// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CancelSubscriptionReader is a Reader for the CancelSubscription structure.
type CancelSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCancelSubscriptionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCancelSubscriptionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /rbs/v1/subscriptions/{id}/cancel] cancelSubscription", response, response.Code())
	}
}

// NewCancelSubscriptionAccepted creates a CancelSubscriptionAccepted with default headers values
func NewCancelSubscriptionAccepted() *CancelSubscriptionAccepted {
	return &CancelSubscriptionAccepted{}
}

/*
CancelSubscriptionAccepted describes a response with status code 202, with default header values.

Successful response.
*/
type CancelSubscriptionAccepted struct {
	Payload *CancelSubscriptionAcceptedBody
}

// IsSuccess returns true when this cancel subscription accepted response has a 2xx status code
func (o *CancelSubscriptionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel subscription accepted response has a 3xx status code
func (o *CancelSubscriptionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel subscription accepted response has a 4xx status code
func (o *CancelSubscriptionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel subscription accepted response has a 5xx status code
func (o *CancelSubscriptionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel subscription accepted response a status code equal to that given
func (o *CancelSubscriptionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the cancel subscription accepted response
func (o *CancelSubscriptionAccepted) Code() int {
	return 202
}

func (o *CancelSubscriptionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionAccepted %s", 202, payload)
}

func (o *CancelSubscriptionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionAccepted %s", 202, payload)
}

func (o *CancelSubscriptionAccepted) GetPayload() *CancelSubscriptionAcceptedBody {
	return o.Payload
}

func (o *CancelSubscriptionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CancelSubscriptionAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSubscriptionBadRequest creates a CancelSubscriptionBadRequest with default headers values
func NewCancelSubscriptionBadRequest() *CancelSubscriptionBadRequest {
	return &CancelSubscriptionBadRequest{}
}

/*
CancelSubscriptionBadRequest describes a response with status code 400, with default header values.

Invalid request.
*/
type CancelSubscriptionBadRequest struct {
	Payload *CancelSubscriptionBadRequestBody
}

// IsSuccess returns true when this cancel subscription bad request response has a 2xx status code
func (o *CancelSubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel subscription bad request response has a 3xx status code
func (o *CancelSubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel subscription bad request response has a 4xx status code
func (o *CancelSubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel subscription bad request response has a 5xx status code
func (o *CancelSubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel subscription bad request response a status code equal to that given
func (o *CancelSubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancel subscription bad request response
func (o *CancelSubscriptionBadRequest) Code() int {
	return 400
}

func (o *CancelSubscriptionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionBadRequest %s", 400, payload)
}

func (o *CancelSubscriptionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionBadRequest %s", 400, payload)
}

func (o *CancelSubscriptionBadRequest) GetPayload() *CancelSubscriptionBadRequestBody {
	return o.Payload
}

func (o *CancelSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CancelSubscriptionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSubscriptionNotFound creates a CancelSubscriptionNotFound with default headers values
func NewCancelSubscriptionNotFound() *CancelSubscriptionNotFound {
	return &CancelSubscriptionNotFound{}
}

/*
CancelSubscriptionNotFound describes a response with status code 404, with default header values.

Not found.
*/
type CancelSubscriptionNotFound struct {
	Payload *CancelSubscriptionNotFoundBody
}

// IsSuccess returns true when this cancel subscription not found response has a 2xx status code
func (o *CancelSubscriptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel subscription not found response has a 3xx status code
func (o *CancelSubscriptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel subscription not found response has a 4xx status code
func (o *CancelSubscriptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel subscription not found response has a 5xx status code
func (o *CancelSubscriptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel subscription not found response a status code equal to that given
func (o *CancelSubscriptionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cancel subscription not found response
func (o *CancelSubscriptionNotFound) Code() int {
	return 404
}

func (o *CancelSubscriptionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionNotFound %s", 404, payload)
}

func (o *CancelSubscriptionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionNotFound %s", 404, payload)
}

func (o *CancelSubscriptionNotFound) GetPayload() *CancelSubscriptionNotFoundBody {
	return o.Payload
}

func (o *CancelSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CancelSubscriptionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSubscriptionBadGateway creates a CancelSubscriptionBadGateway with default headers values
func NewCancelSubscriptionBadGateway() *CancelSubscriptionBadGateway {
	return &CancelSubscriptionBadGateway{}
}

/*
CancelSubscriptionBadGateway describes a response with status code 502, with default header values.

Unexpected system error or system timeout.
*/
type CancelSubscriptionBadGateway struct {
	Payload *CancelSubscriptionBadGatewayBody
}

// IsSuccess returns true when this cancel subscription bad gateway response has a 2xx status code
func (o *CancelSubscriptionBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel subscription bad gateway response has a 3xx status code
func (o *CancelSubscriptionBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel subscription bad gateway response has a 4xx status code
func (o *CancelSubscriptionBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel subscription bad gateway response has a 5xx status code
func (o *CancelSubscriptionBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel subscription bad gateway response a status code equal to that given
func (o *CancelSubscriptionBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the cancel subscription bad gateway response
func (o *CancelSubscriptionBadGateway) Code() int {
	return 502
}

func (o *CancelSubscriptionBadGateway) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionBadGateway %s", 502, payload)
}

func (o *CancelSubscriptionBadGateway) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /rbs/v1/subscriptions/{id}/cancel][%d] cancelSubscriptionBadGateway %s", 502, payload)
}

func (o *CancelSubscriptionBadGateway) GetPayload() *CancelSubscriptionBadGatewayBody {
	return o.Payload
}

func (o *CancelSubscriptionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CancelSubscriptionBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CancelSubscriptionAcceptedBody cancelSubscriptionResponse
swagger:model CancelSubscriptionAcceptedBody
*/
type CancelSubscriptionAcceptedBody struct {

	// links
	Links *CancelSubscriptionAcceptedBodyLinks `json:"_links,omitempty"`

	// An unique identification number generated by Cybersource to identify the submitted request. Returned by all services.
	// It is also appended to the endpoint of the resource.
	// On incremental authorizations, this value with be the same as the identification number returned in the original authorization response.
	//
	// Max Length: 26
	ID string `json:"id,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - ACCEPTED
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by Cybersource for all services.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`

	// subscription information
	SubscriptionInformation *CancelSubscriptionAcceptedBodySubscriptionInformation `json:"subscriptionInformation,omitempty"`
}

// Validate validates this cancel subscription accepted body
func (o *CancelSubscriptionAcceptedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubscriptionInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionAcceptedBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *CancelSubscriptionAcceptedBody) validateID(formats strfmt.Registry) error {
	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("cancelSubscriptionAccepted"+"."+"id", "body", o.ID, 26); err != nil {
		return err
	}

	return nil
}

func (o *CancelSubscriptionAcceptedBody) validateSubscriptionInformation(formats strfmt.Registry) error {
	if swag.IsZero(o.SubscriptionInformation) { // not required
		return nil
	}

	if o.SubscriptionInformation != nil {
		if err := o.SubscriptionInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "subscriptionInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "subscriptionInformation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cancel subscription accepted body based on the context it is used
func (o *CancelSubscriptionAcceptedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSubscriptionInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionAcceptedBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {

		if swag.IsZero(o.Links) { // not required
			return nil
		}

		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *CancelSubscriptionAcceptedBody) contextValidateSubscriptionInformation(ctx context.Context, formats strfmt.Registry) error {

	if o.SubscriptionInformation != nil {

		if swag.IsZero(o.SubscriptionInformation) { // not required
			return nil
		}

		if err := o.SubscriptionInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "subscriptionInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "subscriptionInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBody) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionAcceptedBodyLinks cancel subscription accepted body links
swagger:model CancelSubscriptionAcceptedBodyLinks
*/
type CancelSubscriptionAcceptedBodyLinks struct {

	// self
	Self *CancelSubscriptionAcceptedBodyLinksSelf `json:"self,omitempty"`
}

// Validate validates this cancel subscription accepted body links
func (o *CancelSubscriptionAcceptedBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionAcceptedBodyLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cancel subscription accepted body links based on the context it is used
func (o *CancelSubscriptionAcceptedBodyLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionAcceptedBodyLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {

		if swag.IsZero(o.Self) { // not required
			return nil
		}

		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelSubscriptionAccepted" + "." + "_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelSubscriptionAccepted" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodyLinks) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionAcceptedBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionAcceptedBodyLinksSelf cancel subscription accepted body links self
swagger:model CancelSubscriptionAcceptedBodyLinksSelf
*/
type CancelSubscriptionAcceptedBodyLinksSelf struct {

	// This is the endpoint of the resource that was created by the successful request.
	Href string `json:"href,omitempty"`

	// `method` refers to the HTTP method that you can send to the `self` endpoint to retrieve details of the resource.
	Method string `json:"method,omitempty"`
}

// Validate validates this cancel subscription accepted body links self
func (o *CancelSubscriptionAcceptedBodyLinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel subscription accepted body links self based on context it is used
func (o *CancelSubscriptionAcceptedBodyLinksSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodyLinksSelf) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodyLinksSelf) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionAcceptedBodyLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionAcceptedBodySubscriptionInformation cancel subscription accepted body subscription information
swagger:model CancelSubscriptionAcceptedBodySubscriptionInformation
*/
type CancelSubscriptionAcceptedBodySubscriptionInformation struct {

	// Subscription code.
	//
	// Max Length: 10
	Code string `json:"code,omitempty"`

	// Subscription Status:
	// - `CANCELLED`
	//
	Status string `json:"status,omitempty"`
}

// Validate validates this cancel subscription accepted body subscription information
func (o *CancelSubscriptionAcceptedBodySubscriptionInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionAcceptedBodySubscriptionInformation) validateCode(formats strfmt.Registry) error {
	if swag.IsZero(o.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("cancelSubscriptionAccepted"+"."+"subscriptionInformation"+"."+"code", "body", o.Code, 10); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cancel subscription accepted body subscription information based on context it is used
func (o *CancelSubscriptionAcceptedBodySubscriptionInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodySubscriptionInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionAcceptedBodySubscriptionInformation) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionAcceptedBodySubscriptionInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionBadGatewayBody cancel subscription bad gateway body
swagger:model CancelSubscriptionBadGatewayBody
*/
type CancelSubscriptionBadGatewayBody struct {

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - SYSTEM_ERROR
	//  - SERVER_TIMEOUT
	//  - SERVICE_TIMEOUT
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - SERVER_ERROR
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by Cybersource for all services.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this cancel subscription bad gateway body
func (o *CancelSubscriptionBadGatewayBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel subscription bad gateway body based on context it is used
func (o *CancelSubscriptionBadGatewayBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionBadRequestBody cancel subscription bad request body
swagger:model CancelSubscriptionBadRequestBody
*/
type CancelSubscriptionBadRequestBody struct {

	// details
	Details []*CancelSubscriptionBadRequestBodyDetailsItems0 `json:"details"`

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//  - DUPLICATE_REQUEST
	//  - INVALID_MERCHANT_CONFIGURATION
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - INVALID_REQUEST
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by Cybersource for all services.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this cancel subscription bad request body
func (o *CancelSubscriptionBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionBadRequestBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cancelSubscriptionBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cancelSubscriptionBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cancel subscription bad request body based on the context it is used
func (o *CancelSubscriptionBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelSubscriptionBadRequestBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cancelSubscriptionBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cancelSubscriptionBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionBadRequestBodyDetailsItems0 cancel subscription bad request body details items0
swagger:model CancelSubscriptionBadRequestBodyDetailsItems0
*/
type CancelSubscriptionBadRequestBodyDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	Field string `json:"field,omitempty"`

	// Possible reasons for the error.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this cancel subscription bad request body details items0
func (o *CancelSubscriptionBadRequestBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel subscription bad request body details items0 based on context it is used
func (o *CancelSubscriptionBadRequestBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionBadRequestBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionBadRequestBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionBadRequestBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CancelSubscriptionNotFoundBody cancel subscription not found body
swagger:model CancelSubscriptionNotFoundBody
*/
type CancelSubscriptionNotFoundBody struct {

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - NOT_FOUND
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by Cybersource for all services.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this cancel subscription not found body
func (o *CancelSubscriptionNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel subscription not found body based on context it is used
func (o *CancelSubscriptionNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelSubscriptionNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelSubscriptionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CancelSubscriptionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
