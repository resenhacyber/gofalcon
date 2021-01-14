// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRListAllSessionsParams creates a new RTRListAllSessionsParams object
// with the default values initialized.
func NewRTRListAllSessionsParams() *RTRListAllSessionsParams {
	var ()
	return &RTRListAllSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRTRListAllSessionsParamsWithTimeout creates a new RTRListAllSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRTRListAllSessionsParamsWithTimeout(timeout time.Duration) *RTRListAllSessionsParams {
	var ()
	return &RTRListAllSessionsParams{

		timeout: timeout,
	}
}

// NewRTRListAllSessionsParamsWithContext creates a new RTRListAllSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRTRListAllSessionsParamsWithContext(ctx context.Context) *RTRListAllSessionsParams {
	var ()
	return &RTRListAllSessionsParams{

		Context: ctx,
	}
}

// NewRTRListAllSessionsParamsWithHTTPClient creates a new RTRListAllSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRTRListAllSessionsParamsWithHTTPClient(client *http.Client) *RTRListAllSessionsParams {
	var ()
	return &RTRListAllSessionsParams{
		HTTPClient: client,
	}
}

/*RTRListAllSessionsParams contains all the parameters to send to the API endpoint
for the r t r list all sessions operation typically these are written to a http.Request
*/
type RTRListAllSessionsParams struct {

	/*Filter
	  Optional filter criteria in the form of an FQL query. For more information about FQL queries, see our [FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide). “user_id” can accept a special value ‘@me’ which will restrict results to records with current user’s ID.

	*/
	Filter *string
	/*Limit
	  Number of ids to return.

	*/
	Limit *int64
	/*Offset
	  Starting index of overall result set from which to return ids.

	*/
	Offset *string
	/*Sort
	  Sort by spec. Ex: 'date_created|asc'.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithTimeout(timeout time.Duration) *RTRListAllSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithContext(ctx context.Context) *RTRListAllSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithHTTPClient(client *http.Client) *RTRListAllSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithFilter(filter *string) *RTRListAllSessionsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithLimit(limit *int64) *RTRListAllSessionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithOffset(offset *string) *RTRListAllSessionsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithSort adds the sort to the r t r list all sessions params
func (o *RTRListAllSessionsParams) WithSort(sort *string) *RTRListAllSessionsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the r t r list all sessions params
func (o *RTRListAllSessionsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *RTRListAllSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
