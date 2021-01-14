// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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
)

// NewRTRCheckAdminCommandStatusParams creates a new RTRCheckAdminCommandStatusParams object
// with the default values initialized.
func NewRTRCheckAdminCommandStatusParams() *RTRCheckAdminCommandStatusParams {
	var (
		sequenceIDDefault = int64(0)
	)
	return &RTRCheckAdminCommandStatusParams{
		SequenceID: sequenceIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRTRCheckAdminCommandStatusParamsWithTimeout creates a new RTRCheckAdminCommandStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRTRCheckAdminCommandStatusParamsWithTimeout(timeout time.Duration) *RTRCheckAdminCommandStatusParams {
	var (
		sequenceIDDefault = int64(0)
	)
	return &RTRCheckAdminCommandStatusParams{
		SequenceID: sequenceIDDefault,

		timeout: timeout,
	}
}

// NewRTRCheckAdminCommandStatusParamsWithContext creates a new RTRCheckAdminCommandStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewRTRCheckAdminCommandStatusParamsWithContext(ctx context.Context) *RTRCheckAdminCommandStatusParams {
	var (
		sequenceIdDefault = int64(0)
	)
	return &RTRCheckAdminCommandStatusParams{
		SequenceID: sequenceIdDefault,

		Context: ctx,
	}
}

// NewRTRCheckAdminCommandStatusParamsWithHTTPClient creates a new RTRCheckAdminCommandStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRTRCheckAdminCommandStatusParamsWithHTTPClient(client *http.Client) *RTRCheckAdminCommandStatusParams {
	var (
		sequenceIdDefault = int64(0)
	)
	return &RTRCheckAdminCommandStatusParams{
		SequenceID: sequenceIdDefault,
		HTTPClient: client,
	}
}

/*RTRCheckAdminCommandStatusParams contains all the parameters to send to the API endpoint
for the r t r check admin command status operation typically these are written to a http.Request
*/
type RTRCheckAdminCommandStatusParams struct {

	/*CloudRequestID
	  Cloud Request ID of the executed command to query

	*/
	CloudRequestID string
	/*SequenceID
	  Sequence ID that we want to retrieve. Command responses are chunked across sequences

	*/
	SequenceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) WithTimeout(timeout time.Duration) *RTRCheckAdminCommandStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) WithContext(ctx context.Context) *RTRCheckAdminCommandStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) WithHTTPClient(client *http.Client) *RTRCheckAdminCommandStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudRequestID adds the cloudRequestID to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) WithCloudRequestID(cloudRequestID string) *RTRCheckAdminCommandStatusParams {
	o.SetCloudRequestID(cloudRequestID)
	return o
}

// SetCloudRequestID adds the cloudRequestId to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) SetCloudRequestID(cloudRequestID string) {
	o.CloudRequestID = cloudRequestID
}

// WithSequenceID adds the sequenceID to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) WithSequenceID(sequenceID int64) *RTRCheckAdminCommandStatusParams {
	o.SetSequenceID(sequenceID)
	return o
}

// SetSequenceID adds the sequenceId to the r t r check admin command status params
func (o *RTRCheckAdminCommandStatusParams) SetSequenceID(sequenceID int64) {
	o.SequenceID = sequenceID
}

// WriteToRequest writes these params to a swagger request
func (o *RTRCheckAdminCommandStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloud_request_id
	qrCloudRequestID := o.CloudRequestID
	qCloudRequestID := qrCloudRequestID
	if qCloudRequestID != "" {
		if err := r.SetQueryParam("cloud_request_id", qCloudRequestID); err != nil {
			return err
		}
	}

	// query param sequence_id
	qrSequenceID := o.SequenceID
	qSequenceID := swag.FormatInt64(qrSequenceID)
	if qSequenceID != "" {
		if err := r.SetQueryParam("sequence_id", qSequenceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
