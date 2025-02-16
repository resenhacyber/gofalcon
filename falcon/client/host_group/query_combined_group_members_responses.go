// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

// QueryCombinedGroupMembersReader is a Reader for the QueryCombinedGroupMembers structure.
type QueryCombinedGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedGroupMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedGroupMembersOK creates a QueryCombinedGroupMembersOK with default headers values
func NewQueryCombinedGroupMembersOK() *QueryCombinedGroupMembersOK {
	return &QueryCombinedGroupMembersOK{}
}

/*
QueryCombinedGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedGroupMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesHostGroupMembersV1
}

// IsSuccess returns true when this query combined group members o k response has a 2xx status code
func (o *QueryCombinedGroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined group members o k response has a 3xx status code
func (o *QueryCombinedGroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members o k response has a 4xx status code
func (o *QueryCombinedGroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined group members o k response has a 5xx status code
func (o *QueryCombinedGroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined group members o k response a status code equal to that given
func (o *QueryCombinedGroupMembersOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryCombinedGroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedGroupMembersOK) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedGroupMembersOK) GetPayload() *models.ResponsesHostGroupMembersV1 {
	return o.Payload
}

func (o *QueryCombinedGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesHostGroupMembersV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedGroupMembersBadRequest creates a QueryCombinedGroupMembersBadRequest with default headers values
func NewQueryCombinedGroupMembersBadRequest() *QueryCombinedGroupMembersBadRequest {
	return &QueryCombinedGroupMembersBadRequest{}
}

/*
QueryCombinedGroupMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedGroupMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesHostGroupMembersV1
}

// IsSuccess returns true when this query combined group members bad request response has a 2xx status code
func (o *QueryCombinedGroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined group members bad request response has a 3xx status code
func (o *QueryCombinedGroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members bad request response has a 4xx status code
func (o *QueryCombinedGroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined group members bad request response has a 5xx status code
func (o *QueryCombinedGroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined group members bad request response a status code equal to that given
func (o *QueryCombinedGroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryCombinedGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedGroupMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedGroupMembersBadRequest) GetPayload() *models.ResponsesHostGroupMembersV1 {
	return o.Payload
}

func (o *QueryCombinedGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesHostGroupMembersV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedGroupMembersForbidden creates a QueryCombinedGroupMembersForbidden with default headers values
func NewQueryCombinedGroupMembersForbidden() *QueryCombinedGroupMembersForbidden {
	return &QueryCombinedGroupMembersForbidden{}
}

/*
QueryCombinedGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedGroupMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query combined group members forbidden response has a 2xx status code
func (o *QueryCombinedGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined group members forbidden response has a 3xx status code
func (o *QueryCombinedGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members forbidden response has a 4xx status code
func (o *QueryCombinedGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined group members forbidden response has a 5xx status code
func (o *QueryCombinedGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined group members forbidden response a status code equal to that given
func (o *QueryCombinedGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *QueryCombinedGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedGroupMembersForbidden) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedGroupMembersNotFound creates a QueryCombinedGroupMembersNotFound with default headers values
func NewQueryCombinedGroupMembersNotFound() *QueryCombinedGroupMembersNotFound {
	return &QueryCombinedGroupMembersNotFound{}
}

/*
QueryCombinedGroupMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedGroupMembersNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesHostGroupMembersV1
}

// IsSuccess returns true when this query combined group members not found response has a 2xx status code
func (o *QueryCombinedGroupMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined group members not found response has a 3xx status code
func (o *QueryCombinedGroupMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members not found response has a 4xx status code
func (o *QueryCombinedGroupMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined group members not found response has a 5xx status code
func (o *QueryCombinedGroupMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined group members not found response a status code equal to that given
func (o *QueryCombinedGroupMembersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *QueryCombinedGroupMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedGroupMembersNotFound) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedGroupMembersNotFound) GetPayload() *models.ResponsesHostGroupMembersV1 {
	return o.Payload
}

func (o *QueryCombinedGroupMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesHostGroupMembersV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedGroupMembersTooManyRequests creates a QueryCombinedGroupMembersTooManyRequests with default headers values
func NewQueryCombinedGroupMembersTooManyRequests() *QueryCombinedGroupMembersTooManyRequests {
	return &QueryCombinedGroupMembersTooManyRequests{}
}

/*
QueryCombinedGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedGroupMembersTooManyRequests struct {

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

// IsSuccess returns true when this query combined group members too many requests response has a 2xx status code
func (o *QueryCombinedGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined group members too many requests response has a 3xx status code
func (o *QueryCombinedGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members too many requests response has a 4xx status code
func (o *QueryCombinedGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined group members too many requests response has a 5xx status code
func (o *QueryCombinedGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined group members too many requests response a status code equal to that given
func (o *QueryCombinedGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *QueryCombinedGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedGroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedGroupMembersInternalServerError creates a QueryCombinedGroupMembersInternalServerError with default headers values
func NewQueryCombinedGroupMembersInternalServerError() *QueryCombinedGroupMembersInternalServerError {
	return &QueryCombinedGroupMembersInternalServerError{}
}

/*
QueryCombinedGroupMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedGroupMembersInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesHostGroupMembersV1
}

// IsSuccess returns true when this query combined group members internal server error response has a 2xx status code
func (o *QueryCombinedGroupMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined group members internal server error response has a 3xx status code
func (o *QueryCombinedGroupMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined group members internal server error response has a 4xx status code
func (o *QueryCombinedGroupMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined group members internal server error response has a 5xx status code
func (o *QueryCombinedGroupMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined group members internal server error response a status code equal to that given
func (o *QueryCombinedGroupMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryCombinedGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedGroupMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedGroupMembersInternalServerError) GetPayload() *models.ResponsesHostGroupMembersV1 {
	return o.Payload
}

func (o *QueryCombinedGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesHostGroupMembersV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedGroupMembersDefault creates a QueryCombinedGroupMembersDefault with default headers values
func NewQueryCombinedGroupMembersDefault(code int) *QueryCombinedGroupMembersDefault {
	return &QueryCombinedGroupMembersDefault{
		_statusCode: code,
	}
}

/*
QueryCombinedGroupMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedGroupMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesHostGroupMembersV1
}

// Code gets the status code for the query combined group members default response
func (o *QueryCombinedGroupMembersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this query combined group members default response has a 2xx status code
func (o *QueryCombinedGroupMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query combined group members default response has a 3xx status code
func (o *QueryCombinedGroupMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query combined group members default response has a 4xx status code
func (o *QueryCombinedGroupMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query combined group members default response has a 5xx status code
func (o *QueryCombinedGroupMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query combined group members default response a status code equal to that given
func (o *QueryCombinedGroupMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QueryCombinedGroupMembersDefault) Error() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedGroupMembersDefault) String() string {
	return fmt.Sprintf("[GET /devices/combined/host-group-members/v1][%d] queryCombinedGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedGroupMembersDefault) GetPayload() *models.ResponsesHostGroupMembersV1 {
	return o.Payload
}

func (o *QueryCombinedGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesHostGroupMembersV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
