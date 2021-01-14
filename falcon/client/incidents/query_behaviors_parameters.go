// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewQueryBehaviorsParams creates a new QueryBehaviorsParams object
// with the default values initialized.
func NewQueryBehaviorsParams() *QueryBehaviorsParams {
	var ()
	return &QueryBehaviorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryBehaviorsParamsWithTimeout creates a new QueryBehaviorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryBehaviorsParamsWithTimeout(timeout time.Duration) *QueryBehaviorsParams {
	var ()
	return &QueryBehaviorsParams{

		timeout: timeout,
	}
}

// NewQueryBehaviorsParamsWithContext creates a new QueryBehaviorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryBehaviorsParamsWithContext(ctx context.Context) *QueryBehaviorsParams {
	var ()
	return &QueryBehaviorsParams{

		Context: ctx,
	}
}

// NewQueryBehaviorsParamsWithHTTPClient creates a new QueryBehaviorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryBehaviorsParamsWithHTTPClient(client *http.Client) *QueryBehaviorsParams {
	var ()
	return &QueryBehaviorsParams{
		HTTPClient: client,
	}
}

/*QueryBehaviorsParams contains all the parameters to send to the API endpoint
for the query behaviors operation typically these are written to a http.Request
*/
type QueryBehaviorsParams struct {

	/*Filter
	  Optional filter and sort criteria in the form of an FQL query. For more information about FQL queries, see [our FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).

	*/
	Filter *string
	/*Limit
	  The maximum records to return. [1-500]

	*/
	Limit *int64
	/*Offset
	  Starting index of overall result set from which to return ids.

	*/
	Offset *string
	/*Sort
	  The property to sort on, followed by a dot (.), followed by the sort direction, either "asc" or "desc".

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query behaviors params
func (o *QueryBehaviorsParams) WithTimeout(timeout time.Duration) *QueryBehaviorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query behaviors params
func (o *QueryBehaviorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query behaviors params
func (o *QueryBehaviorsParams) WithContext(ctx context.Context) *QueryBehaviorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query behaviors params
func (o *QueryBehaviorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query behaviors params
func (o *QueryBehaviorsParams) WithHTTPClient(client *http.Client) *QueryBehaviorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query behaviors params
func (o *QueryBehaviorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query behaviors params
func (o *QueryBehaviorsParams) WithFilter(filter *string) *QueryBehaviorsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query behaviors params
func (o *QueryBehaviorsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query behaviors params
func (o *QueryBehaviorsParams) WithLimit(limit *int64) *QueryBehaviorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query behaviors params
func (o *QueryBehaviorsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query behaviors params
func (o *QueryBehaviorsParams) WithOffset(offset *string) *QueryBehaviorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query behaviors params
func (o *QueryBehaviorsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the query behaviors params
func (o *QueryBehaviorsParams) WithSort(sort *string) *QueryBehaviorsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query behaviors params
func (o *QueryBehaviorsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryBehaviorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
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
