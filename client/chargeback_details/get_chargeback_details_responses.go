// Code generated by go-swagger; DO NOT EDIT.

package chargeback_details

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

// GetChargebackDetailsReader is a Reader for the GetChargebackDetails structure.
type GetChargebackDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChargebackDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChargebackDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /reporting/v3/chargeback-details] getChargebackDetails", response, response.Code())
	}
}

// NewGetChargebackDetailsOK creates a GetChargebackDetailsOK with default headers values
func NewGetChargebackDetailsOK() *GetChargebackDetailsOK {
	return &GetChargebackDetailsOK{}
}

/*
GetChargebackDetailsOK describes a response with status code 200, with default header values.

Ok
*/
type GetChargebackDetailsOK struct {
	Payload *GetChargebackDetailsOKBody
}

// IsSuccess returns true when this get chargeback details o k response has a 2xx status code
func (o *GetChargebackDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get chargeback details o k response has a 3xx status code
func (o *GetChargebackDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get chargeback details o k response has a 4xx status code
func (o *GetChargebackDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get chargeback details o k response has a 5xx status code
func (o *GetChargebackDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get chargeback details o k response a status code equal to that given
func (o *GetChargebackDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get chargeback details o k response
func (o *GetChargebackDetailsOK) Code() int {
	return 200
}

func (o *GetChargebackDetailsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/chargeback-details][%d] getChargebackDetailsOK %s", 200, payload)
}

func (o *GetChargebackDetailsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/chargeback-details][%d] getChargebackDetailsOK %s", 200, payload)
}

func (o *GetChargebackDetailsOK) GetPayload() *GetChargebackDetailsOKBody {
	return o.Payload
}

func (o *GetChargebackDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetChargebackDetailsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetChargebackDetailsOKBody reportingV3ChargebackDetailsGet200Response
swagger:model GetChargebackDetailsOKBody
*/
type GetChargebackDetailsOKBody struct {

	// List of Chargeback Details list.
	ChargebackDetails []*GetChargebackDetailsOKBodyChargebackDetailsItems0 `json:"chargebackDetails"`

	// Report Start Date (ISO 8601 Extended)
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty" xml:"ReportEndDate,attr,omitempty"`

	// Organization Id
	// Example: testrest
	OrganizationID string `json:"organizationId,omitempty" xml:"MerchantId,attr,omitempty"`

	// Report Start Date (ISO 8601 Extended)
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty" xml:"ReportStartDate,attr,omitempty"`
}

// Validate validates this get chargeback details o k body
func (o *GetChargebackDetailsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChargebackDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetChargebackDetailsOKBody) validateChargebackDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.ChargebackDetails) { // not required
		return nil
	}

	for i := 0; i < len(o.ChargebackDetails); i++ {
		if swag.IsZero(o.ChargebackDetails[i]) { // not required
			continue
		}

		if o.ChargebackDetails[i] != nil {
			if err := o.ChargebackDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getChargebackDetailsOK" + "." + "chargebackDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getChargebackDetailsOK" + "." + "chargebackDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetChargebackDetailsOKBody) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getChargebackDetailsOK"+"."+"endTime", "body", "date-time", o.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetChargebackDetailsOKBody) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getChargebackDetailsOK"+"."+"startTime", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get chargeback details o k body based on the context it is used
func (o *GetChargebackDetailsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateChargebackDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetChargebackDetailsOKBody) contextValidateChargebackDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ChargebackDetails); i++ {

		if o.ChargebackDetails[i] != nil {

			if swag.IsZero(o.ChargebackDetails[i]) { // not required
				return nil
			}

			if err := o.ChargebackDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getChargebackDetailsOK" + "." + "chargebackDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getChargebackDetailsOK" + "." + "chargebackDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetChargebackDetailsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChargebackDetailsOKBody) UnmarshalBinary(b []byte) error {
	var res GetChargebackDetailsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetChargebackDetailsOKBodyChargebackDetailsItems0 get chargeback details o k body chargeback details items0
swagger:model GetChargebackDetailsOKBodyChargebackDetailsItems0
*/
type GetChargebackDetailsOKBodyChargebackDetailsItems0 struct {

	// Chargeback Action
	// Example: 3
	Action string `json:"action,omitempty"`

	// Chargeback Action Description
	// Example: Financial transaction
	ActionDescription string `json:"actionDescription,omitempty"`

	// Chargeback Alert Type
	// Example: 2
	AlertType string `json:"alertType,omitempty"`

	// Chargeback Amount
	// Example: 5
	Amount string `json:"amount,omitempty"`

	// ICS Request Applications
	// Example: ics_bill
	Applications string `json:"applications,omitempty"`

	// Card Type
	// Example: American Express
	CardType string `json:"cardType,omitempty"`

	// Valid ISO 4217 ALPHA-3 currency code
	// Example: USD
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Customer Id
	// Example: 937999JFK
	CustomerID string `json:"customerId,omitempty"`

	// Event Request Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	EventRequestedTime strfmt.DateTime `json:"eventRequestedTime,omitempty"`

	// Merchant Name
	// Example: Revolutionary Entertainment Inc
	MerchantName string `json:"merchantName,omitempty"`

	// Merchant Reference Number
	// Example: X03434388DEADBEEF
	MerchantReferenceNumber string `json:"merchantReferenceNumber,omitempty"`

	// Nature of Dispute
	// Example: Chargeback
	NatureOfDispute string `json:"natureOfDispute,omitempty"`

	// Original Settlement Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	OriginalSettlementTime strfmt.DateTime `json:"originalSettlementTime,omitempty"`

	// Pre Dispute Flag
	// Example: N
	PreDisputeFlag string `json:"preDisputeFlag,omitempty"`

	// Processor Merchant Id
	// Example: 174263416896
	ProcessorMerchantID string `json:"processorMerchantId,omitempty"`

	// Chargeback Reason Code
	// Example: 1050
	ReasonCode string `json:"reasonCode,omitempty"`

	// Representment CP Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	RepresentmentCPTime strfmt.DateTime `json:"representmentCPTime,omitempty"`

	// Request Id
	// Example: 5060113732046412501541
	RequestID string `json:"requestId,omitempty"`

	// Response Due Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	ResponseDueTime strfmt.DateTime `json:"responseDueTime,omitempty"`

	// Chargeback Sign
	// Example: C
	Sign string `json:"sign,omitempty"`

	// Chargeback Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Tracking Number
	// Example: 990175
	TrackingNumber string `json:"trackingNumber,omitempty"`

	// Transaction Reference Number
	// Example: 93983883073
	TransactionReferenceNumber string `json:"transactionReferenceNumber,omitempty"`
}

// Validate validates this get chargeback details o k body chargeback details items0
func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEventRequestedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginalSettlementTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRepresentmentCPTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResponseDueTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) validateEventRequestedTime(formats strfmt.Registry) error {
	if swag.IsZero(o.EventRequestedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventRequestedTime", "body", "date-time", o.EventRequestedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) validateOriginalSettlementTime(formats strfmt.Registry) error {
	if swag.IsZero(o.OriginalSettlementTime) { // not required
		return nil
	}

	if err := validate.FormatOf("originalSettlementTime", "body", "date-time", o.OriginalSettlementTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) validateRepresentmentCPTime(formats strfmt.Registry) error {
	if swag.IsZero(o.RepresentmentCPTime) { // not required
		return nil
	}

	if err := validate.FormatOf("representmentCPTime", "body", "date-time", o.RepresentmentCPTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) validateResponseDueTime(formats strfmt.Registry) error {
	if swag.IsZero(o.ResponseDueTime) { // not required
		return nil
	}

	if err := validate.FormatOf("responseDueTime", "body", "date-time", o.ResponseDueTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(o.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", o.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get chargeback details o k body chargeback details items0 based on context it is used
func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChargebackDetailsOKBodyChargebackDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetChargebackDetailsOKBodyChargebackDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
