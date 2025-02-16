// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// NewDiscoverCloudAzureDownloadCertificateParams creates a new DiscoverCloudAzureDownloadCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscoverCloudAzureDownloadCertificateParams() *DiscoverCloudAzureDownloadCertificateParams {
	return &DiscoverCloudAzureDownloadCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscoverCloudAzureDownloadCertificateParamsWithTimeout creates a new DiscoverCloudAzureDownloadCertificateParams object
// with the ability to set a timeout on a request.
func NewDiscoverCloudAzureDownloadCertificateParamsWithTimeout(timeout time.Duration) *DiscoverCloudAzureDownloadCertificateParams {
	return &DiscoverCloudAzureDownloadCertificateParams{
		timeout: timeout,
	}
}

// NewDiscoverCloudAzureDownloadCertificateParamsWithContext creates a new DiscoverCloudAzureDownloadCertificateParams object
// with the ability to set a context for a request.
func NewDiscoverCloudAzureDownloadCertificateParamsWithContext(ctx context.Context) *DiscoverCloudAzureDownloadCertificateParams {
	return &DiscoverCloudAzureDownloadCertificateParams{
		Context: ctx,
	}
}

// NewDiscoverCloudAzureDownloadCertificateParamsWithHTTPClient creates a new DiscoverCloudAzureDownloadCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscoverCloudAzureDownloadCertificateParamsWithHTTPClient(client *http.Client) *DiscoverCloudAzureDownloadCertificateParams {
	return &DiscoverCloudAzureDownloadCertificateParams{
		HTTPClient: client,
	}
}

/*
DiscoverCloudAzureDownloadCertificateParams contains all the parameters to send to the API endpoint

	for the discover cloud azure download certificate operation.

	Typically these are written to a http.Request.
*/
type DiscoverCloudAzureDownloadCertificateParams struct {

	// Refresh.
	//
	// Default: "false"
	Refresh *string

	/* TenantID.

	   Azure Tenant ID
	*/
	TenantID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the discover cloud azure download certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscoverCloudAzureDownloadCertificateParams) WithDefaults() *DiscoverCloudAzureDownloadCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discover cloud azure download certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscoverCloudAzureDownloadCertificateParams) SetDefaults() {
	var (
		refreshDefault = string("false")
	)

	val := DiscoverCloudAzureDownloadCertificateParams{
		Refresh: &refreshDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) WithTimeout(timeout time.Duration) *DiscoverCloudAzureDownloadCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) WithContext(ctx context.Context) *DiscoverCloudAzureDownloadCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) WithHTTPClient(client *http.Client) *DiscoverCloudAzureDownloadCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefresh adds the refresh to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) WithRefresh(refresh *string) *DiscoverCloudAzureDownloadCertificateParams {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) SetRefresh(refresh *string) {
	o.Refresh = refresh
}

// WithTenantID adds the tenantID to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) WithTenantID(tenantID []string) *DiscoverCloudAzureDownloadCertificateParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the discover cloud azure download certificate params
func (o *DiscoverCloudAzureDownloadCertificateParams) SetTenantID(tenantID []string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *DiscoverCloudAzureDownloadCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Refresh != nil {

		// query param refresh
		var qrRefresh string

		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := qrRefresh
		if qRefresh != "" {

			if err := r.SetQueryParam("refresh", qRefresh); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// binding items for tenant_id
		joinedTenantID := o.bindParamTenantID(reg)

		// query array param tenant_id
		if err := r.SetQueryParam("tenant_id", joinedTenantID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDiscoverCloudAzureDownloadCertificate binds the parameter tenant_id
func (o *DiscoverCloudAzureDownloadCertificateParams) bindParamTenantID(formats strfmt.Registry) []string {
	tenantIDIR := o.TenantID

	var tenantIDIC []string
	for _, tenantIDIIR := range tenantIDIR { // explode []string

		tenantIDIIV := tenantIDIIR // string as string
		tenantIDIC = append(tenantIDIC, tenantIDIIV)
	}

	// items.CollectionFormat: "multi"
	tenantIDIS := swag.JoinByFormat(tenantIDIC, "multi")

	return tenantIDIS
}
