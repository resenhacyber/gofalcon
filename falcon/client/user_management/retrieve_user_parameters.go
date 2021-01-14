// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// NewRetrieveUserParams creates a new RetrieveUserParams object
// with the default values initialized.
func NewRetrieveUserParams() *RetrieveUserParams {
	var ()
	return &RetrieveUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveUserParamsWithTimeout creates a new RetrieveUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveUserParamsWithTimeout(timeout time.Duration) *RetrieveUserParams {
	var ()
	return &RetrieveUserParams{

		timeout: timeout,
	}
}

// NewRetrieveUserParamsWithContext creates a new RetrieveUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveUserParamsWithContext(ctx context.Context) *RetrieveUserParams {
	var ()
	return &RetrieveUserParams{

		Context: ctx,
	}
}

// NewRetrieveUserParamsWithHTTPClient creates a new RetrieveUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveUserParamsWithHTTPClient(client *http.Client) *RetrieveUserParams {
	var ()
	return &RetrieveUserParams{
		HTTPClient: client,
	}
}

/*RetrieveUserParams contains all the parameters to send to the API endpoint
for the retrieve user operation typically these are written to a http.Request
*/
type RetrieveUserParams struct {

	/*Ids
	  ID of a user. Find a user's ID from `/users/entities/user/v1`.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve user params
func (o *RetrieveUserParams) WithTimeout(timeout time.Duration) *RetrieveUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve user params
func (o *RetrieveUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve user params
func (o *RetrieveUserParams) WithContext(ctx context.Context) *RetrieveUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve user params
func (o *RetrieveUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve user params
func (o *RetrieveUserParams) WithHTTPClient(client *http.Client) *RetrieveUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve user params
func (o *RetrieveUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the retrieve user params
func (o *RetrieveUserParams) WithIds(ids []string) *RetrieveUserParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the retrieve user params
func (o *RetrieveUserParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
