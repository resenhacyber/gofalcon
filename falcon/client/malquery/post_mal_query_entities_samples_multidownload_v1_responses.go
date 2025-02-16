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

// PostMalQueryEntitiesSamplesMultidownloadV1Reader is a Reader for the PostMalQueryEntitiesSamplesMultidownloadV1 structure.
type PostMalQueryEntitiesSamplesMultidownloadV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostMalQueryEntitiesSamplesMultidownloadV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1OK creates a PostMalQueryEntitiesSamplesMultidownloadV1OK with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1OK() *PostMalQueryEntitiesSamplesMultidownloadV1OK {
	return &PostMalQueryEntitiesSamplesMultidownloadV1OK{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1OK describes a response with status code 200, with default header values.

OK
*/
type PostMalQueryEntitiesSamplesMultidownloadV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 o k response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 o k response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 o k response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post mal query entities samples multidownload v1 o k response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 o k response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1OK  %+v", 200, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1OK  %+v", 200, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1BadRequest creates a PostMalQueryEntitiesSamplesMultidownloadV1BadRequest with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1BadRequest() *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest {
	return &PostMalQueryEntitiesSamplesMultidownloadV1BadRequest{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostMalQueryEntitiesSamplesMultidownloadV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 bad request response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 bad request response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 bad request response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query entities samples multidownload v1 bad request response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 bad request response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1BadRequest  %+v", 400, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1Unauthorized creates a PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1Unauthorized() *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized {
	return &PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 unauthorized response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 unauthorized response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 unauthorized response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query entities samples multidownload v1 unauthorized response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 unauthorized response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1Unauthorized  %+v", 401, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1Unauthorized  %+v", 401, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostMalQueryEntitiesSamplesMultidownloadV1Forbidden creates a PostMalQueryEntitiesSamplesMultidownloadV1Forbidden with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1Forbidden() *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden {
	return &PostMalQueryEntitiesSamplesMultidownloadV1Forbidden{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostMalQueryEntitiesSamplesMultidownloadV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 forbidden response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 forbidden response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 forbidden response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query entities samples multidownload v1 forbidden response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 forbidden response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1Forbidden  %+v", 403, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPostMalQueryEntitiesSamplesMultidownloadV1NotFound creates a PostMalQueryEntitiesSamplesMultidownloadV1NotFound with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1NotFound() *PostMalQueryEntitiesSamplesMultidownloadV1NotFound {
	return &PostMalQueryEntitiesSamplesMultidownloadV1NotFound{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostMalQueryEntitiesSamplesMultidownloadV1NotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 not found response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 not found response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 not found response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query entities samples multidownload v1 not found response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 not found response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1NotFound  %+v", 404, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1NotFound  %+v", 404, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests creates a PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests() *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests {
	return &PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 too many requests response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 too many requests response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 too many requests response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post mal query entities samples multidownload v1 too many requests response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post mal query entities samples multidownload v1 too many requests response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1InternalServerError creates a PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1InternalServerError() *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError {
	return &PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError{}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryExternalQueryResponse
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 internal server error response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 internal server error response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post mal query entities samples multidownload v1 internal server error response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post mal query entities samples multidownload v1 internal server error response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post mal query entities samples multidownload v1 internal server error response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] postMalQueryEntitiesSamplesMultidownloadV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMalQueryEntitiesSamplesMultidownloadV1Default creates a PostMalQueryEntitiesSamplesMultidownloadV1Default with default headers values
func NewPostMalQueryEntitiesSamplesMultidownloadV1Default(code int) *PostMalQueryEntitiesSamplesMultidownloadV1Default {
	return &PostMalQueryEntitiesSamplesMultidownloadV1Default{
		_statusCode: code,
	}
}

/*
PostMalQueryEntitiesSamplesMultidownloadV1Default describes a response with status code -1, with default header values.

OK
*/
type PostMalQueryEntitiesSamplesMultidownloadV1Default struct {
	_statusCode int

	Payload *models.MalqueryExternalQueryResponse
}

// Code gets the status code for the post mal query entities samples multidownload v1 default response
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post mal query entities samples multidownload v1 default response has a 2xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post mal query entities samples multidownload v1 default response has a 3xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post mal query entities samples multidownload v1 default response has a 4xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post mal query entities samples multidownload v1 default response has a 5xx status code
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post mal query entities samples multidownload v1 default response a status code equal to that given
func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) Error() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] PostMalQueryEntitiesSamplesMultidownloadV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) String() string {
	return fmt.Sprintf("[POST /malquery/entities/samples-multidownload/v1][%d] PostMalQueryEntitiesSamplesMultidownloadV1 default  %+v", o._statusCode, o.Payload)
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) GetPayload() *models.MalqueryExternalQueryResponse {
	return o.Payload
}

func (o *PostMalQueryEntitiesSamplesMultidownloadV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MalqueryExternalQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
