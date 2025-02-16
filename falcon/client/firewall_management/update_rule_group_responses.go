// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// UpdateRuleGroupReader is a Reader for the UpdateRuleGroup structure.
type UpdateRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRuleGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRuleGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRuleGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRuleGroupOK creates a UpdateRuleGroupOK with default headers values
func NewUpdateRuleGroupOK() *UpdateRuleGroupOK {
	return &UpdateRuleGroupOK{}
}

/*
UpdateRuleGroupOK describes a response with status code 200, with default header values.

OK
*/
type UpdateRuleGroupOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIQueryResponse
}

// IsSuccess returns true when this update rule group o k response has a 2xx status code
func (o *UpdateRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rule group o k response has a 3xx status code
func (o *UpdateRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group o k response has a 4xx status code
func (o *UpdateRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rule group o k response has a 5xx status code
func (o *UpdateRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group o k response a status code equal to that given
func (o *UpdateRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRuleGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupOK) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupOK) GetPayload() *models.FwmgrAPIQueryResponse {
	return o.Payload
}

func (o *UpdateRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupBadRequest creates a UpdateRuleGroupBadRequest with default headers values
func NewUpdateRuleGroupBadRequest() *UpdateRuleGroupBadRequest {
	return &UpdateRuleGroupBadRequest{}
}

/*
UpdateRuleGroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRuleGroupBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaReplyMetaOnly
}

// IsSuccess returns true when this update rule group bad request response has a 2xx status code
func (o *UpdateRuleGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group bad request response has a 3xx status code
func (o *UpdateRuleGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group bad request response has a 4xx status code
func (o *UpdateRuleGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group bad request response has a 5xx status code
func (o *UpdateRuleGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group bad request response a status code equal to that given
func (o *UpdateRuleGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRuleGroupBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupBadRequest) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupBadRequest) GetPayload() *models.FwmgrMsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupForbidden creates a UpdateRuleGroupForbidden with default headers values
func NewUpdateRuleGroupForbidden() *UpdateRuleGroupForbidden {
	return &UpdateRuleGroupForbidden{}
}

/*
UpdateRuleGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRuleGroupForbidden struct {

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

// IsSuccess returns true when this update rule group forbidden response has a 2xx status code
func (o *UpdateRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group forbidden response has a 3xx status code
func (o *UpdateRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group forbidden response has a 4xx status code
func (o *UpdateRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group forbidden response has a 5xx status code
func (o *UpdateRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group forbidden response a status code equal to that given
func (o *UpdateRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRuleGroupForbidden) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupForbidden) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupTooManyRequests creates a UpdateRuleGroupTooManyRequests with default headers values
func NewUpdateRuleGroupTooManyRequests() *UpdateRuleGroupTooManyRequests {
	return &UpdateRuleGroupTooManyRequests{}
}

/*
UpdateRuleGroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRuleGroupTooManyRequests struct {

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

// IsSuccess returns true when this update rule group too many requests response has a 2xx status code
func (o *UpdateRuleGroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group too many requests response has a 3xx status code
func (o *UpdateRuleGroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group too many requests response has a 4xx status code
func (o *UpdateRuleGroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group too many requests response has a 5xx status code
func (o *UpdateRuleGroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group too many requests response a status code equal to that given
func (o *UpdateRuleGroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateRuleGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] updateRuleGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupDefault creates a UpdateRuleGroupDefault with default headers values
func NewUpdateRuleGroupDefault(code int) *UpdateRuleGroupDefault {
	return &UpdateRuleGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateRuleGroupDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateRuleGroupDefault struct {
	_statusCode int

	Payload *models.FwmgrAPIQueryResponse
}

// Code gets the status code for the update rule group default response
func (o *UpdateRuleGroupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update rule group default response has a 2xx status code
func (o *UpdateRuleGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update rule group default response has a 3xx status code
func (o *UpdateRuleGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update rule group default response has a 4xx status code
func (o *UpdateRuleGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update rule group default response has a 5xx status code
func (o *UpdateRuleGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update rule group default response a status code equal to that given
func (o *UpdateRuleGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateRuleGroupDefault) Error() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] update-rule-group default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupDefault) String() string {
	return fmt.Sprintf("[PATCH /fwmgr/entities/rule-groups/v1][%d] update-rule-group default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRuleGroupDefault) GetPayload() *models.FwmgrAPIQueryResponse {
	return o.Payload
}

func (o *UpdateRuleGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrAPIQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
