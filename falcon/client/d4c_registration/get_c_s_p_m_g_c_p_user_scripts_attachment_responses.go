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

// GetCSPMGCPUserScriptsAttachmentReader is a Reader for the GetCSPMGCPUserScriptsAttachment structure.
type GetCSPMGCPUserScriptsAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMGCPUserScriptsAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMGCPUserScriptsAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMGCPUserScriptsAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMGCPUserScriptsAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMGCPUserScriptsAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMGCPUserScriptsAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMGCPUserScriptsAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMGCPUserScriptsAttachmentOK creates a GetCSPMGCPUserScriptsAttachmentOK with default headers values
func NewGetCSPMGCPUserScriptsAttachmentOK() *GetCSPMGCPUserScriptsAttachmentOK {
	return &GetCSPMGCPUserScriptsAttachmentOK{}
}

/*
GetCSPMGCPUserScriptsAttachmentOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMGCPUserScriptsAttachmentOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get c s p m g c p user scripts attachment o k response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment o k response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p user scripts attachment o k response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m g c p user scripts attachment o k response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p user scripts attachment o k response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentBadRequest creates a GetCSPMGCPUserScriptsAttachmentBadRequest with default headers values
func NewGetCSPMGCPUserScriptsAttachmentBadRequest() *GetCSPMGCPUserScriptsAttachmentBadRequest {
	return &GetCSPMGCPUserScriptsAttachmentBadRequest{}
}

/*
GetCSPMGCPUserScriptsAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMGCPUserScriptsAttachmentBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get c s p m g c p user scripts attachment bad request response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment bad request response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p user scripts attachment bad request response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p user scripts attachment bad request response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p user scripts attachment bad request response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentForbidden creates a GetCSPMGCPUserScriptsAttachmentForbidden with default headers values
func NewGetCSPMGCPUserScriptsAttachmentForbidden() *GetCSPMGCPUserScriptsAttachmentForbidden {
	return &GetCSPMGCPUserScriptsAttachmentForbidden{}
}

/*
GetCSPMGCPUserScriptsAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMGCPUserScriptsAttachmentForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this get c s p m g c p user scripts attachment forbidden response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment forbidden response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p user scripts attachment forbidden response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p user scripts attachment forbidden response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p user scripts attachment forbidden response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMGCPUserScriptsAttachmentTooManyRequests creates a GetCSPMGCPUserScriptsAttachmentTooManyRequests with default headers values
func NewGetCSPMGCPUserScriptsAttachmentTooManyRequests() *GetCSPMGCPUserScriptsAttachmentTooManyRequests {
	return &GetCSPMGCPUserScriptsAttachmentTooManyRequests{}
}

/*
GetCSPMGCPUserScriptsAttachmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMGCPUserScriptsAttachmentTooManyRequests struct {

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

// IsSuccess returns true when this get c s p m g c p user scripts attachment too many requests response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment too many requests response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p user scripts attachment too many requests response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c s p m g c p user scripts attachment too many requests response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c s p m g c p user scripts attachment too many requests response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMGCPUserScriptsAttachmentInternalServerError creates a GetCSPMGCPUserScriptsAttachmentInternalServerError with default headers values
func NewGetCSPMGCPUserScriptsAttachmentInternalServerError() *GetCSPMGCPUserScriptsAttachmentInternalServerError {
	return &GetCSPMGCPUserScriptsAttachmentInternalServerError{}
}

/*
GetCSPMGCPUserScriptsAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMGCPUserScriptsAttachmentInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// IsSuccess returns true when this get c s p m g c p user scripts attachment internal server error response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment internal server error response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c s p m g c p user scripts attachment internal server error response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c s p m g c p user scripts attachment internal server error response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get c s p m g c p user scripts attachment internal server error response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentDefault creates a GetCSPMGCPUserScriptsAttachmentDefault with default headers values
func NewGetCSPMGCPUserScriptsAttachmentDefault(code int) *GetCSPMGCPUserScriptsAttachmentDefault {
	return &GetCSPMGCPUserScriptsAttachmentDefault{
		_statusCode: code,
	}
}

/*
GetCSPMGCPUserScriptsAttachmentDefault describes a response with status code -1, with default header values.

OK
*/
type GetCSPMGCPUserScriptsAttachmentDefault struct {
	_statusCode int

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// Code gets the status code for the get c s p m g c p user scripts attachment default response
func (o *GetCSPMGCPUserScriptsAttachmentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get c s p m g c p user scripts attachment default response has a 2xx status code
func (o *GetCSPMGCPUserScriptsAttachmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get c s p m g c p user scripts attachment default response has a 3xx status code
func (o *GetCSPMGCPUserScriptsAttachmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get c s p m g c p user scripts attachment default response has a 4xx status code
func (o *GetCSPMGCPUserScriptsAttachmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get c s p m g c p user scripts attachment default response has a 5xx status code
func (o *GetCSPMGCPUserScriptsAttachmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get c s p m g c p user scripts attachment default response a status code equal to that given
func (o *GetCSPMGCPUserScriptsAttachmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] GetCSPMGCPUserScriptsAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] GetCSPMGCPUserScriptsAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
