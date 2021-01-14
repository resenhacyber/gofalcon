// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// NewQueryCombinedFirewallPoliciesParams creates a new QueryCombinedFirewallPoliciesParams object
// with the default values initialized.
func NewQueryCombinedFirewallPoliciesParams() *QueryCombinedFirewallPoliciesParams {
	var ()
	return &QueryCombinedFirewallPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryCombinedFirewallPoliciesParamsWithTimeout creates a new QueryCombinedFirewallPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryCombinedFirewallPoliciesParamsWithTimeout(timeout time.Duration) *QueryCombinedFirewallPoliciesParams {
	var ()
	return &QueryCombinedFirewallPoliciesParams{

		timeout: timeout,
	}
}

// NewQueryCombinedFirewallPoliciesParamsWithContext creates a new QueryCombinedFirewallPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryCombinedFirewallPoliciesParamsWithContext(ctx context.Context) *QueryCombinedFirewallPoliciesParams {
	var ()
	return &QueryCombinedFirewallPoliciesParams{

		Context: ctx,
	}
}

// NewQueryCombinedFirewallPoliciesParamsWithHTTPClient creates a new QueryCombinedFirewallPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryCombinedFirewallPoliciesParamsWithHTTPClient(client *http.Client) *QueryCombinedFirewallPoliciesParams {
	var ()
	return &QueryCombinedFirewallPoliciesParams{
		HTTPClient: client,
	}
}

/*QueryCombinedFirewallPoliciesParams contains all the parameters to send to the API endpoint
for the query combined firewall policies operation typically these are written to a http.Request
*/
type QueryCombinedFirewallPoliciesParams struct {

	/*Filter
	  The filter expression that should be used to limit the results

	*/
	Filter *string
	/*Limit
	  The maximum records to return. [1-5000]

	*/
	Limit *int64
	/*Offset
	  The offset to start retrieving records from

	*/
	Offset *int64
	/*Sort
	  The property to sort by

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithTimeout(timeout time.Duration) *QueryCombinedFirewallPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithContext(ctx context.Context) *QueryCombinedFirewallPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithHTTPClient(client *http.Client) *QueryCombinedFirewallPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithFilter(filter *string) *QueryCombinedFirewallPoliciesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithLimit(limit *int64) *QueryCombinedFirewallPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithOffset(offset *int64) *QueryCombinedFirewallPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) WithSort(sort *string) *QueryCombinedFirewallPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query combined firewall policies params
func (o *QueryCombinedFirewallPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryCombinedFirewallPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
