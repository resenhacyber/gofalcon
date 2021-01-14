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
)

// NewGetIntelReportPDFParams creates a new GetIntelReportPDFParams object
// with the default values initialized.
func NewGetIntelReportPDFParams() *GetIntelReportPDFParams {
	var ()
	return &GetIntelReportPDFParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntelReportPDFParamsWithTimeout creates a new GetIntelReportPDFParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntelReportPDFParamsWithTimeout(timeout time.Duration) *GetIntelReportPDFParams {
	var ()
	return &GetIntelReportPDFParams{

		timeout: timeout,
	}
}

// NewGetIntelReportPDFParamsWithContext creates a new GetIntelReportPDFParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntelReportPDFParamsWithContext(ctx context.Context) *GetIntelReportPDFParams {
	var ()
	return &GetIntelReportPDFParams{

		Context: ctx,
	}
}

// NewGetIntelReportPDFParamsWithHTTPClient creates a new GetIntelReportPDFParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntelReportPDFParamsWithHTTPClient(client *http.Client) *GetIntelReportPDFParams {
	var ()
	return &GetIntelReportPDFParams{
		HTTPClient: client,
	}
}

/*GetIntelReportPDFParams contains all the parameters to send to the API endpoint
for the get intel report p d f operation typically these are written to a http.Request
*/
type GetIntelReportPDFParams struct {

	/*ID
	  The ID of the report you want to download as a PDF.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get intel report p d f params
func (o *GetIntelReportPDFParams) WithTimeout(timeout time.Duration) *GetIntelReportPDFParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intel report p d f params
func (o *GetIntelReportPDFParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intel report p d f params
func (o *GetIntelReportPDFParams) WithContext(ctx context.Context) *GetIntelReportPDFParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intel report p d f params
func (o *GetIntelReportPDFParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intel report p d f params
func (o *GetIntelReportPDFParams) WithHTTPClient(client *http.Client) *GetIntelReportPDFParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intel report p d f params
func (o *GetIntelReportPDFParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get intel report p d f params
func (o *GetIntelReportPDFParams) WithID(id string) *GetIntelReportPDFParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get intel report p d f params
func (o *GetIntelReportPDFParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntelReportPDFParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
