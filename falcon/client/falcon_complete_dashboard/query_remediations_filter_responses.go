// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// QueryRemediationsFilterReader is a Reader for the QueryRemediationsFilter structure.
type QueryRemediationsFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRemediationsFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRemediationsFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryRemediationsFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryRemediationsFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryRemediationsFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryRemediationsFilterOK creates a QueryRemediationsFilterOK with default headers values
func NewQueryRemediationsFilterOK() *QueryRemediationsFilterOK {
	return &QueryRemediationsFilterOK{}
}

/*
QueryRemediationsFilterOK describes a response with status code 200, with default header values.

OK
*/
type QueryRemediationsFilterOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query remediations filter o k response has a 2xx status code
func (o *QueryRemediationsFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query remediations filter o k response has a 3xx status code
func (o *QueryRemediationsFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query remediations filter o k response has a 4xx status code
func (o *QueryRemediationsFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query remediations filter o k response has a 5xx status code
func (o *QueryRemediationsFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query remediations filter o k response a status code equal to that given
func (o *QueryRemediationsFilterOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryRemediationsFilterOK) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterOK  %+v", 200, o.Payload)
}

func (o *QueryRemediationsFilterOK) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterOK  %+v", 200, o.Payload)
}

func (o *QueryRemediationsFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryRemediationsFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRemediationsFilterForbidden creates a QueryRemediationsFilterForbidden with default headers values
func NewQueryRemediationsFilterForbidden() *QueryRemediationsFilterForbidden {
	return &QueryRemediationsFilterForbidden{}
}

/*
QueryRemediationsFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryRemediationsFilterForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query remediations filter forbidden response has a 2xx status code
func (o *QueryRemediationsFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query remediations filter forbidden response has a 3xx status code
func (o *QueryRemediationsFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query remediations filter forbidden response has a 4xx status code
func (o *QueryRemediationsFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query remediations filter forbidden response has a 5xx status code
func (o *QueryRemediationsFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query remediations filter forbidden response a status code equal to that given
func (o *QueryRemediationsFilterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryRemediationsFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryRemediationsFilterForbidden) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryRemediationsFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRemediationsFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRemediationsFilterTooManyRequests creates a QueryRemediationsFilterTooManyRequests with default headers values
func NewQueryRemediationsFilterTooManyRequests() *QueryRemediationsFilterTooManyRequests {
	return &QueryRemediationsFilterTooManyRequests{}
}

/*
QueryRemediationsFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryRemediationsFilterTooManyRequests struct {

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

// IsSuccess returns true when this query remediations filter too many requests response has a 2xx status code
func (o *QueryRemediationsFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query remediations filter too many requests response has a 3xx status code
func (o *QueryRemediationsFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query remediations filter too many requests response has a 4xx status code
func (o *QueryRemediationsFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query remediations filter too many requests response has a 5xx status code
func (o *QueryRemediationsFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query remediations filter too many requests response a status code equal to that given
func (o *QueryRemediationsFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryRemediationsFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRemediationsFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] queryRemediationsFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRemediationsFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRemediationsFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRemediationsFilterDefault creates a QueryRemediationsFilterDefault with default headers values
func NewQueryRemediationsFilterDefault(code int) *QueryRemediationsFilterDefault {
	return &QueryRemediationsFilterDefault{
		_statusCode: code,
	}
}

/*
QueryRemediationsFilterDefault describes a response with status code -1, with default header values.

OK
*/
type QueryRemediationsFilterDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query remediations filter default response
func (o *QueryRemediationsFilterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query remediations filter default response has a 2xx status code
func (o *QueryRemediationsFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query remediations filter default response has a 3xx status code
func (o *QueryRemediationsFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query remediations filter default response has a 4xx status code
func (o *QueryRemediationsFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query remediations filter default response has a 5xx status code
func (o *QueryRemediationsFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query remediations filter default response a status code equal to that given
func (o *QueryRemediationsFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryRemediationsFilterDefault) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] QueryRemediationsFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryRemediationsFilterDefault) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/remediations/v1][%d] QueryRemediationsFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryRemediationsFilterDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryRemediationsFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
