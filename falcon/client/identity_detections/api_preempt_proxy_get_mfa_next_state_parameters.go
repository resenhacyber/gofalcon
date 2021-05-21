// Code generated by go-swagger; DO NOT EDIT.

package identity_detections

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

// NewAPIPreemptProxyGetMfaNextStateParams creates a new APIPreemptProxyGetMfaNextStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAPIPreemptProxyGetMfaNextStateParams() *APIPreemptProxyGetMfaNextStateParams {
	return &APIPreemptProxyGetMfaNextStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAPIPreemptProxyGetMfaNextStateParamsWithTimeout creates a new APIPreemptProxyGetMfaNextStateParams object
// with the ability to set a timeout on a request.
func NewAPIPreemptProxyGetMfaNextStateParamsWithTimeout(timeout time.Duration) *APIPreemptProxyGetMfaNextStateParams {
	return &APIPreemptProxyGetMfaNextStateParams{
		timeout: timeout,
	}
}

// NewAPIPreemptProxyGetMfaNextStateParamsWithContext creates a new APIPreemptProxyGetMfaNextStateParams object
// with the ability to set a context for a request.
func NewAPIPreemptProxyGetMfaNextStateParamsWithContext(ctx context.Context) *APIPreemptProxyGetMfaNextStateParams {
	return &APIPreemptProxyGetMfaNextStateParams{
		Context: ctx,
	}
}

// NewAPIPreemptProxyGetMfaNextStateParamsWithHTTPClient creates a new APIPreemptProxyGetMfaNextStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAPIPreemptProxyGetMfaNextStateParamsWithHTTPClient(client *http.Client) *APIPreemptProxyGetMfaNextStateParams {
	return &APIPreemptProxyGetMfaNextStateParams{
		HTTPClient: client,
	}
}

/* APIPreemptProxyGetMfaNextStateParams contains all the parameters to send to the API endpoint
   for the api preempt proxy get mfa next state operation.

   Typically these are written to a http.Request.
*/
type APIPreemptProxyGetMfaNextStateParams struct {

	/* Authorization.

	   Authorization Header
	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the api preempt proxy get mfa next state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIPreemptProxyGetMfaNextStateParams) WithDefaults() *APIPreemptProxyGetMfaNextStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the api preempt proxy get mfa next state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *APIPreemptProxyGetMfaNextStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) WithTimeout(timeout time.Duration) *APIPreemptProxyGetMfaNextStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) WithContext(ctx context.Context) *APIPreemptProxyGetMfaNextStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) WithHTTPClient(client *http.Client) *APIPreemptProxyGetMfaNextStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) WithAuthorization(authorization string) *APIPreemptProxyGetMfaNextStateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the api preempt proxy get mfa next state params
func (o *APIPreemptProxyGetMfaNextStateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *APIPreemptProxyGetMfaNextStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
