// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// NewGetSampleV3Params creates a new GetSampleV3Params object
// with the default values initialized.
func NewGetSampleV3Params() *GetSampleV3Params {
	var (
		passwordProtectedDefault = string(false)
	)
	return &GetSampleV3Params{
		PasswordProtected: &passwordProtectedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSampleV3ParamsWithTimeout creates a new GetSampleV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSampleV3ParamsWithTimeout(timeout time.Duration) *GetSampleV3Params {
	var (
		passwordProtectedDefault = string(false)
	)
	return &GetSampleV3Params{
		PasswordProtected: &passwordProtectedDefault,

		timeout: timeout,
	}
}

// NewGetSampleV3ParamsWithContext creates a new GetSampleV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetSampleV3ParamsWithContext(ctx context.Context) *GetSampleV3Params {
	var (
		passwordProtectedDefault = string(false)
	)
	return &GetSampleV3Params{
		PasswordProtected: &passwordProtectedDefault,

		Context: ctx,
	}
}

// NewGetSampleV3ParamsWithHTTPClient creates a new GetSampleV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSampleV3ParamsWithHTTPClient(client *http.Client) *GetSampleV3Params {
	var (
		passwordProtectedDefault = string(false)
	)
	return &GetSampleV3Params{
		PasswordProtected: &passwordProtectedDefault,
		HTTPClient:        client,
	}
}

/*GetSampleV3Params contains all the parameters to send to the API endpoint
for the get sample v3 operation typically these are written to a http.Request
*/
type GetSampleV3Params struct {

	/*XCSUSERUUID
	  User UUID

	*/
	XCSUSERUUID *string
	/*Ids
	  The file SHA256.

	*/
	Ids string
	/*PasswordProtected
	  Flag whether the sample should be zipped and password protected with pass='infected'

	*/
	PasswordProtected *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sample v3 params
func (o *GetSampleV3Params) WithTimeout(timeout time.Duration) *GetSampleV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sample v3 params
func (o *GetSampleV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sample v3 params
func (o *GetSampleV3Params) WithContext(ctx context.Context) *GetSampleV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sample v3 params
func (o *GetSampleV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sample v3 params
func (o *GetSampleV3Params) WithHTTPClient(client *http.Client) *GetSampleV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sample v3 params
func (o *GetSampleV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the get sample v3 params
func (o *GetSampleV3Params) WithXCSUSERUUID(xCSUSERUUID *string) *GetSampleV3Params {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the get sample v3 params
func (o *GetSampleV3Params) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithIds adds the ids to the get sample v3 params
func (o *GetSampleV3Params) WithIds(ids string) *GetSampleV3Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get sample v3 params
func (o *GetSampleV3Params) SetIds(ids string) {
	o.Ids = ids
}

// WithPasswordProtected adds the passwordProtected to the get sample v3 params
func (o *GetSampleV3Params) WithPasswordProtected(passwordProtected *string) *GetSampleV3Params {
	o.SetPasswordProtected(passwordProtected)
	return o
}

// SetPasswordProtected adds the passwordProtected to the get sample v3 params
func (o *GetSampleV3Params) SetPasswordProtected(passwordProtected *string) {
	o.PasswordProtected = passwordProtected
}

// WriteToRequest writes these params to a swagger request
func (o *GetSampleV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
		}

	}

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if o.PasswordProtected != nil {

		// query param password_protected
		var qrPasswordProtected string
		if o.PasswordProtected != nil {
			qrPasswordProtected = *o.PasswordProtected
		}
		qPasswordProtected := qrPasswordProtected
		if qPasswordProtected != "" {
			if err := r.SetQueryParam("password_protected", qPasswordProtected); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
