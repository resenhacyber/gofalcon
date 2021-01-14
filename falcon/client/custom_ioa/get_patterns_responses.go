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

// GetPatternsReader is a Reader for the GetPatterns structure.
type GetPatternsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatternsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatternsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPatternsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPatternsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPatternsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPatternsOK creates a GetPatternsOK with default headers values
func NewGetPatternsOK() *GetPatternsOK {
	return &GetPatternsOK{}
}

/*GetPatternsOK handles this case with default header values.

OK
*/
type GetPatternsOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIPatternsResponse
}

func (o *GetPatternsOK) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/pattern-severities/v1][%d] getPatternsOK  %+v", 200, o.Payload)
}

func (o *GetPatternsOK) GetPayload() *models.APIPatternsResponse {
	return o.Payload
}

func (o *GetPatternsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.APIPatternsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatternsForbidden creates a GetPatternsForbidden with default headers values
func NewGetPatternsForbidden() *GetPatternsForbidden {
	return &GetPatternsForbidden{}
}

/*GetPatternsForbidden handles this case with default header values.

Forbidden
*/
type GetPatternsForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetPatternsForbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/pattern-severities/v1][%d] getPatternsForbidden  %+v", 403, o.Payload)
}

func (o *GetPatternsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPatternsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatternsTooManyRequests creates a GetPatternsTooManyRequests with default headers values
func NewGetPatternsTooManyRequests() *GetPatternsTooManyRequests {
	return &GetPatternsTooManyRequests{}
}

/*GetPatternsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetPatternsTooManyRequests struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
	/*Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetPatternsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/pattern-severities/v1][%d] getPatternsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPatternsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPatternsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-RetryAfter
	xRateLimitRetryAfter, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-RetryAfter"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", response.GetHeader("X-RateLimit-RetryAfter"))
	}
	o.XRateLimitRetryAfter = xRateLimitRetryAfter

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatternsDefault creates a GetPatternsDefault with default headers values
func NewGetPatternsDefault(code int) *GetPatternsDefault {
	return &GetPatternsDefault{
		_statusCode: code,
	}
}

/*GetPatternsDefault handles this case with default header values.

OK
*/
type GetPatternsDefault struct {
	_statusCode int

	Payload *models.APIPatternsResponse
}

// Code gets the status code for the get patterns default response
func (o *GetPatternsDefault) Code() int {
	return o._statusCode
}

func (o *GetPatternsDefault) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/pattern-severities/v1][%d] get-patterns default  %+v", o._statusCode, o.Payload)
}

func (o *GetPatternsDefault) GetPayload() *models.APIPatternsResponse {
	return o.Payload
}

func (o *GetPatternsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPatternsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
