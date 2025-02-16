// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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

// AuditEventsQueryReader is a Reader for the AuditEventsQuery structure.
type AuditEventsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditEventsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuditEventsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuditEventsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuditEventsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAuditEventsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuditEventsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuditEventsQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditEventsQueryOK creates a AuditEventsQueryOK with default headers values
func NewAuditEventsQueryOK() *AuditEventsQueryOK {
	return &AuditEventsQueryOK{}
}

/*
AuditEventsQueryOK describes a response with status code 200, with default header values.

OK
*/
type AuditEventsQueryOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this audit events query o k response has a 2xx status code
func (o *AuditEventsQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this audit events query o k response has a 3xx status code
func (o *AuditEventsQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit events query o k response has a 4xx status code
func (o *AuditEventsQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this audit events query o k response has a 5xx status code
func (o *AuditEventsQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this audit events query o k response a status code equal to that given
func (o *AuditEventsQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuditEventsQueryOK) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryOK  %+v", 200, o.Payload)
}

func (o *AuditEventsQueryOK) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryOK  %+v", 200, o.Payload)
}

func (o *AuditEventsQueryOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *AuditEventsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditEventsQueryBadRequest creates a AuditEventsQueryBadRequest with default headers values
func NewAuditEventsQueryBadRequest() *AuditEventsQueryBadRequest {
	return &AuditEventsQueryBadRequest{}
}

/*
AuditEventsQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuditEventsQueryBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this audit events query bad request response has a 2xx status code
func (o *AuditEventsQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this audit events query bad request response has a 3xx status code
func (o *AuditEventsQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit events query bad request response has a 4xx status code
func (o *AuditEventsQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this audit events query bad request response has a 5xx status code
func (o *AuditEventsQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this audit events query bad request response a status code equal to that given
func (o *AuditEventsQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuditEventsQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *AuditEventsQueryBadRequest) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *AuditEventsQueryBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AuditEventsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAuditEventsQueryForbidden creates a AuditEventsQueryForbidden with default headers values
func NewAuditEventsQueryForbidden() *AuditEventsQueryForbidden {
	return &AuditEventsQueryForbidden{}
}

/*
AuditEventsQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuditEventsQueryForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this audit events query forbidden response has a 2xx status code
func (o *AuditEventsQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this audit events query forbidden response has a 3xx status code
func (o *AuditEventsQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit events query forbidden response has a 4xx status code
func (o *AuditEventsQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this audit events query forbidden response has a 5xx status code
func (o *AuditEventsQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this audit events query forbidden response a status code equal to that given
func (o *AuditEventsQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AuditEventsQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryForbidden  %+v", 403, o.Payload)
}

func (o *AuditEventsQueryForbidden) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryForbidden  %+v", 403, o.Payload)
}

func (o *AuditEventsQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AuditEventsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAuditEventsQueryTooManyRequests creates a AuditEventsQueryTooManyRequests with default headers values
func NewAuditEventsQueryTooManyRequests() *AuditEventsQueryTooManyRequests {
	return &AuditEventsQueryTooManyRequests{}
}

/*
AuditEventsQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AuditEventsQueryTooManyRequests struct {

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

// IsSuccess returns true when this audit events query too many requests response has a 2xx status code
func (o *AuditEventsQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this audit events query too many requests response has a 3xx status code
func (o *AuditEventsQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit events query too many requests response has a 4xx status code
func (o *AuditEventsQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this audit events query too many requests response has a 5xx status code
func (o *AuditEventsQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this audit events query too many requests response a status code equal to that given
func (o *AuditEventsQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AuditEventsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *AuditEventsQueryTooManyRequests) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *AuditEventsQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AuditEventsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAuditEventsQueryInternalServerError creates a AuditEventsQueryInternalServerError with default headers values
func NewAuditEventsQueryInternalServerError() *AuditEventsQueryInternalServerError {
	return &AuditEventsQueryInternalServerError{}
}

/*
AuditEventsQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AuditEventsQueryInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this audit events query internal server error response has a 2xx status code
func (o *AuditEventsQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this audit events query internal server error response has a 3xx status code
func (o *AuditEventsQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit events query internal server error response has a 4xx status code
func (o *AuditEventsQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this audit events query internal server error response has a 5xx status code
func (o *AuditEventsQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this audit events query internal server error response a status code equal to that given
func (o *AuditEventsQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuditEventsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *AuditEventsQueryInternalServerError) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] auditEventsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *AuditEventsQueryInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AuditEventsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAuditEventsQueryDefault creates a AuditEventsQueryDefault with default headers values
func NewAuditEventsQueryDefault(code int) *AuditEventsQueryDefault {
	return &AuditEventsQueryDefault{
		_statusCode: code,
	}
}

/*
AuditEventsQueryDefault describes a response with status code -1, with default header values.

OK
*/
type AuditEventsQueryDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the audit events query default response
func (o *AuditEventsQueryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this audit events query default response has a 2xx status code
func (o *AuditEventsQueryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this audit events query default response has a 3xx status code
func (o *AuditEventsQueryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this audit events query default response has a 4xx status code
func (o *AuditEventsQueryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this audit events query default response has a 5xx status code
func (o *AuditEventsQueryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this audit events query default response a status code equal to that given
func (o *AuditEventsQueryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AuditEventsQueryDefault) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] audit-events-query default  %+v", o._statusCode, o.Payload)
}

func (o *AuditEventsQueryDefault) String() string {
	return fmt.Sprintf("[GET /installation-tokens/queries/audit-events/v1][%d] audit-events-query default  %+v", o._statusCode, o.Payload)
}

func (o *AuditEventsQueryDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *AuditEventsQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
