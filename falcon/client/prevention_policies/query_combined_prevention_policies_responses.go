// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// QueryCombinedPreventionPoliciesReader is a Reader for the QueryCombinedPreventionPolicies structure.
type QueryCombinedPreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedPreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedPreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedPreventionPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedPreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedPreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedPreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedPreventionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedPreventionPoliciesOK creates a QueryCombinedPreventionPoliciesOK with default headers values
func NewQueryCombinedPreventionPoliciesOK() *QueryCombinedPreventionPoliciesOK {
	return &QueryCombinedPreventionPoliciesOK{}
}

/*
QueryCombinedPreventionPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedPreventionPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this query combined prevention policies o k response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined prevention policies o k response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined prevention policies o k response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined prevention policies o k response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined prevention policies o k response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryCombinedPreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesOK) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPoliciesBadRequest creates a QueryCombinedPreventionPoliciesBadRequest with default headers values
func NewQueryCombinedPreventionPoliciesBadRequest() *QueryCombinedPreventionPoliciesBadRequest {
	return &QueryCombinedPreventionPoliciesBadRequest{}
}

/*
QueryCombinedPreventionPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedPreventionPoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this query combined prevention policies bad request response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined prevention policies bad request response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined prevention policies bad request response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined prevention policies bad request response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined prevention policies bad request response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryCombinedPreventionPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesBadRequest) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPoliciesForbidden creates a QueryCombinedPreventionPoliciesForbidden with default headers values
func NewQueryCombinedPreventionPoliciesForbidden() *QueryCombinedPreventionPoliciesForbidden {
	return &QueryCombinedPreventionPoliciesForbidden{}
}

/*
QueryCombinedPreventionPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedPreventionPoliciesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query combined prevention policies forbidden response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined prevention policies forbidden response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined prevention policies forbidden response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined prevention policies forbidden response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined prevention policies forbidden response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryCombinedPreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPoliciesTooManyRequests creates a QueryCombinedPreventionPoliciesTooManyRequests with default headers values
func NewQueryCombinedPreventionPoliciesTooManyRequests() *QueryCombinedPreventionPoliciesTooManyRequests {
	return &QueryCombinedPreventionPoliciesTooManyRequests{}
}

/*
QueryCombinedPreventionPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedPreventionPoliciesTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this query combined prevention policies too many requests response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined prevention policies too many requests response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined prevention policies too many requests response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined prevention policies too many requests response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined prevention policies too many requests response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryCombinedPreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewQueryCombinedPreventionPoliciesInternalServerError creates a QueryCombinedPreventionPoliciesInternalServerError with default headers values
func NewQueryCombinedPreventionPoliciesInternalServerError() *QueryCombinedPreventionPoliciesInternalServerError {
	return &QueryCombinedPreventionPoliciesInternalServerError{}
}

/*
QueryCombinedPreventionPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedPreventionPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPreventionPoliciesV1
}

// IsSuccess returns true when this query combined prevention policies internal server error response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined prevention policies internal server error response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined prevention policies internal server error response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined prevention policies internal server error response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined prevention policies internal server error response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryCombinedPreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesInternalServerError) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedPreventionPoliciesDefault creates a QueryCombinedPreventionPoliciesDefault with default headers values
func NewQueryCombinedPreventionPoliciesDefault(code int) *QueryCombinedPreventionPoliciesDefault {
	return &QueryCombinedPreventionPoliciesDefault{
		_statusCode: code,
	}
}

/*
QueryCombinedPreventionPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedPreventionPoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesPreventionPoliciesV1
}

// Code gets the status code for the query combined prevention policies default response
func (o *QueryCombinedPreventionPoliciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query combined prevention policies default response has a 2xx status code
func (o *QueryCombinedPreventionPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query combined prevention policies default response has a 3xx status code
func (o *QueryCombinedPreventionPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query combined prevention policies default response has a 4xx status code
func (o *QueryCombinedPreventionPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query combined prevention policies default response has a 5xx status code
func (o *QueryCombinedPreventionPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query combined prevention policies default response a status code equal to that given
func (o *QueryCombinedPreventionPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryCombinedPreventionPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesDefault) String() string {
	return fmt.Sprintf("[GET /policy/combined/prevention/v1][%d] queryCombinedPreventionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedPreventionPoliciesDefault) GetPayload() *models.ResponsesPreventionPoliciesV1 {
	return o.Payload
}

func (o *QueryCombinedPreventionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPreventionPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
