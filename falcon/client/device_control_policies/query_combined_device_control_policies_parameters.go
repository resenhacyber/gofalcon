// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// NewQueryCombinedDeviceControlPoliciesParams creates a new QueryCombinedDeviceControlPoliciesParams object
// with the default values initialized.
func NewQueryCombinedDeviceControlPoliciesParams() *QueryCombinedDeviceControlPoliciesParams {
	var ()
	return &QueryCombinedDeviceControlPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryCombinedDeviceControlPoliciesParamsWithTimeout creates a new QueryCombinedDeviceControlPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryCombinedDeviceControlPoliciesParamsWithTimeout(timeout time.Duration) *QueryCombinedDeviceControlPoliciesParams {
	var ()
	return &QueryCombinedDeviceControlPoliciesParams{

		timeout: timeout,
	}
}

// NewQueryCombinedDeviceControlPoliciesParamsWithContext creates a new QueryCombinedDeviceControlPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryCombinedDeviceControlPoliciesParamsWithContext(ctx context.Context) *QueryCombinedDeviceControlPoliciesParams {
	var ()
	return &QueryCombinedDeviceControlPoliciesParams{

		Context: ctx,
	}
}

// NewQueryCombinedDeviceControlPoliciesParamsWithHTTPClient creates a new QueryCombinedDeviceControlPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryCombinedDeviceControlPoliciesParamsWithHTTPClient(client *http.Client) *QueryCombinedDeviceControlPoliciesParams {
	var ()
	return &QueryCombinedDeviceControlPoliciesParams{
		HTTPClient: client,
	}
}

/*QueryCombinedDeviceControlPoliciesParams contains all the parameters to send to the API endpoint
for the query combined device control policies operation typically these are written to a http.Request
*/
type QueryCombinedDeviceControlPoliciesParams struct {

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

// WithTimeout adds the timeout to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithTimeout(timeout time.Duration) *QueryCombinedDeviceControlPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithContext(ctx context.Context) *QueryCombinedDeviceControlPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithHTTPClient(client *http.Client) *QueryCombinedDeviceControlPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithFilter(filter *string) *QueryCombinedDeviceControlPoliciesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithLimit(limit *int64) *QueryCombinedDeviceControlPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithOffset(offset *int64) *QueryCombinedDeviceControlPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) WithSort(sort *string) *QueryCombinedDeviceControlPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query combined device control policies params
func (o *QueryCombinedDeviceControlPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryCombinedDeviceControlPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
