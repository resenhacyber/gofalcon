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

// TokensUpdateReader is a Reader for the TokensUpdate structure.
type TokensUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokensUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokensUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTokensUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTokensUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTokensUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTokensUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTokensUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTokensUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTokensUpdateOK creates a TokensUpdateOK with default headers values
func NewTokensUpdateOK() *TokensUpdateOK {
	return &TokensUpdateOK{}
}

/*TokensUpdateOK handles this case with default header values.

OK
*/
type TokensUpdateOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *TokensUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateOK  %+v", 200, o.Payload)
}

func (o *TokensUpdateOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *TokensUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokensUpdateBadRequest creates a TokensUpdateBadRequest with default headers values
func NewTokensUpdateBadRequest() *TokensUpdateBadRequest {
	return &TokensUpdateBadRequest{}
}

/*TokensUpdateBadRequest handles this case with default header values.

Bad Request
*/
type TokensUpdateBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *TokensUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TokensUpdateBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensUpdateForbidden creates a TokensUpdateForbidden with default headers values
func NewTokensUpdateForbidden() *TokensUpdateForbidden {
	return &TokensUpdateForbidden{}
}

/*TokensUpdateForbidden handles this case with default header values.

Forbidden
*/
type TokensUpdateForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *TokensUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateForbidden  %+v", 403, o.Payload)
}

func (o *TokensUpdateForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensUpdateNotFound creates a TokensUpdateNotFound with default headers values
func NewTokensUpdateNotFound() *TokensUpdateNotFound {
	return &TokensUpdateNotFound{}
}

/*TokensUpdateNotFound handles this case with default header values.

Not Found
*/
type TokensUpdateNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *TokensUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateNotFound  %+v", 404, o.Payload)
}

func (o *TokensUpdateNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *TokensUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokensUpdateTooManyRequests creates a TokensUpdateTooManyRequests with default headers values
func NewTokensUpdateTooManyRequests() *TokensUpdateTooManyRequests {
	return &TokensUpdateTooManyRequests{}
}

/*TokensUpdateTooManyRequests handles this case with default header values.

Too Many Requests
*/
type TokensUpdateTooManyRequests struct {
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

func (o *TokensUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *TokensUpdateTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensUpdateInternalServerError creates a TokensUpdateInternalServerError with default headers values
func NewTokensUpdateInternalServerError() *TokensUpdateInternalServerError {
	return &TokensUpdateInternalServerError{}
}

/*TokensUpdateInternalServerError handles this case with default header values.

Internal Server Error
*/
type TokensUpdateInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *TokensUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokensUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *TokensUpdateInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensUpdateDefault creates a TokensUpdateDefault with default headers values
func NewTokensUpdateDefault(code int) *TokensUpdateDefault {
	return &TokensUpdateDefault{
		_statusCode: code,
	}
}

/*TokensUpdateDefault handles this case with default header values.

OK
*/
type TokensUpdateDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the tokens update default response
func (o *TokensUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TokensUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/tokens/v1][%d] tokens-update default  %+v", o._statusCode, o.Payload)
}

func (o *TokensUpdateDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *TokensUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
