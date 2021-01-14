// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewBatchGetCmdParams creates a new BatchGetCmdParams object
// with the default values initialized.
func NewBatchGetCmdParams() *BatchGetCmdParams {
	var (
		timeoutDefault         = int64(30)
		timeoutDurationDefault = string("30s")
	)
	return &BatchGetCmdParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewBatchGetCmdParamsWithTimeout creates a new BatchGetCmdParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchGetCmdParamsWithTimeout(timeout time.Duration) *BatchGetCmdParams {
	var (
		timeoutDefault         = int64(30)
		timeoutDurationDefault = string("30s")
	)
	return &BatchGetCmdParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,

		requestTimeout: timeout,
	}
}

// NewBatchGetCmdParamsWithContext creates a new BatchGetCmdParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchGetCmdParamsWithContext(ctx context.Context) *BatchGetCmdParams {
	var (
		timeoutDefault         = int64(30)
		timeoutDurationDefault = string("30s")
	)
	return &BatchGetCmdParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,

		Context: ctx,
	}
}

// NewBatchGetCmdParamsWithHTTPClient creates a new BatchGetCmdParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchGetCmdParamsWithHTTPClient(client *http.Client) *BatchGetCmdParams {
	var (
		timeoutDefault         = int64(30)
		timeoutDurationDefault = string("30s")
	)
	return &BatchGetCmdParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,
		HTTPClient:      client,
	}
}

/*BatchGetCmdParams contains all the parameters to send to the API endpoint
for the batch get cmd operation typically these are written to a http.Request
*/
type BatchGetCmdParams struct {

	/*Body
	  **`batch_id`** Batch ID to execute the command on.  Received from `/real-time-response/combined/init-sessions/v1`.
	**`file_path`** Full path to the file that is to be retrieved from each host in the batch.
	**`optional_hosts`** List of a subset of hosts we want to run the command on.  If this list is supplied, only these hosts will receive the command.

	*/
	Body *models.DomainBatchGetCommandRequest
	/*Timeout
	  Timeout for how long to wait for the request in seconds, default timeout is 30 seconds. Maximum is 10 minutes.

	*/
	Timeout *int64
	/*TimeoutDuration
	  Timeout duration for for how long to wait for the request in duration syntax. Example, `10s`. Valid units: `ns, us, ms, s, m, h`. Maximum is 10 minutes.

	*/
	TimeoutDuration *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the batch get cmd params
func (o *BatchGetCmdParams) WithRequestTimeout(timeout time.Duration) *BatchGetCmdParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the batch get cmd params
func (o *BatchGetCmdParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the batch get cmd params
func (o *BatchGetCmdParams) WithContext(ctx context.Context) *BatchGetCmdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch get cmd params
func (o *BatchGetCmdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch get cmd params
func (o *BatchGetCmdParams) WithHTTPClient(client *http.Client) *BatchGetCmdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch get cmd params
func (o *BatchGetCmdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch get cmd params
func (o *BatchGetCmdParams) WithBody(body *models.DomainBatchGetCommandRequest) *BatchGetCmdParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch get cmd params
func (o *BatchGetCmdParams) SetBody(body *models.DomainBatchGetCommandRequest) {
	o.Body = body
}

// WithTimeout adds the timeout to the batch get cmd params
func (o *BatchGetCmdParams) WithTimeout(timeout *int64) *BatchGetCmdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch get cmd params
func (o *BatchGetCmdParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithTimeoutDuration adds the timeoutDuration to the batch get cmd params
func (o *BatchGetCmdParams) WithTimeoutDuration(timeoutDuration *string) *BatchGetCmdParams {
	o.SetTimeoutDuration(timeoutDuration)
	return o
}

// SetTimeoutDuration adds the timeoutDuration to the batch get cmd params
func (o *BatchGetCmdParams) SetTimeoutDuration(timeoutDuration *string) {
	o.TimeoutDuration = timeoutDuration
}

// WriteToRequest writes these params to a swagger request
func (o *BatchGetCmdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if o.TimeoutDuration != nil {

		// query param timeout_duration
		var qrTimeoutDuration string
		if o.TimeoutDuration != nil {
			qrTimeoutDuration = *o.TimeoutDuration
		}
		qTimeoutDuration := qrTimeoutDuration
		if qTimeoutDuration != "" {
			if err := r.SetQueryParam("timeout_duration", qTimeoutDuration); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
