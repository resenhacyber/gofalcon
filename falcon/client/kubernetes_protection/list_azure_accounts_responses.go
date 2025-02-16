// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// ListAzureAccountsReader is a Reader for the ListAzureAccounts structure.
type ListAzureAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewListAzureAccountsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAzureAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAzureAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAzureAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAzureAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAzureAccountsOK creates a ListAzureAccountsOK with default headers values
func NewListAzureAccountsOK() *ListAzureAccountsOK {
	return &ListAzureAccountsOK{}
}

/*
ListAzureAccountsOK describes a response with status code 200, with default header values.

OK
*/
type ListAzureAccountsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureSubscriptionsResp
}

// IsSuccess returns true when this list azure accounts o k response has a 2xx status code
func (o *ListAzureAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure accounts o k response has a 3xx status code
func (o *ListAzureAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts o k response has a 4xx status code
func (o *ListAzureAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure accounts o k response has a 5xx status code
func (o *ListAzureAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure accounts o k response a status code equal to that given
func (o *ListAzureAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAzureAccountsOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAzureAccountsOK) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAzureAccountsOK) GetPayload() *models.K8sregGetAzureSubscriptionsResp {
	return o.Payload
}

func (o *ListAzureAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureSubscriptionsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAccountsMultiStatus creates a ListAzureAccountsMultiStatus with default headers values
func NewListAzureAccountsMultiStatus() *ListAzureAccountsMultiStatus {
	return &ListAzureAccountsMultiStatus{}
}

/*
ListAzureAccountsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type ListAzureAccountsMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureSubscriptionsResp
}

// IsSuccess returns true when this list azure accounts multi status response has a 2xx status code
func (o *ListAzureAccountsMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure accounts multi status response has a 3xx status code
func (o *ListAzureAccountsMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts multi status response has a 4xx status code
func (o *ListAzureAccountsMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure accounts multi status response has a 5xx status code
func (o *ListAzureAccountsMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure accounts multi status response a status code equal to that given
func (o *ListAzureAccountsMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *ListAzureAccountsMultiStatus) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsMultiStatus  %+v", 207, o.Payload)
}

func (o *ListAzureAccountsMultiStatus) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsMultiStatus  %+v", 207, o.Payload)
}

func (o *ListAzureAccountsMultiStatus) GetPayload() *models.K8sregGetAzureSubscriptionsResp {
	return o.Payload
}

func (o *ListAzureAccountsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureSubscriptionsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAccountsBadRequest creates a ListAzureAccountsBadRequest with default headers values
func NewListAzureAccountsBadRequest() *ListAzureAccountsBadRequest {
	return &ListAzureAccountsBadRequest{}
}

/*
ListAzureAccountsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListAzureAccountsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureSubscriptionsResp
}

// IsSuccess returns true when this list azure accounts bad request response has a 2xx status code
func (o *ListAzureAccountsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list azure accounts bad request response has a 3xx status code
func (o *ListAzureAccountsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts bad request response has a 4xx status code
func (o *ListAzureAccountsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list azure accounts bad request response has a 5xx status code
func (o *ListAzureAccountsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure accounts bad request response a status code equal to that given
func (o *ListAzureAccountsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListAzureAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAzureAccountsBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAzureAccountsBadRequest) GetPayload() *models.K8sregGetAzureSubscriptionsResp {
	return o.Payload
}

func (o *ListAzureAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureSubscriptionsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAccountsForbidden creates a ListAzureAccountsForbidden with default headers values
func NewListAzureAccountsForbidden() *ListAzureAccountsForbidden {
	return &ListAzureAccountsForbidden{}
}

/*
ListAzureAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListAzureAccountsForbidden struct {

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

// IsSuccess returns true when this list azure accounts forbidden response has a 2xx status code
func (o *ListAzureAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list azure accounts forbidden response has a 3xx status code
func (o *ListAzureAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts forbidden response has a 4xx status code
func (o *ListAzureAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list azure accounts forbidden response has a 5xx status code
func (o *ListAzureAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure accounts forbidden response a status code equal to that given
func (o *ListAzureAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListAzureAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsForbidden  %+v", 403, o.Payload)
}

func (o *ListAzureAccountsForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsForbidden  %+v", 403, o.Payload)
}

func (o *ListAzureAccountsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListAzureAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAzureAccountsTooManyRequests creates a ListAzureAccountsTooManyRequests with default headers values
func NewListAzureAccountsTooManyRequests() *ListAzureAccountsTooManyRequests {
	return &ListAzureAccountsTooManyRequests{}
}

/*
ListAzureAccountsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ListAzureAccountsTooManyRequests struct {

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

// IsSuccess returns true when this list azure accounts too many requests response has a 2xx status code
func (o *ListAzureAccountsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list azure accounts too many requests response has a 3xx status code
func (o *ListAzureAccountsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts too many requests response has a 4xx status code
func (o *ListAzureAccountsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list azure accounts too many requests response has a 5xx status code
func (o *ListAzureAccountsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure accounts too many requests response a status code equal to that given
func (o *ListAzureAccountsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListAzureAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAzureAccountsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAzureAccountsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ListAzureAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAzureAccountsInternalServerError creates a ListAzureAccountsInternalServerError with default headers values
func NewListAzureAccountsInternalServerError() *ListAzureAccountsInternalServerError {
	return &ListAzureAccountsInternalServerError{}
}

/*
ListAzureAccountsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ListAzureAccountsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureSubscriptionsResp
}

// IsSuccess returns true when this list azure accounts internal server error response has a 2xx status code
func (o *ListAzureAccountsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list azure accounts internal server error response has a 3xx status code
func (o *ListAzureAccountsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure accounts internal server error response has a 4xx status code
func (o *ListAzureAccountsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure accounts internal server error response has a 5xx status code
func (o *ListAzureAccountsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list azure accounts internal server error response a status code equal to that given
func (o *ListAzureAccountsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListAzureAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAzureAccountsInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/accounts/azure/v1][%d] listAzureAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAzureAccountsInternalServerError) GetPayload() *models.K8sregGetAzureSubscriptionsResp {
	return o.Payload
}

func (o *ListAzureAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureSubscriptionsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
