// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// NewGetIntelRuleEntitiesParams creates a new GetIntelRuleEntitiesParams object
// with the default values initialized.
func NewGetIntelRuleEntitiesParams() *GetIntelRuleEntitiesParams {
	var ()
	return &GetIntelRuleEntitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntelRuleEntitiesParamsWithTimeout creates a new GetIntelRuleEntitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntelRuleEntitiesParamsWithTimeout(timeout time.Duration) *GetIntelRuleEntitiesParams {
	var ()
	return &GetIntelRuleEntitiesParams{

		timeout: timeout,
	}
}

// NewGetIntelRuleEntitiesParamsWithContext creates a new GetIntelRuleEntitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntelRuleEntitiesParamsWithContext(ctx context.Context) *GetIntelRuleEntitiesParams {
	var ()
	return &GetIntelRuleEntitiesParams{

		Context: ctx,
	}
}

// NewGetIntelRuleEntitiesParamsWithHTTPClient creates a new GetIntelRuleEntitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntelRuleEntitiesParamsWithHTTPClient(client *http.Client) *GetIntelRuleEntitiesParams {
	var ()
	return &GetIntelRuleEntitiesParams{
		HTTPClient: client,
	}
}

/*GetIntelRuleEntitiesParams contains all the parameters to send to the API endpoint
for the get intel rule entities operation typically these are written to a http.Request
*/
type GetIntelRuleEntitiesParams struct {

	/*Ids
	  The ids of rules to return.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) WithTimeout(timeout time.Duration) *GetIntelRuleEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) WithContext(ctx context.Context) *GetIntelRuleEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) WithHTTPClient(client *http.Client) *GetIntelRuleEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) WithIds(ids []string) *GetIntelRuleEntitiesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get intel rule entities params
func (o *GetIntelRuleEntitiesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntelRuleEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
