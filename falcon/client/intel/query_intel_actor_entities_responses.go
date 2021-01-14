// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// QueryIntelActorEntitiesReader is a Reader for the QueryIntelActorEntities structure.
type QueryIntelActorEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryIntelActorEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryIntelActorEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryIntelActorEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryIntelActorEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryIntelActorEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryIntelActorEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryIntelActorEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryIntelActorEntitiesOK creates a QueryIntelActorEntitiesOK with default headers values
func NewQueryIntelActorEntitiesOK() *QueryIntelActorEntitiesOK {
	return &QueryIntelActorEntitiesOK{}
}

/*QueryIntelActorEntitiesOK handles this case with default header values.

OK
*/
type QueryIntelActorEntitiesOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainActorsResponse
}

func (o *QueryIntelActorEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] queryIntelActorEntitiesOK  %+v", 200, o.Payload)
}

func (o *QueryIntelActorEntitiesOK) GetPayload() *models.DomainActorsResponse {
	return o.Payload
}

func (o *QueryIntelActorEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainActorsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelActorEntitiesBadRequest creates a QueryIntelActorEntitiesBadRequest with default headers values
func NewQueryIntelActorEntitiesBadRequest() *QueryIntelActorEntitiesBadRequest {
	return &QueryIntelActorEntitiesBadRequest{}
}

/*QueryIntelActorEntitiesBadRequest handles this case with default header values.

Bad Request
*/
type QueryIntelActorEntitiesBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelActorEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] queryIntelActorEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryIntelActorEntitiesBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelActorEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelActorEntitiesForbidden creates a QueryIntelActorEntitiesForbidden with default headers values
func NewQueryIntelActorEntitiesForbidden() *QueryIntelActorEntitiesForbidden {
	return &QueryIntelActorEntitiesForbidden{}
}

/*QueryIntelActorEntitiesForbidden handles this case with default header values.

Forbidden
*/
type QueryIntelActorEntitiesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryIntelActorEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] queryIntelActorEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *QueryIntelActorEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelActorEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelActorEntitiesTooManyRequests creates a QueryIntelActorEntitiesTooManyRequests with default headers values
func NewQueryIntelActorEntitiesTooManyRequests() *QueryIntelActorEntitiesTooManyRequests {
	return &QueryIntelActorEntitiesTooManyRequests{}
}

/*QueryIntelActorEntitiesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryIntelActorEntitiesTooManyRequests struct {
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

func (o *QueryIntelActorEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] queryIntelActorEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryIntelActorEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelActorEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelActorEntitiesInternalServerError creates a QueryIntelActorEntitiesInternalServerError with default headers values
func NewQueryIntelActorEntitiesInternalServerError() *QueryIntelActorEntitiesInternalServerError {
	return &QueryIntelActorEntitiesInternalServerError{}
}

/*QueryIntelActorEntitiesInternalServerError handles this case with default header values.

Internal Server Error
*/
type QueryIntelActorEntitiesInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelActorEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] queryIntelActorEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryIntelActorEntitiesInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelActorEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelActorEntitiesDefault creates a QueryIntelActorEntitiesDefault with default headers values
func NewQueryIntelActorEntitiesDefault(code int) *QueryIntelActorEntitiesDefault {
	return &QueryIntelActorEntitiesDefault{
		_statusCode: code,
	}
}

/*QueryIntelActorEntitiesDefault handles this case with default header values.

OK
*/
type QueryIntelActorEntitiesDefault struct {
	_statusCode int

	Payload *models.DomainActorsResponse
}

// Code gets the status code for the query intel actor entities default response
func (o *QueryIntelActorEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *QueryIntelActorEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /intel/combined/actors/v1][%d] QueryIntelActorEntities default  %+v", o._statusCode, o.Payload)
}

func (o *QueryIntelActorEntitiesDefault) GetPayload() *models.DomainActorsResponse {
	return o.Payload
}

func (o *QueryIntelActorEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainActorsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
