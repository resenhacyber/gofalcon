// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// QueryPreventionPoliciesReader is a Reader for the QueryPreventionPolicies structure.
type QueryPreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryPreventionPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryPreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryPreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryPreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryPreventionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryPreventionPoliciesOK creates a QueryPreventionPoliciesOK with default headers values
func NewQueryPreventionPoliciesOK() *QueryPreventionPoliciesOK {
	return &QueryPreventionPoliciesOK{}
}

/*QueryPreventionPoliciesOK handles this case with default header values.

OK
*/
type QueryPreventionPoliciesOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryPreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryPreventionPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesBadRequest creates a QueryPreventionPoliciesBadRequest with default headers values
func NewQueryPreventionPoliciesBadRequest() *QueryPreventionPoliciesBadRequest {
	return &QueryPreventionPoliciesBadRequest{}
}

/*QueryPreventionPoliciesBadRequest handles this case with default header values.

Bad Request
*/
type QueryPreventionPoliciesBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryPreventionPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryPreventionPoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesForbidden creates a QueryPreventionPoliciesForbidden with default headers values
func NewQueryPreventionPoliciesForbidden() *QueryPreventionPoliciesForbidden {
	return &QueryPreventionPoliciesForbidden{}
}

/*QueryPreventionPoliciesForbidden handles this case with default header values.

Forbidden
*/
type QueryPreventionPoliciesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryPreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryPreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryPreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesTooManyRequests creates a QueryPreventionPoliciesTooManyRequests with default headers values
func NewQueryPreventionPoliciesTooManyRequests() *QueryPreventionPoliciesTooManyRequests {
	return &QueryPreventionPoliciesTooManyRequests{}
}

/*QueryPreventionPoliciesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryPreventionPoliciesTooManyRequests struct {
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

func (o *QueryPreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesInternalServerError creates a QueryPreventionPoliciesInternalServerError with default headers values
func NewQueryPreventionPoliciesInternalServerError() *QueryPreventionPoliciesInternalServerError {
	return &QueryPreventionPoliciesInternalServerError{}
}

/*QueryPreventionPoliciesInternalServerError handles this case with default header values.

Internal Server Error
*/
type QueryPreventionPoliciesInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryPreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryPreventionPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesDefault creates a QueryPreventionPoliciesDefault with default headers values
func NewQueryPreventionPoliciesDefault(code int) *QueryPreventionPoliciesDefault {
	return &QueryPreventionPoliciesDefault{
		_statusCode: code,
	}
}

/*QueryPreventionPoliciesDefault handles this case with default header values.

OK
*/
type QueryPreventionPoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query prevention policies default response
func (o *QueryPreventionPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *QueryPreventionPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryPreventionPoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
