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

// GetPlatformsMixin0Reader is a Reader for the GetPlatformsMixin0 structure.
type GetPlatformsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlatformsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPlatformsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPlatformsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPlatformsMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformsMixin0OK creates a GetPlatformsMixin0OK with default headers values
func NewGetPlatformsMixin0OK() *GetPlatformsMixin0OK {
	return &GetPlatformsMixin0OK{}
}

/*
GetPlatformsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type GetPlatformsMixin0OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIPlatformsResponse
}

// IsSuccess returns true when this get platforms mixin0 o k response has a 2xx status code
func (o *GetPlatformsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get platforms mixin0 o k response has a 3xx status code
func (o *GetPlatformsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms mixin0 o k response has a 4xx status code
func (o *GetPlatformsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get platforms mixin0 o k response has a 5xx status code
func (o *GetPlatformsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms mixin0 o k response a status code equal to that given
func (o *GetPlatformsMixin0OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPlatformsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetPlatformsMixin0OK) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetPlatformsMixin0OK) GetPayload() *models.APIPlatformsResponse {
	return o.Payload
}

func (o *GetPlatformsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIPlatformsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformsMixin0Forbidden creates a GetPlatformsMixin0Forbidden with default headers values
func NewGetPlatformsMixin0Forbidden() *GetPlatformsMixin0Forbidden {
	return &GetPlatformsMixin0Forbidden{}
}

/*
GetPlatformsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPlatformsMixin0Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get platforms mixin0 forbidden response has a 2xx status code
func (o *GetPlatformsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get platforms mixin0 forbidden response has a 3xx status code
func (o *GetPlatformsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms mixin0 forbidden response has a 4xx status code
func (o *GetPlatformsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get platforms mixin0 forbidden response has a 5xx status code
func (o *GetPlatformsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms mixin0 forbidden response a status code equal to that given
func (o *GetPlatformsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPlatformsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetPlatformsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetPlatformsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPlatformsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPlatformsMixin0TooManyRequests creates a GetPlatformsMixin0TooManyRequests with default headers values
func NewGetPlatformsMixin0TooManyRequests() *GetPlatformsMixin0TooManyRequests {
	return &GetPlatformsMixin0TooManyRequests{}
}

/*
GetPlatformsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetPlatformsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this get platforms mixin0 too many requests response has a 2xx status code
func (o *GetPlatformsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get platforms mixin0 too many requests response has a 3xx status code
func (o *GetPlatformsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get platforms mixin0 too many requests response has a 4xx status code
func (o *GetPlatformsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get platforms mixin0 too many requests response has a 5xx status code
func (o *GetPlatformsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get platforms mixin0 too many requests response a status code equal to that given
func (o *GetPlatformsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPlatformsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPlatformsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] getPlatformsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPlatformsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPlatformsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPlatformsMixin0Default creates a GetPlatformsMixin0Default with default headers values
func NewGetPlatformsMixin0Default(code int) *GetPlatformsMixin0Default {
	return &GetPlatformsMixin0Default{
		_statusCode: code,
	}
}

/*
GetPlatformsMixin0Default describes a response with status code -1, with default header values.

OK
*/
type GetPlatformsMixin0Default struct {
	_statusCode int

	Payload *models.APIPlatformsResponse
}

// Code gets the status code for the get platforms mixin0 default response
func (o *GetPlatformsMixin0Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get platforms mixin0 default response has a 2xx status code
func (o *GetPlatformsMixin0Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get platforms mixin0 default response has a 3xx status code
func (o *GetPlatformsMixin0Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get platforms mixin0 default response has a 4xx status code
func (o *GetPlatformsMixin0Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get platforms mixin0 default response has a 5xx status code
func (o *GetPlatformsMixin0Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get platforms mixin0 default response a status code equal to that given
func (o *GetPlatformsMixin0Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPlatformsMixin0Default) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] get-platformsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *GetPlatformsMixin0Default) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/platforms/v1][%d] get-platformsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *GetPlatformsMixin0Default) GetPayload() *models.APIPlatformsResponse {
	return o.Payload
}

func (o *GetPlatformsMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPlatformsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
