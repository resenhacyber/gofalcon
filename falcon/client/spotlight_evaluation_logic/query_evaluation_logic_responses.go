// Code generated by go-swagger; DO NOT EDIT.

package spotlight_evaluation_logic

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

// QueryEvaluationLogicReader is a Reader for the QueryEvaluationLogic structure.
type QueryEvaluationLogicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryEvaluationLogicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryEvaluationLogicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryEvaluationLogicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryEvaluationLogicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryEvaluationLogicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryEvaluationLogicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryEvaluationLogicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryEvaluationLogicOK creates a QueryEvaluationLogicOK with default headers values
func NewQueryEvaluationLogicOK() *QueryEvaluationLogicOK {
	return &QueryEvaluationLogicOK{}
}

/*
QueryEvaluationLogicOK describes a response with status code 200, with default header values.

OK
*/
type QueryEvaluationLogicOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIQueryResponse
}

// IsSuccess returns true when this query evaluation logic o k response has a 2xx status code
func (o *QueryEvaluationLogicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query evaluation logic o k response has a 3xx status code
func (o *QueryEvaluationLogicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query evaluation logic o k response has a 4xx status code
func (o *QueryEvaluationLogicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query evaluation logic o k response has a 5xx status code
func (o *QueryEvaluationLogicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query evaluation logic o k response a status code equal to that given
func (o *QueryEvaluationLogicOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryEvaluationLogicOK) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicOK  %+v", 200, o.Payload)
}

func (o *QueryEvaluationLogicOK) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicOK  %+v", 200, o.Payload)
}

func (o *QueryEvaluationLogicOK) GetPayload() *models.DomainSPAPIQueryResponse {
	return o.Payload
}

func (o *QueryEvaluationLogicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEvaluationLogicBadRequest creates a QueryEvaluationLogicBadRequest with default headers values
func NewQueryEvaluationLogicBadRequest() *QueryEvaluationLogicBadRequest {
	return &QueryEvaluationLogicBadRequest{}
}

/*
QueryEvaluationLogicBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryEvaluationLogicBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this query evaluation logic bad request response has a 2xx status code
func (o *QueryEvaluationLogicBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query evaluation logic bad request response has a 3xx status code
func (o *QueryEvaluationLogicBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query evaluation logic bad request response has a 4xx status code
func (o *QueryEvaluationLogicBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query evaluation logic bad request response has a 5xx status code
func (o *QueryEvaluationLogicBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query evaluation logic bad request response a status code equal to that given
func (o *QueryEvaluationLogicBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryEvaluationLogicBadRequest) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicBadRequest ", 400)
}

func (o *QueryEvaluationLogicBadRequest) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicBadRequest ", 400)
}

func (o *QueryEvaluationLogicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewQueryEvaluationLogicForbidden creates a QueryEvaluationLogicForbidden with default headers values
func NewQueryEvaluationLogicForbidden() *QueryEvaluationLogicForbidden {
	return &QueryEvaluationLogicForbidden{}
}

/*
QueryEvaluationLogicForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryEvaluationLogicForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query evaluation logic forbidden response has a 2xx status code
func (o *QueryEvaluationLogicForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query evaluation logic forbidden response has a 3xx status code
func (o *QueryEvaluationLogicForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query evaluation logic forbidden response has a 4xx status code
func (o *QueryEvaluationLogicForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query evaluation logic forbidden response has a 5xx status code
func (o *QueryEvaluationLogicForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query evaluation logic forbidden response a status code equal to that given
func (o *QueryEvaluationLogicForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryEvaluationLogicForbidden) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicForbidden  %+v", 403, o.Payload)
}

func (o *QueryEvaluationLogicForbidden) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicForbidden  %+v", 403, o.Payload)
}

func (o *QueryEvaluationLogicForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryEvaluationLogicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEvaluationLogicTooManyRequests creates a QueryEvaluationLogicTooManyRequests with default headers values
func NewQueryEvaluationLogicTooManyRequests() *QueryEvaluationLogicTooManyRequests {
	return &QueryEvaluationLogicTooManyRequests{}
}

/*
QueryEvaluationLogicTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryEvaluationLogicTooManyRequests struct {

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

// IsSuccess returns true when this query evaluation logic too many requests response has a 2xx status code
func (o *QueryEvaluationLogicTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query evaluation logic too many requests response has a 3xx status code
func (o *QueryEvaluationLogicTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query evaluation logic too many requests response has a 4xx status code
func (o *QueryEvaluationLogicTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query evaluation logic too many requests response has a 5xx status code
func (o *QueryEvaluationLogicTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query evaluation logic too many requests response a status code equal to that given
func (o *QueryEvaluationLogicTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryEvaluationLogicTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryEvaluationLogicTooManyRequests) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryEvaluationLogicTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryEvaluationLogicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryEvaluationLogicInternalServerError creates a QueryEvaluationLogicInternalServerError with default headers values
func NewQueryEvaluationLogicInternalServerError() *QueryEvaluationLogicInternalServerError {
	return &QueryEvaluationLogicInternalServerError{}
}

/*
QueryEvaluationLogicInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryEvaluationLogicInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this query evaluation logic internal server error response has a 2xx status code
func (o *QueryEvaluationLogicInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query evaluation logic internal server error response has a 3xx status code
func (o *QueryEvaluationLogicInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query evaluation logic internal server error response has a 4xx status code
func (o *QueryEvaluationLogicInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query evaluation logic internal server error response has a 5xx status code
func (o *QueryEvaluationLogicInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query evaluation logic internal server error response a status code equal to that given
func (o *QueryEvaluationLogicInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryEvaluationLogicInternalServerError) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicInternalServerError ", 500)
}

func (o *QueryEvaluationLogicInternalServerError) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogicInternalServerError ", 500)
}

func (o *QueryEvaluationLogicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewQueryEvaluationLogicDefault creates a QueryEvaluationLogicDefault with default headers values
func NewQueryEvaluationLogicDefault(code int) *QueryEvaluationLogicDefault {
	return &QueryEvaluationLogicDefault{
		_statusCode: code,
	}
}

/*
QueryEvaluationLogicDefault describes a response with status code -1, with default header values.

OK
*/
type QueryEvaluationLogicDefault struct {
	_statusCode int

	Payload *models.DomainSPAPIQueryResponse
}

// Code gets the status code for the query evaluation logic default response
func (o *QueryEvaluationLogicDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query evaluation logic default response has a 2xx status code
func (o *QueryEvaluationLogicDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query evaluation logic default response has a 3xx status code
func (o *QueryEvaluationLogicDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query evaluation logic default response has a 4xx status code
func (o *QueryEvaluationLogicDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query evaluation logic default response has a 5xx status code
func (o *QueryEvaluationLogicDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query evaluation logic default response a status code equal to that given
func (o *QueryEvaluationLogicDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryEvaluationLogicDefault) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogic default  %+v", o._statusCode, o.Payload)
}

func (o *QueryEvaluationLogicDefault) String() string {
	return fmt.Sprintf("[GET /spotlight/queries/evaluation-logic/v1][%d] queryEvaluationLogic default  %+v", o._statusCode, o.Payload)
}

func (o *QueryEvaluationLogicDefault) GetPayload() *models.DomainSPAPIQueryResponse {
	return o.Payload
}

func (o *QueryEvaluationLogicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainSPAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
