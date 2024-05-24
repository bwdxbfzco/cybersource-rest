// Code generated by go-swagger; DO NOT EDIT.

package key_management_scmp

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
)

// UpdateSCMPReader is a Reader for the UpdateSCMP structure.
type UpdateSCMPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSCMPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSCMPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSCMPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSCMPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /kms/v2/keys-scmp/{keyId}] UpdateSCMP", response, response.Code())
	}
}

// NewUpdateSCMPOK creates a UpdateSCMPOK with default headers values
func NewUpdateSCMPOK() *UpdateSCMPOK {
	return &UpdateSCMPOK{}
}

/*
UpdateSCMPOK describes a response with status code 200, with default header values.

Successful response.
*/
type UpdateSCMPOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update s c m p o k response has a 2xx status code
func (o *UpdateSCMPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update s c m p o k response has a 3xx status code
func (o *UpdateSCMPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update s c m p o k response has a 4xx status code
func (o *UpdateSCMPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update s c m p o k response has a 5xx status code
func (o *UpdateSCMPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update s c m p o k response a status code equal to that given
func (o *UpdateSCMPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update s c m p o k response
func (o *UpdateSCMPOK) Code() int {
	return 200
}

func (o *UpdateSCMPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPOK %s", 200, payload)
}

func (o *UpdateSCMPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPOK %s", 200, payload)
}

func (o *UpdateSCMPOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateSCMPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSCMPBadRequest creates a UpdateSCMPBadRequest with default headers values
func NewUpdateSCMPBadRequest() *UpdateSCMPBadRequest {
	return &UpdateSCMPBadRequest{}
}

/*
UpdateSCMPBadRequest describes a response with status code 400, with default header values.

Invalid request.
*/
type UpdateSCMPBadRequest struct {
	Payload *UpdateSCMPBadRequestBody
}

// IsSuccess returns true when this update s c m p bad request response has a 2xx status code
func (o *UpdateSCMPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update s c m p bad request response has a 3xx status code
func (o *UpdateSCMPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update s c m p bad request response has a 4xx status code
func (o *UpdateSCMPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update s c m p bad request response has a 5xx status code
func (o *UpdateSCMPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update s c m p bad request response a status code equal to that given
func (o *UpdateSCMPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update s c m p bad request response
func (o *UpdateSCMPBadRequest) Code() int {
	return 400
}

func (o *UpdateSCMPBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPBadRequest %s", 400, payload)
}

func (o *UpdateSCMPBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPBadRequest %s", 400, payload)
}

func (o *UpdateSCMPBadRequest) GetPayload() *UpdateSCMPBadRequestBody {
	return o.Payload
}

func (o *UpdateSCMPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSCMPBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSCMPInternalServerError creates a UpdateSCMPInternalServerError with default headers values
func NewUpdateSCMPInternalServerError() *UpdateSCMPInternalServerError {
	return &UpdateSCMPInternalServerError{}
}

/*
UpdateSCMPInternalServerError describes a response with status code 500, with default header values.

Unexpected system error or system timeout.
*/
type UpdateSCMPInternalServerError struct {
	Payload *UpdateSCMPInternalServerErrorBody
}

// IsSuccess returns true when this update s c m p internal server error response has a 2xx status code
func (o *UpdateSCMPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update s c m p internal server error response has a 3xx status code
func (o *UpdateSCMPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update s c m p internal server error response has a 4xx status code
func (o *UpdateSCMPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update s c m p internal server error response has a 5xx status code
func (o *UpdateSCMPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update s c m p internal server error response a status code equal to that given
func (o *UpdateSCMPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update s c m p internal server error response
func (o *UpdateSCMPInternalServerError) Code() int {
	return 500
}

func (o *UpdateSCMPInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPInternalServerError %s", 500, payload)
}

func (o *UpdateSCMPInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /kms/v2/keys-scmp/{keyId}][%d] updateSCMPInternalServerError %s", 500, payload)
}

func (o *UpdateSCMPInternalServerError) GetPayload() *UpdateSCMPInternalServerErrorBody {
	return o.Payload
}

func (o *UpdateSCMPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateSCMPInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateSCMPBadRequestBody update s c m p bad request body
swagger:model UpdateSCMPBadRequestBody
*/
type UpdateSCMPBadRequestBody struct {

	// details
	Details []*UpdateSCMPBadRequestBodyDetailsItems0 `json:"details"`

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
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

// Validate validates this update s c m p bad request body
func (o *UpdateSCMPBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSCMPBadRequestBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateSCMPBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSCMPBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update s c m p bad request body based on the context it is used
func (o *UpdateSCMPBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSCMPBadRequestBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {

			if swag.IsZero(o.Details[i]) { // not required
				return nil
			}

			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSCMPBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSCMPBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSCMPBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSCMPBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateSCMPBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateSCMPBadRequestBodyDetailsItems0 update s c m p bad request body details items0
swagger:model UpdateSCMPBadRequestBodyDetailsItems0
*/
type UpdateSCMPBadRequestBodyDetailsItems0 struct {

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

// Validate validates this update s c m p bad request body details items0
func (o *UpdateSCMPBadRequestBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update s c m p bad request body details items0 based on context it is used
func (o *UpdateSCMPBadRequestBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSCMPBadRequestBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSCMPBadRequestBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateSCMPBadRequestBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateSCMPBody update s c m p body
swagger:model UpdateSCMPBody
*/
type UpdateSCMPBody struct {

	// Comment. Optional field. Can be provided along with Expiration Date and Version
	Comment string `json:"comment,omitempty"`

	// Expiration Date. Required field to update the SCMP_API key
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// Only inactive status is applicable for SCMP_API. Only status as inactive needs to be provided to deactivate scmp.
	Status string `json:"status,omitempty"`

	// Version. Required field to update the SCMP_API key
	Version string `json:"version,omitempty"`
}

// Validate validates this update s c m p body
func (o *UpdateSCMPBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update s c m p body based on context it is used
func (o *UpdateSCMPBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSCMPBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSCMPBody) UnmarshalBinary(b []byte) error {
	var res UpdateSCMPBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateSCMPInternalServerErrorBody update s c m p internal server error body
swagger:model UpdateSCMPInternalServerErrorBody
*/
type UpdateSCMPInternalServerErrorBody struct {

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

	// The status of the submitted request.
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

// Validate validates this update s c m p internal server error body
func (o *UpdateSCMPInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update s c m p internal server error body based on context it is used
func (o *UpdateSCMPInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSCMPInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSCMPInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res UpdateSCMPInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
