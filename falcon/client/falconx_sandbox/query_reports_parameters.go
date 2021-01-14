// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// NewQueryReportsParams creates a new QueryReportsParams object
// with the default values initialized.
func NewQueryReportsParams() *QueryReportsParams {
	var ()
	return &QueryReportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryReportsParamsWithTimeout creates a new QueryReportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryReportsParamsWithTimeout(timeout time.Duration) *QueryReportsParams {
	var ()
	return &QueryReportsParams{

		timeout: timeout,
	}
}

// NewQueryReportsParamsWithContext creates a new QueryReportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryReportsParamsWithContext(ctx context.Context) *QueryReportsParams {
	var ()
	return &QueryReportsParams{

		Context: ctx,
	}
}

// NewQueryReportsParamsWithHTTPClient creates a new QueryReportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryReportsParamsWithHTTPClient(client *http.Client) *QueryReportsParams {
	var ()
	return &QueryReportsParams{
		HTTPClient: client,
	}
}

/*QueryReportsParams contains all the parameters to send to the API endpoint
for the query reports operation typically these are written to a http.Request
*/
type QueryReportsParams struct {

	/*Filter
	  Optional filter and sort criteria in the form of an FQL query. For more information about FQL queries, see [our FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).

	*/
	Filter *string
	/*Limit
	  Maximum number of report IDs to return. Max: 5000.

	*/
	Limit *int64
	/*Offset
	  The offset to start retrieving reports from.

	*/
	Offset *string
	/*Sort
	  Sort order: `asc` or `desc`.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query reports params
func (o *QueryReportsParams) WithTimeout(timeout time.Duration) *QueryReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query reports params
func (o *QueryReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query reports params
func (o *QueryReportsParams) WithContext(ctx context.Context) *QueryReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query reports params
func (o *QueryReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query reports params
func (o *QueryReportsParams) WithHTTPClient(client *http.Client) *QueryReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query reports params
func (o *QueryReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query reports params
func (o *QueryReportsParams) WithFilter(filter *string) *QueryReportsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query reports params
func (o *QueryReportsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query reports params
func (o *QueryReportsParams) WithLimit(limit *int64) *QueryReportsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query reports params
func (o *QueryReportsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query reports params
func (o *QueryReportsParams) WithOffset(offset *string) *QueryReportsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query reports params
func (o *QueryReportsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the query reports params
func (o *QueryReportsParams) WithSort(sort *string) *QueryReportsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query reports params
func (o *QueryReportsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
