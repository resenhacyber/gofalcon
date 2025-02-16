// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// GetRTResponsePoliciesReader is a Reader for the GetRTResponsePolicies structure.
type GetRTResponsePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRTResponsePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRTResponsePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRTResponsePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRTResponsePoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRTResponsePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRTResponsePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRTResponsePoliciesOK creates a GetRTResponsePoliciesOK with default headers values
func NewGetRTResponsePoliciesOK() *GetRTResponsePoliciesOK {
	return &GetRTResponsePoliciesOK{}
}

/*
GetRTResponsePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetRTResponsePoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this get r t response policies o k response has a 2xx status code
func (o *GetRTResponsePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get r t response policies o k response has a 3xx status code
func (o *GetRTResponsePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get r t response policies o k response has a 4xx status code
func (o *GetRTResponsePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get r t response policies o k response has a 5xx status code
func (o *GetRTResponsePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get r t response policies o k response a status code equal to that given
func (o *GetRTResponsePoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRTResponsePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetRTResponsePoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetRTResponsePoliciesOK) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *GetRTResponsePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRTResponsePoliciesForbidden creates a GetRTResponsePoliciesForbidden with default headers values
func NewGetRTResponsePoliciesForbidden() *GetRTResponsePoliciesForbidden {
	return &GetRTResponsePoliciesForbidden{}
}

/*
GetRTResponsePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRTResponsePoliciesForbidden struct {

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

// IsSuccess returns true when this get r t response policies forbidden response has a 2xx status code
func (o *GetRTResponsePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get r t response policies forbidden response has a 3xx status code
func (o *GetRTResponsePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get r t response policies forbidden response has a 4xx status code
func (o *GetRTResponsePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get r t response policies forbidden response has a 5xx status code
func (o *GetRTResponsePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get r t response policies forbidden response a status code equal to that given
func (o *GetRTResponsePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRTResponsePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetRTResponsePoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetRTResponsePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetRTResponsePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRTResponsePoliciesNotFound creates a GetRTResponsePoliciesNotFound with default headers values
func NewGetRTResponsePoliciesNotFound() *GetRTResponsePoliciesNotFound {
	return &GetRTResponsePoliciesNotFound{}
}

/*
GetRTResponsePoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRTResponsePoliciesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this get r t response policies not found response has a 2xx status code
func (o *GetRTResponsePoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get r t response policies not found response has a 3xx status code
func (o *GetRTResponsePoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get r t response policies not found response has a 4xx status code
func (o *GetRTResponsePoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get r t response policies not found response has a 5xx status code
func (o *GetRTResponsePoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get r t response policies not found response a status code equal to that given
func (o *GetRTResponsePoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRTResponsePoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetRTResponsePoliciesNotFound) String() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetRTResponsePoliciesNotFound) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *GetRTResponsePoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRTResponsePoliciesTooManyRequests creates a GetRTResponsePoliciesTooManyRequests with default headers values
func NewGetRTResponsePoliciesTooManyRequests() *GetRTResponsePoliciesTooManyRequests {
	return &GetRTResponsePoliciesTooManyRequests{}
}

/*
GetRTResponsePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRTResponsePoliciesTooManyRequests struct {

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

// IsSuccess returns true when this get r t response policies too many requests response has a 2xx status code
func (o *GetRTResponsePoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get r t response policies too many requests response has a 3xx status code
func (o *GetRTResponsePoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get r t response policies too many requests response has a 4xx status code
func (o *GetRTResponsePoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get r t response policies too many requests response has a 5xx status code
func (o *GetRTResponsePoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get r t response policies too many requests response a status code equal to that given
func (o *GetRTResponsePoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRTResponsePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRTResponsePoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRTResponsePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRTResponsePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRTResponsePoliciesInternalServerError creates a GetRTResponsePoliciesInternalServerError with default headers values
func NewGetRTResponsePoliciesInternalServerError() *GetRTResponsePoliciesInternalServerError {
	return &GetRTResponsePoliciesInternalServerError{}
}

/*
GetRTResponsePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRTResponsePoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

// IsSuccess returns true when this get r t response policies internal server error response has a 2xx status code
func (o *GetRTResponsePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get r t response policies internal server error response has a 3xx status code
func (o *GetRTResponsePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get r t response policies internal server error response has a 4xx status code
func (o *GetRTResponsePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get r t response policies internal server error response has a 5xx status code
func (o *GetRTResponsePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get r t response policies internal server error response a status code equal to that given
func (o *GetRTResponsePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRTResponsePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRTResponsePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/entities/response/v1][%d] getRTResponsePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRTResponsePoliciesInternalServerError) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *GetRTResponsePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
