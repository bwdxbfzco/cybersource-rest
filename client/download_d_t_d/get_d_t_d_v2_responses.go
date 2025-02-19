// Code generated by go-swagger; DO NOT EDIT.

package download_d_t_d

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDTDV2Reader is a Reader for the GetDTDV2 structure.
type GetDTDV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDTDV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDTDV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDTDV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDTDV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDTDV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}] getDTDV2", response, response.Code())
	}
}

// NewGetDTDV2OK creates a GetDTDV2OK with default headers values
func NewGetDTDV2OK() *GetDTDV2OK {
	return &GetDTDV2OK{}
}

/*
GetDTDV2OK describes a response with status code 200, with default header values.

Ok
*/
type GetDTDV2OK struct {
}

// IsSuccess returns true when this get d t d v2 o k response has a 2xx status code
func (o *GetDTDV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d t d v2 o k response has a 3xx status code
func (o *GetDTDV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d t d v2 o k response has a 4xx status code
func (o *GetDTDV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d t d v2 o k response has a 5xx status code
func (o *GetDTDV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d t d v2 o k response a status code equal to that given
func (o *GetDTDV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get d t d v2 o k response
func (o *GetDTDV2OK) Code() int {
	return 200
}

func (o *GetDTDV2OK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2OK", 200)
}

func (o *GetDTDV2OK) String() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2OK", 200)
}

func (o *GetDTDV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2BadRequest creates a GetDTDV2BadRequest with default headers values
func NewGetDTDV2BadRequest() *GetDTDV2BadRequest {
	return &GetDTDV2BadRequest{}
}

/*
GetDTDV2BadRequest describes a response with status code 400, with default header values.

Bad request. DTD file name may be invalid
*/
type GetDTDV2BadRequest struct {
}

// IsSuccess returns true when this get d t d v2 bad request response has a 2xx status code
func (o *GetDTDV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d t d v2 bad request response has a 3xx status code
func (o *GetDTDV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d t d v2 bad request response has a 4xx status code
func (o *GetDTDV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d t d v2 bad request response has a 5xx status code
func (o *GetDTDV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get d t d v2 bad request response a status code equal to that given
func (o *GetDTDV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get d t d v2 bad request response
func (o *GetDTDV2BadRequest) Code() int {
	return 400
}

func (o *GetDTDV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2BadRequest", 400)
}

func (o *GetDTDV2BadRequest) String() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2BadRequest", 400)
}

func (o *GetDTDV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2NotFound creates a GetDTDV2NotFound with default headers values
func NewGetDTDV2NotFound() *GetDTDV2NotFound {
	return &GetDTDV2NotFound{}
}

/*
GetDTDV2NotFound describes a response with status code 404, with default header values.

DTD file not found
*/
type GetDTDV2NotFound struct {
}

// IsSuccess returns true when this get d t d v2 not found response has a 2xx status code
func (o *GetDTDV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d t d v2 not found response has a 3xx status code
func (o *GetDTDV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d t d v2 not found response has a 4xx status code
func (o *GetDTDV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d t d v2 not found response has a 5xx status code
func (o *GetDTDV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get d t d v2 not found response a status code equal to that given
func (o *GetDTDV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get d t d v2 not found response
func (o *GetDTDV2NotFound) Code() int {
	return 404
}

func (o *GetDTDV2NotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2NotFound", 404)
}

func (o *GetDTDV2NotFound) String() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2NotFound", 404)
}

func (o *GetDTDV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2InternalServerError creates a GetDTDV2InternalServerError with default headers values
func NewGetDTDV2InternalServerError() *GetDTDV2InternalServerError {
	return &GetDTDV2InternalServerError{}
}

/*
GetDTDV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDTDV2InternalServerError struct {
}

// IsSuccess returns true when this get d t d v2 internal server error response has a 2xx status code
func (o *GetDTDV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d t d v2 internal server error response has a 3xx status code
func (o *GetDTDV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d t d v2 internal server error response has a 4xx status code
func (o *GetDTDV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d t d v2 internal server error response has a 5xx status code
func (o *GetDTDV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get d t d v2 internal server error response a status code equal to that given
func (o *GetDTDV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get d t d v2 internal server error response
func (o *GetDTDV2InternalServerError) Code() int {
	return 500
}

func (o *GetDTDV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2InternalServerError", 500)
}

func (o *GetDTDV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /reporting/v3/dtds/{reportDefinitionNameVersion}][%d] getDTDV2InternalServerError", 500)
}

func (o *GetDTDV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
