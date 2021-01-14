// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// NewGetScansParams creates a new GetScansParams object
// with the default values initialized.
func NewGetScansParams() *GetScansParams {
	var ()
	return &GetScansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScansParamsWithTimeout creates a new GetScansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScansParamsWithTimeout(timeout time.Duration) *GetScansParams {
	var ()
	return &GetScansParams{

		timeout: timeout,
	}
}

// NewGetScansParamsWithContext creates a new GetScansParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScansParamsWithContext(ctx context.Context) *GetScansParams {
	var ()
	return &GetScansParams{

		Context: ctx,
	}
}

// NewGetScansParamsWithHTTPClient creates a new GetScansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScansParamsWithHTTPClient(client *http.Client) *GetScansParams {
	var ()
	return &GetScansParams{
		HTTPClient: client,
	}
}

/*GetScansParams contains all the parameters to send to the API endpoint
for the get scans operation typically these are written to a http.Request
*/
type GetScansParams struct {

	/*Ids
	  ID of a submitted scan

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scans params
func (o *GetScansParams) WithTimeout(timeout time.Duration) *GetScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scans params
func (o *GetScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scans params
func (o *GetScansParams) WithContext(ctx context.Context) *GetScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scans params
func (o *GetScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scans params
func (o *GetScansParams) WithHTTPClient(client *http.Client) *GetScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scans params
func (o *GetScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get scans params
func (o *GetScansParams) WithIds(ids []string) *GetScansParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get scans params
func (o *GetScansParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
