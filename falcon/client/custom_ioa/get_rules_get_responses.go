// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// GetRulesGetReader is a Reader for the GetRulesGet structure.
type GetRulesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRulesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRulesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRulesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRulesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRulesGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRulesGetOK creates a GetRulesGetOK with default headers values
func NewGetRulesGetOK() *GetRulesGetOK {
	return &GetRulesGetOK{}
}

/*
GetRulesGetOK describes a response with status code 200, with default header values.

OK
*/
type GetRulesGetOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRulesResponse
}

// IsSuccess returns true when this get rules get o k response has a 2xx status code
func (o *GetRulesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rules get o k response has a 3xx status code
func (o *GetRulesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules get o k response has a 4xx status code
func (o *GetRulesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rules get o k response has a 5xx status code
func (o *GetRulesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules get o k response a status code equal to that given
func (o *GetRulesGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRulesGetOK) Error() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetOK  %+v", 200, o.Payload)
}

func (o *GetRulesGetOK) String() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetOK  %+v", 200, o.Payload)
}

func (o *GetRulesGetOK) GetPayload() *models.APIRulesResponse {
	return o.Payload
}

func (o *GetRulesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesGetForbidden creates a GetRulesGetForbidden with default headers values
func NewGetRulesGetForbidden() *GetRulesGetForbidden {
	return &GetRulesGetForbidden{}
}

/*
GetRulesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRulesGetForbidden struct {

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

// IsSuccess returns true when this get rules get forbidden response has a 2xx status code
func (o *GetRulesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules get forbidden response has a 3xx status code
func (o *GetRulesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules get forbidden response has a 4xx status code
func (o *GetRulesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules get forbidden response has a 5xx status code
func (o *GetRulesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules get forbidden response a status code equal to that given
func (o *GetRulesGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRulesGetForbidden) Error() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetForbidden  %+v", 403, o.Payload)
}

func (o *GetRulesGetForbidden) String() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetForbidden  %+v", 403, o.Payload)
}

func (o *GetRulesGetForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRulesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRulesGetNotFound creates a GetRulesGetNotFound with default headers values
func NewGetRulesGetNotFound() *GetRulesGetNotFound {
	return &GetRulesGetNotFound{}
}

/*
GetRulesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRulesGetNotFound struct {

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

// IsSuccess returns true when this get rules get not found response has a 2xx status code
func (o *GetRulesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules get not found response has a 3xx status code
func (o *GetRulesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules get not found response has a 4xx status code
func (o *GetRulesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules get not found response has a 5xx status code
func (o *GetRulesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules get not found response a status code equal to that given
func (o *GetRulesGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRulesGetNotFound) Error() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetNotFound  %+v", 404, o.Payload)
}

func (o *GetRulesGetNotFound) String() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetNotFound  %+v", 404, o.Payload)
}

func (o *GetRulesGetNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRulesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRulesGetTooManyRequests creates a GetRulesGetTooManyRequests with default headers values
func NewGetRulesGetTooManyRequests() *GetRulesGetTooManyRequests {
	return &GetRulesGetTooManyRequests{}
}

/*
GetRulesGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRulesGetTooManyRequests struct {

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

// IsSuccess returns true when this get rules get too many requests response has a 2xx status code
func (o *GetRulesGetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules get too many requests response has a 3xx status code
func (o *GetRulesGetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules get too many requests response has a 4xx status code
func (o *GetRulesGetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules get too many requests response has a 5xx status code
func (o *GetRulesGetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules get too many requests response a status code equal to that given
func (o *GetRulesGetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRulesGetTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesGetTooManyRequests) String() string {
	return fmt.Sprintf("[POST /ioarules/entities/rules/GET/v1][%d] getRulesGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesGetTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRulesGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
