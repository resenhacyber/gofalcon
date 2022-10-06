// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// NewQueryRulesV1Params creates a new QueryRulesV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryRulesV1Params() *QueryRulesV1Params {
	return &QueryRulesV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryRulesV1ParamsWithTimeout creates a new QueryRulesV1Params object
// with the ability to set a timeout on a request.
func NewQueryRulesV1ParamsWithTimeout(timeout time.Duration) *QueryRulesV1Params {
	return &QueryRulesV1Params{
		timeout: timeout,
	}
}

// NewQueryRulesV1ParamsWithContext creates a new QueryRulesV1Params object
// with the ability to set a context for a request.
func NewQueryRulesV1ParamsWithContext(ctx context.Context) *QueryRulesV1Params {
	return &QueryRulesV1Params{
		Context: ctx,
	}
}

// NewQueryRulesV1ParamsWithHTTPClient creates a new QueryRulesV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewQueryRulesV1ParamsWithHTTPClient(client *http.Client) *QueryRulesV1Params {
	return &QueryRulesV1Params{
		HTTPClient: client,
	}
}

/*
QueryRulesV1Params contains all the parameters to send to the API endpoint

	for the query rules v1 operation.

	Typically these are written to a http.Request.
*/
type QueryRulesV1Params struct {

	/* Filter.

	   FQL query to filter rules by. Possible filter properties are: `permissions`, `status`, `breach_monitoring_enabled`, `substring_matching_enabled`, `cid`, `user_uuid`, `priority`, `filter`, `created_timestamp`, `last_updated_timestamp`, `id`, `topic`.
	*/
	Filter *string

	/* Limit.

	   Number of IDs to return. Offset + limit should NOT be above 10K.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return IDs.
	*/
	Offset *int64

	/* Q.

	   Free text search across all indexed fields.
	*/
	Q *string

	/* Sort.

	   Possible order by fields: created_timestamp, last_updated_timestamp. Ex: `last_updated_timestamp|desc`.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query rules v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRulesV1Params) WithDefaults() *QueryRulesV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query rules v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRulesV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query rules v1 params
func (o *QueryRulesV1Params) WithTimeout(timeout time.Duration) *QueryRulesV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query rules v1 params
func (o *QueryRulesV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query rules v1 params
func (o *QueryRulesV1Params) WithContext(ctx context.Context) *QueryRulesV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query rules v1 params
func (o *QueryRulesV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query rules v1 params
func (o *QueryRulesV1Params) WithHTTPClient(client *http.Client) *QueryRulesV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query rules v1 params
func (o *QueryRulesV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query rules v1 params
func (o *QueryRulesV1Params) WithFilter(filter *string) *QueryRulesV1Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query rules v1 params
func (o *QueryRulesV1Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query rules v1 params
func (o *QueryRulesV1Params) WithLimit(limit *int64) *QueryRulesV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query rules v1 params
func (o *QueryRulesV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query rules v1 params
func (o *QueryRulesV1Params) WithOffset(offset *int64) *QueryRulesV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query rules v1 params
func (o *QueryRulesV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query rules v1 params
func (o *QueryRulesV1Params) WithQ(q *string) *QueryRulesV1Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query rules v1 params
func (o *QueryRulesV1Params) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query rules v1 params
func (o *QueryRulesV1Params) WithSort(sort *string) *QueryRulesV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query rules v1 params
func (o *QueryRulesV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryRulesV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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
