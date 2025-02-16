// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// GetMalQueryMetadataV1Reader is a Reader for the GetMalQueryMetadataV1 structure.
type GetMalQueryMetadataV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMalQueryMetadataV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMalQueryMetadataV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMalQueryMetadataV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMalQueryMetadataV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMalQueryMetadataV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMalQueryMetadataV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMalQueryMetadataV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMalQueryMetadataV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMalQueryMetadataV1OK creates a GetMalQueryMetadataV1OK with default headers values
func NewGetMalQueryMetadataV1OK() *GetMalQueryMetadataV1OK {
	return &GetMalQueryMetadataV1OK{}
}

/*
GetMalQueryMetadataV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMalQueryMetadataV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalquerySampleMetadataResponse
}

// IsSuccess returns true when this get mal query metadata v1 o k response has a 2xx status code
func (o *GetMalQueryMetadataV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mal query metadata v1 o k response has a 3xx status code
func (o *GetMalQueryMetadataV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 o k response has a 4xx status code
func (o *GetMalQueryMetadataV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query metadata v1 o k response has a 5xx status code
func (o *GetMalQueryMetadataV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query metadata v1 o k response a status code equal to that given
func (o *GetMalQueryMetadataV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMalQueryMetadataV1OK) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryMetadataV1OK) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryMetadataV1OK) GetPayload() *models.MalquerySampleMetadataResponse {
	return o.Payload
}

func (o *GetMalQueryMetadataV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalquerySampleMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryMetadataV1BadRequest creates a GetMalQueryMetadataV1BadRequest with default headers values
func NewGetMalQueryMetadataV1BadRequest() *GetMalQueryMetadataV1BadRequest {
	return &GetMalQueryMetadataV1BadRequest{}
}

/*
GetMalQueryMetadataV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMalQueryMetadataV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalquerySampleMetadataResponse
}

// IsSuccess returns true when this get mal query metadata v1 bad request response has a 2xx status code
func (o *GetMalQueryMetadataV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query metadata v1 bad request response has a 3xx status code
func (o *GetMalQueryMetadataV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 bad request response has a 4xx status code
func (o *GetMalQueryMetadataV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query metadata v1 bad request response has a 5xx status code
func (o *GetMalQueryMetadataV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query metadata v1 bad request response a status code equal to that given
func (o *GetMalQueryMetadataV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetMalQueryMetadataV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryMetadataV1BadRequest) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryMetadataV1BadRequest) GetPayload() *models.MalquerySampleMetadataResponse {
	return o.Payload
}

func (o *GetMalQueryMetadataV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalquerySampleMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryMetadataV1Unauthorized creates a GetMalQueryMetadataV1Unauthorized with default headers values
func NewGetMalQueryMetadataV1Unauthorized() *GetMalQueryMetadataV1Unauthorized {
	return &GetMalQueryMetadataV1Unauthorized{}
}

/*
GetMalQueryMetadataV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetMalQueryMetadataV1Unauthorized struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get mal query metadata v1 unauthorized response has a 2xx status code
func (o *GetMalQueryMetadataV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query metadata v1 unauthorized response has a 3xx status code
func (o *GetMalQueryMetadataV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 unauthorized response has a 4xx status code
func (o *GetMalQueryMetadataV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query metadata v1 unauthorized response has a 5xx status code
func (o *GetMalQueryMetadataV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query metadata v1 unauthorized response a status code equal to that given
func (o *GetMalQueryMetadataV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetMalQueryMetadataV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryMetadataV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryMetadataV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryMetadataV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryMetadataV1Forbidden creates a GetMalQueryMetadataV1Forbidden with default headers values
func NewGetMalQueryMetadataV1Forbidden() *GetMalQueryMetadataV1Forbidden {
	return &GetMalQueryMetadataV1Forbidden{}
}

/*
GetMalQueryMetadataV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMalQueryMetadataV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get mal query metadata v1 forbidden response has a 2xx status code
func (o *GetMalQueryMetadataV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query metadata v1 forbidden response has a 3xx status code
func (o *GetMalQueryMetadataV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 forbidden response has a 4xx status code
func (o *GetMalQueryMetadataV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query metadata v1 forbidden response has a 5xx status code
func (o *GetMalQueryMetadataV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query metadata v1 forbidden response a status code equal to that given
func (o *GetMalQueryMetadataV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetMalQueryMetadataV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryMetadataV1Forbidden) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryMetadataV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryMetadataV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryMetadataV1TooManyRequests creates a GetMalQueryMetadataV1TooManyRequests with default headers values
func NewGetMalQueryMetadataV1TooManyRequests() *GetMalQueryMetadataV1TooManyRequests {
	return &GetMalQueryMetadataV1TooManyRequests{}
}

/*
GetMalQueryMetadataV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMalQueryMetadataV1TooManyRequests struct {

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

// IsSuccess returns true when this get mal query metadata v1 too many requests response has a 2xx status code
func (o *GetMalQueryMetadataV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query metadata v1 too many requests response has a 3xx status code
func (o *GetMalQueryMetadataV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 too many requests response has a 4xx status code
func (o *GetMalQueryMetadataV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mal query metadata v1 too many requests response has a 5xx status code
func (o *GetMalQueryMetadataV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get mal query metadata v1 too many requests response a status code equal to that given
func (o *GetMalQueryMetadataV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetMalQueryMetadataV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryMetadataV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryMetadataV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMalQueryMetadataV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryMetadataV1InternalServerError creates a GetMalQueryMetadataV1InternalServerError with default headers values
func NewGetMalQueryMetadataV1InternalServerError() *GetMalQueryMetadataV1InternalServerError {
	return &GetMalQueryMetadataV1InternalServerError{}
}

/*
GetMalQueryMetadataV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMalQueryMetadataV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalquerySampleMetadataResponse
}

// IsSuccess returns true when this get mal query metadata v1 internal server error response has a 2xx status code
func (o *GetMalQueryMetadataV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mal query metadata v1 internal server error response has a 3xx status code
func (o *GetMalQueryMetadataV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mal query metadata v1 internal server error response has a 4xx status code
func (o *GetMalQueryMetadataV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mal query metadata v1 internal server error response has a 5xx status code
func (o *GetMalQueryMetadataV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get mal query metadata v1 internal server error response a status code equal to that given
func (o *GetMalQueryMetadataV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetMalQueryMetadataV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryMetadataV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] getMalQueryMetadataV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryMetadataV1InternalServerError) GetPayload() *models.MalquerySampleMetadataResponse {
	return o.Payload
}

func (o *GetMalQueryMetadataV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalquerySampleMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryMetadataV1Default creates a GetMalQueryMetadataV1Default with default headers values
func NewGetMalQueryMetadataV1Default(code int) *GetMalQueryMetadataV1Default {
	return &GetMalQueryMetadataV1Default{
		_statusCode: code,
	}
}

/*
GetMalQueryMetadataV1Default describes a response with status code -1, with default header values.

OK
*/
type GetMalQueryMetadataV1Default struct {
	_statusCode int

	Payload *models.MalquerySampleMetadataResponse
}

// Code gets the status code for the get mal query metadata v1 default response
func (o *GetMalQueryMetadataV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get mal query metadata v1 default response has a 2xx status code
func (o *GetMalQueryMetadataV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get mal query metadata v1 default response has a 3xx status code
func (o *GetMalQueryMetadataV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get mal query metadata v1 default response has a 4xx status code
func (o *GetMalQueryMetadataV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get mal query metadata v1 default response has a 5xx status code
func (o *GetMalQueryMetadataV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get mal query metadata v1 default response a status code equal to that given
func (o *GetMalQueryMetadataV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetMalQueryMetadataV1Default) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] GetMalQueryMetadataV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetMalQueryMetadataV1Default) String() string {
	return fmt.Sprintf("[GET /malquery/entities/metadata/v1][%d] GetMalQueryMetadataV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetMalQueryMetadataV1Default) GetPayload() *models.MalquerySampleMetadataResponse {
	return o.Payload
}

func (o *GetMalQueryMetadataV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MalquerySampleMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
