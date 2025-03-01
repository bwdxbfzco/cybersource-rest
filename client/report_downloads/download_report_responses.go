// Code generated by go-swagger; DO NOT EDIT.

package report_downloads

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

// DownloadReportReader is a Reader for the DownloadReport structure.
type DownloadReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reporting/v3/report-downloads] downloadReport", response, response.Code())
	}
}

// NewDownloadReportOK creates a DownloadReportOK with default headers values
func NewDownloadReportOK() *DownloadReportOK {
	return &DownloadReportOK{}
}

/*
DownloadReportOK describes a response with status code 200, with default header values.

OK
*/
type DownloadReportOK struct {
}

// IsSuccess returns true when this download report o k response has a 2xx status code
func (o *DownloadReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download report o k response has a 3xx status code
func (o *DownloadReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download report o k response has a 4xx status code
func (o *DownloadReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download report o k response has a 5xx status code
func (o *DownloadReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download report o k response a status code equal to that given
func (o *DownloadReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download report o k response
func (o *DownloadReportOK) Code() int {
	return 200
}

func (o *DownloadReportOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportOK", 200)
}

func (o *DownloadReportOK) String() string {
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportOK", 200)
}

func (o *DownloadReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadReportBadRequest creates a DownloadReportBadRequest with default headers values
func NewDownloadReportBadRequest() *DownloadReportBadRequest {
	return &DownloadReportBadRequest{}
}

/*
DownloadReportBadRequest describes a response with status code 400, with default header values.

Invalid Request
*/
type DownloadReportBadRequest struct {
	Payload *DownloadReportBadRequestBody
}

// IsSuccess returns true when this download report bad request response has a 2xx status code
func (o *DownloadReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download report bad request response has a 3xx status code
func (o *DownloadReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download report bad request response has a 4xx status code
func (o *DownloadReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this download report bad request response has a 5xx status code
func (o *DownloadReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this download report bad request response a status code equal to that given
func (o *DownloadReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the download report bad request response
func (o *DownloadReportBadRequest) Code() int {
	return 400
}

func (o *DownloadReportBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportBadRequest %s", 400, payload)
}

func (o *DownloadReportBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportBadRequest %s", 400, payload)
}

func (o *DownloadReportBadRequest) GetPayload() *DownloadReportBadRequestBody {
	return o.Payload
}

func (o *DownloadReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DownloadReportBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadReportNotFound creates a DownloadReportNotFound with default headers values
func NewDownloadReportNotFound() *DownloadReportNotFound {
	return &DownloadReportNotFound{}
}

/*
DownloadReportNotFound describes a response with status code 404, with default header values.

No Reports Found
*/
type DownloadReportNotFound struct {
}

// IsSuccess returns true when this download report not found response has a 2xx status code
func (o *DownloadReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download report not found response has a 3xx status code
func (o *DownloadReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download report not found response has a 4xx status code
func (o *DownloadReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this download report not found response has a 5xx status code
func (o *DownloadReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this download report not found response a status code equal to that given
func (o *DownloadReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the download report not found response
func (o *DownloadReportNotFound) Code() int {
	return 404
}

func (o *DownloadReportNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportNotFound", 404)
}

func (o *DownloadReportNotFound) String() string {
	return fmt.Sprintf("[GET /reporting/v3/report-downloads][%d] downloadReportNotFound", 404)
}

func (o *DownloadReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
DownloadReportBadRequestBody reportingv3ReportDownloadsGet400Response
//
// HTTP status code for client application
swagger:model DownloadReportBadRequestBody
*/
type DownloadReportBadRequestBody struct {

	// Error field list
	//
	// Required: true
	Details []*DownloadReportBadRequestBodyDetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Example: One or more fields contains invalid data
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Example: INVALID_DATA
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Example: 2016-08-11T22:47:57Z
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this download report bad request body
func (o *DownloadReportBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DownloadReportBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("downloadReportBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downloadReportBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("downloadReportBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *DownloadReportBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("downloadReportBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *DownloadReportBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("downloadReportBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *DownloadReportBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("downloadReportBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("downloadReportBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this download report bad request body based on the context it is used
func (o *DownloadReportBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DownloadReportBadRequestBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downloadReportBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("downloadReportBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DownloadReportBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadReportBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DownloadReportBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DownloadReportBadRequestBodyDetailsItems0 Provides failed validation input field detail
//
swagger:model DownloadReportBadRequestBodyDetailsItems0
*/
type DownloadReportBadRequestBodyDetailsItems0 struct {

	// Field in request that caused an error
	//
	Field string `json:"field,omitempty"`

	// Documented reason code
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this download report bad request body details items0
func (o *DownloadReportBadRequestBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download report bad request body details items0 based on context it is used
func (o *DownloadReportBadRequestBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadReportBadRequestBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadReportBadRequestBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DownloadReportBadRequestBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
