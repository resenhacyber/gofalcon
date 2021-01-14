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

// NewQueryIntelActorIdsParams creates a new QueryIntelActorIdsParams object
// with the default values initialized.
func NewQueryIntelActorIdsParams() *QueryIntelActorIdsParams {
	var ()
	return &QueryIntelActorIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryIntelActorIdsParamsWithTimeout creates a new QueryIntelActorIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryIntelActorIdsParamsWithTimeout(timeout time.Duration) *QueryIntelActorIdsParams {
	var ()
	return &QueryIntelActorIdsParams{

		timeout: timeout,
	}
}

// NewQueryIntelActorIdsParamsWithContext creates a new QueryIntelActorIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryIntelActorIdsParamsWithContext(ctx context.Context) *QueryIntelActorIdsParams {
	var ()
	return &QueryIntelActorIdsParams{

		Context: ctx,
	}
}

// NewQueryIntelActorIdsParamsWithHTTPClient creates a new QueryIntelActorIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryIntelActorIdsParamsWithHTTPClient(client *http.Client) *QueryIntelActorIdsParams {
	var ()
	return &QueryIntelActorIdsParams{
		HTTPClient: client,
	}
}

/*QueryIntelActorIdsParams contains all the parameters to send to the API endpoint
for the query intel actor ids operation typically these are written to a http.Request
*/
type QueryIntelActorIdsParams struct {

	/*Filter
	  Filter your query by specifying FQL filter parameters. Filter parameters include:

	actors, actors.id, actors.name, actors.slug, actors.url, created_date, description, id, last_modified_date, motivations, motivations.id, motivations.slug, motivations.value, name, name.raw, short_description, slug, sub_type, sub_type.id, sub_type.name, sub_type.slug, tags, tags.id, tags.slug, tags.value, target_countries, target_countries.id, target_countries.slug, target_countries.value, target_industries, target_industries.id, target_industries.slug, target_industries.value, type, type.id, type.name, type.slug, url.

	*/
	Filter *string
	/*Limit
	  Set the number of actor IDs to return. The value must be between 1 and 5000.

	*/
	Limit *int64
	/*Offset
	  Set the starting row number to return actors IDs from. Defaults to 0.

	*/
	Offset *int64
	/*Q
	  Perform a generic substring search across all fields.

	*/
	Q *string
	/*Sort
	  Order fields in ascending or descending order.

	Ex: created_date|asc.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithTimeout(timeout time.Duration) *QueryIntelActorIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithContext(ctx context.Context) *QueryIntelActorIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithHTTPClient(client *http.Client) *QueryIntelActorIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithFilter(filter *string) *QueryIntelActorIdsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithLimit(limit *int64) *QueryIntelActorIdsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithOffset(offset *int64) *QueryIntelActorIdsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithQ(q *string) *QueryIntelActorIdsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query intel actor ids params
func (o *QueryIntelActorIdsParams) WithSort(sort *string) *QueryIntelActorIdsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query intel actor ids params
func (o *QueryIntelActorIdsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryIntelActorIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
