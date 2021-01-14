// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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
)

// NewGetCSPMAwsAccountScriptsAttachmentParams creates a new GetCSPMAwsAccountScriptsAttachmentParams object
// with the default values initialized.
func NewGetCSPMAwsAccountScriptsAttachmentParams() *GetCSPMAwsAccountScriptsAttachmentParams {

	return &GetCSPMAwsAccountScriptsAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMAwsAccountScriptsAttachmentParamsWithTimeout creates a new GetCSPMAwsAccountScriptsAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCSPMAwsAccountScriptsAttachmentParamsWithTimeout(timeout time.Duration) *GetCSPMAwsAccountScriptsAttachmentParams {

	return &GetCSPMAwsAccountScriptsAttachmentParams{

		timeout: timeout,
	}
}

// NewGetCSPMAwsAccountScriptsAttachmentParamsWithContext creates a new GetCSPMAwsAccountScriptsAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCSPMAwsAccountScriptsAttachmentParamsWithContext(ctx context.Context) *GetCSPMAwsAccountScriptsAttachmentParams {

	return &GetCSPMAwsAccountScriptsAttachmentParams{

		Context: ctx,
	}
}

// NewGetCSPMAwsAccountScriptsAttachmentParamsWithHTTPClient creates a new GetCSPMAwsAccountScriptsAttachmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCSPMAwsAccountScriptsAttachmentParamsWithHTTPClient(client *http.Client) *GetCSPMAwsAccountScriptsAttachmentParams {

	return &GetCSPMAwsAccountScriptsAttachmentParams{
		HTTPClient: client,
	}
}

/*GetCSPMAwsAccountScriptsAttachmentParams contains all the parameters to send to the API endpoint
for the get c s p m aws account scripts attachment operation typically these are written to a http.Request
*/
type GetCSPMAwsAccountScriptsAttachmentParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) WithTimeout(timeout time.Duration) *GetCSPMAwsAccountScriptsAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) WithContext(ctx context.Context) *GetCSPMAwsAccountScriptsAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) WithHTTPClient(client *http.Client) *GetCSPMAwsAccountScriptsAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m aws account scripts attachment params
func (o *GetCSPMAwsAccountScriptsAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMAwsAccountScriptsAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
