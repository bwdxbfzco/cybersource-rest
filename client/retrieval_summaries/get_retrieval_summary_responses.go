// Code generated by go-swagger; DO NOT EDIT.

package retrieval_summaries

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

// GetRetrievalSummaryReader is a Reader for the GetRetrievalSummary structure.
type GetRetrievalSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRetrievalSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRetrievalSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /reporting/v3/retrieval-summaries] getRetrievalSummary", response, response.Code())
	}
}

// NewGetRetrievalSummaryOK creates a GetRetrievalSummaryOK with default headers values
func NewGetRetrievalSummaryOK() *GetRetrievalSummaryOK {
	return &GetRetrievalSummaryOK{}
}

/*
GetRetrievalSummaryOK describes a response with status code 200, with default header values.

Ok
*/
type GetRetrievalSummaryOK struct {
	Payload *GetRetrievalSummaryOKBody
}

// IsSuccess returns true when this get retrieval summary o k response has a 2xx status code
func (o *GetRetrievalSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get retrieval summary o k response has a 3xx status code
func (o *GetRetrievalSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get retrieval summary o k response has a 4xx status code
func (o *GetRetrievalSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get retrieval summary o k response has a 5xx status code
func (o *GetRetrievalSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get retrieval summary o k response a status code equal to that given
func (o *GetRetrievalSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get retrieval summary o k response
func (o *GetRetrievalSummaryOK) Code() int {
	return 200
}

func (o *GetRetrievalSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/retrieval-summaries][%d] getRetrievalSummaryOK %s", 200, payload)
}

func (o *GetRetrievalSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/retrieval-summaries][%d] getRetrievalSummaryOK %s", 200, payload)
}

func (o *GetRetrievalSummaryOK) GetPayload() *GetRetrievalSummaryOKBody {
	return o.Payload
}

func (o *GetRetrievalSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRetrievalSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetRetrievalSummaryOKBody reportingV3RetrievalSummariesGet200Response
swagger:model GetRetrievalSummaryOKBody
*/
type GetRetrievalSummaryOKBody struct {

	// Report Start Date
	// Example: 2017-10-01T10:10:10+05:00
	EndTime string `json:"endTime,omitempty" xml:"ReportEndDate,attr,omitempty"`

	// Organization Id
	// Example: testrest
	OrganizationID string `json:"organizationId,omitempty" xml:"MerchantId,attr,omitempty"`

	// List of Summary values
	RetrievalSummaries []*GetRetrievalSummaryOKBodyRetrievalSummariesItems0 `json:"retrievalSummaries"`

	// Report Start Date
	// Example: 2017-10-01T10:10:10+05:00
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty" xml:"ReportStartDate,attr,omitempty"`
}

// Validate validates this get retrieval summary o k body
func (o *GetRetrievalSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRetrievalSummaries(formats); err != nil {
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

func (o *GetRetrievalSummaryOKBody) validateRetrievalSummaries(formats strfmt.Registry) error {
	if swag.IsZero(o.RetrievalSummaries) { // not required
		return nil
	}

	for i := 0; i < len(o.RetrievalSummaries); i++ {
		if swag.IsZero(o.RetrievalSummaries[i]) { // not required
			continue
		}

		if o.RetrievalSummaries[i] != nil {
			if err := o.RetrievalSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRetrievalSummaryOK" + "." + "retrievalSummaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRetrievalSummaryOK" + "." + "retrievalSummaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRetrievalSummaryOKBody) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getRetrievalSummaryOK"+"."+"startTime", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get retrieval summary o k body based on the context it is used
func (o *GetRetrievalSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRetrievalSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRetrievalSummaryOKBody) contextValidateRetrievalSummaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RetrievalSummaries); i++ {

		if o.RetrievalSummaries[i] != nil {

			if swag.IsZero(o.RetrievalSummaries[i]) { // not required
				return nil
			}

			if err := o.RetrievalSummaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRetrievalSummaryOK" + "." + "retrievalSummaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRetrievalSummaryOK" + "." + "retrievalSummaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRetrievalSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRetrievalSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res GetRetrievalSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRetrievalSummaryOKBodyRetrievalSummariesItems0 get retrieval summary o k body retrieval summaries items0
swagger:model GetRetrievalSummaryOKBodyRetrievalSummariesItems0
*/
type GetRetrievalSummaryOKBodyRetrievalSummariesItems0 struct {

	// Account Id
	// Example: testrest_acct
	AccountID string `json:"accountId,omitempty"`

	// Chargeback summary list count
	// Example: 8
	Count float64 `json:"count,omitempty"`

	// Summary Date
	// Example: 2018-01-04T11:33:06.000-0800
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this get retrieval summary o k body retrieval summaries items0
func (o *GetRetrievalSummaryOKBodyRetrievalSummariesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRetrievalSummaryOKBodyRetrievalSummariesItems0) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(o.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", o.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get retrieval summary o k body retrieval summaries items0 based on context it is used
func (o *GetRetrievalSummaryOKBodyRetrievalSummariesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRetrievalSummaryOKBodyRetrievalSummariesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRetrievalSummaryOKBodyRetrievalSummariesItems0) UnmarshalBinary(b []byte) error {
	var res GetRetrievalSummaryOKBodyRetrievalSummariesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
