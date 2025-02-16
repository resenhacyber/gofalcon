// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// RTRListScriptsReader is a Reader for the RTRListScripts structure.
type RTRListScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListScriptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRListScriptsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRListScriptsOK creates a RTRListScriptsOK with default headers values
func NewRTRListScriptsOK() *RTRListScriptsOK {
	return &RTRListScriptsOK{}
}

/*
RTRListScriptsOK describes a response with status code 200, with default header values.

OK
*/
type RTRListScriptsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.BinservclientMsaPutFileResponse
}

// IsSuccess returns true when this r t r list scripts o k response has a 2xx status code
func (o *RTRListScriptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r list scripts o k response has a 3xx status code
func (o *RTRListScriptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list scripts o k response has a 4xx status code
func (o *RTRListScriptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list scripts o k response has a 5xx status code
func (o *RTRListScriptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list scripts o k response a status code equal to that given
func (o *RTRListScriptsOK) IsCode(code int) bool {
	return code == 200
}

func (o *RTRListScriptsOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRListScriptsOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRListScriptsOK) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsBadRequest creates a RTRListScriptsBadRequest with default headers values
func NewRTRListScriptsBadRequest() *RTRListScriptsBadRequest {
	return &RTRListScriptsBadRequest{}
}

/*
RTRListScriptsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListScriptsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list scripts bad request response has a 2xx status code
func (o *RTRListScriptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list scripts bad request response has a 3xx status code
func (o *RTRListScriptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list scripts bad request response has a 4xx status code
func (o *RTRListScriptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list scripts bad request response has a 5xx status code
func (o *RTRListScriptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list scripts bad request response a status code equal to that given
func (o *RTRListScriptsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RTRListScriptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListScriptsBadRequest) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListScriptsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsForbidden creates a RTRListScriptsForbidden with default headers values
func NewRTRListScriptsForbidden() *RTRListScriptsForbidden {
	return &RTRListScriptsForbidden{}
}

/*
RTRListScriptsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListScriptsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r list scripts forbidden response has a 2xx status code
func (o *RTRListScriptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list scripts forbidden response has a 3xx status code
func (o *RTRListScriptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list scripts forbidden response has a 4xx status code
func (o *RTRListScriptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list scripts forbidden response has a 5xx status code
func (o *RTRListScriptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list scripts forbidden response a status code equal to that given
func (o *RTRListScriptsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RTRListScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListScriptsForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListScriptsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsNotFound creates a RTRListScriptsNotFound with default headers values
func NewRTRListScriptsNotFound() *RTRListScriptsNotFound {
	return &RTRListScriptsNotFound{}
}

/*
RTRListScriptsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListScriptsNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list scripts not found response has a 2xx status code
func (o *RTRListScriptsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list scripts not found response has a 3xx status code
func (o *RTRListScriptsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list scripts not found response has a 4xx status code
func (o *RTRListScriptsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list scripts not found response has a 5xx status code
func (o *RTRListScriptsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list scripts not found response a status code equal to that given
func (o *RTRListScriptsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RTRListScriptsNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListScriptsNotFound) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListScriptsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListScriptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsTooManyRequests creates a RTRListScriptsTooManyRequests with default headers values
func NewRTRListScriptsTooManyRequests() *RTRListScriptsTooManyRequests {
	return &RTRListScriptsTooManyRequests{}
}

/*
RTRListScriptsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListScriptsTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r list scripts too many requests response has a 2xx status code
func (o *RTRListScriptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list scripts too many requests response has a 3xx status code
func (o *RTRListScriptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list scripts too many requests response has a 4xx status code
func (o *RTRListScriptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list scripts too many requests response has a 5xx status code
func (o *RTRListScriptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list scripts too many requests response a status code equal to that given
func (o *RTRListScriptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *RTRListScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListScriptsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] rTRListScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListScriptsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListScriptsDefault creates a RTRListScriptsDefault with default headers values
func NewRTRListScriptsDefault(code int) *RTRListScriptsDefault {
	return &RTRListScriptsDefault{
		_statusCode: code,
	}
}

/*
RTRListScriptsDefault describes a response with status code -1, with default header values.

OK
*/
type RTRListScriptsDefault struct {
	_statusCode int

	Payload *models.BinservclientMsaPutFileResponse
}

// Code gets the status code for the r t r list scripts default response
func (o *RTRListScriptsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this r t r list scripts default response has a 2xx status code
func (o *RTRListScriptsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r list scripts default response has a 3xx status code
func (o *RTRListScriptsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r list scripts default response has a 4xx status code
func (o *RTRListScriptsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r list scripts default response has a 5xx status code
func (o *RTRListScriptsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r list scripts default response a status code equal to that given
func (o *RTRListScriptsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RTRListScriptsDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] RTR-ListScripts default  %+v", o._statusCode, o.Payload)
}

func (o *RTRListScriptsDefault) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/scripts/v1][%d] RTR-ListScripts default  %+v", o._statusCode, o.Payload)
}

func (o *RTRListScriptsDefault) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListScriptsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
