// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetCSPMCGPAccountReader is a Reader for the GetCSPMCGPAccount structure.
type GetCSPMCGPAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMCGPAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMCGPAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCSPMCGPAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMCGPAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMCGPAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMCGPAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMCGPAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMCGPAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMCGPAccountOK creates a GetCSPMCGPAccountOK with default headers values
func NewGetCSPMCGPAccountOK() *GetCSPMCGPAccountOK {
	return &GetCSPMCGPAccountOK{}
}

/*GetCSPMCGPAccountOK handles this case with default header values.

OK
*/
type GetCSPMCGPAccountOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

func (o *GetCSPMCGPAccountOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountOK  %+v", 200, o.Payload)
}

func (o *GetCSPMCGPAccountOK) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountMultiStatus creates a GetCSPMCGPAccountMultiStatus with default headers values
func NewGetCSPMCGPAccountMultiStatus() *GetCSPMCGPAccountMultiStatus {
	return &GetCSPMCGPAccountMultiStatus{}
}

/*GetCSPMCGPAccountMultiStatus handles this case with default header values.

Multi-Status
*/
type GetCSPMCGPAccountMultiStatus struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

func (o *GetCSPMCGPAccountMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCSPMCGPAccountMultiStatus) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountBadRequest creates a GetCSPMCGPAccountBadRequest with default headers values
func NewGetCSPMCGPAccountBadRequest() *GetCSPMCGPAccountBadRequest {
	return &GetCSPMCGPAccountBadRequest{}
}

/*GetCSPMCGPAccountBadRequest handles this case with default header values.

Bad Request
*/
type GetCSPMCGPAccountBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

func (o *GetCSPMCGPAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMCGPAccountBadRequest) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountForbidden creates a GetCSPMCGPAccountForbidden with default headers values
func NewGetCSPMCGPAccountForbidden() *GetCSPMCGPAccountForbidden {
	return &GetCSPMCGPAccountForbidden{}
}

/*GetCSPMCGPAccountForbidden handles this case with default header values.

Forbidden
*/
type GetCSPMCGPAccountForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMCGPAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMCGPAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMCGPAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMCGPAccountTooManyRequests creates a GetCSPMCGPAccountTooManyRequests with default headers values
func NewGetCSPMCGPAccountTooManyRequests() *GetCSPMCGPAccountTooManyRequests {
	return &GetCSPMCGPAccountTooManyRequests{}
}

/*GetCSPMCGPAccountTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCSPMCGPAccountTooManyRequests struct {
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

func (o *GetCSPMCGPAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMCGPAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMCGPAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMCGPAccountInternalServerError creates a GetCSPMCGPAccountInternalServerError with default headers values
func NewGetCSPMCGPAccountInternalServerError() *GetCSPMCGPAccountInternalServerError {
	return &GetCSPMCGPAccountInternalServerError{}
}

/*GetCSPMCGPAccountInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCSPMCGPAccountInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

func (o *GetCSPMCGPAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getCSPMCGPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMCGPAccountInternalServerError) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMCGPAccountDefault creates a GetCSPMCGPAccountDefault with default headers values
func NewGetCSPMCGPAccountDefault(code int) *GetCSPMCGPAccountDefault {
	return &GetCSPMCGPAccountDefault{
		_statusCode: code,
	}
}

/*GetCSPMCGPAccountDefault handles this case with default header values.

OK
*/
type GetCSPMCGPAccountDefault struct {
	_statusCode int

	Payload *models.RegistrationGCPAccountResponseV1
}

// Code gets the status code for the get c s p m c g p account default response
func (o *GetCSPMCGPAccountDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMCGPAccountDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] GetCSPMCGPAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMCGPAccountDefault) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetCSPMCGPAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
