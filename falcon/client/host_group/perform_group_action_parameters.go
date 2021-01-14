// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewPerformGroupActionParams creates a new PerformGroupActionParams object
// with the default values initialized.
func NewPerformGroupActionParams() *PerformGroupActionParams {
	var ()
	return &PerformGroupActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPerformGroupActionParamsWithTimeout creates a new PerformGroupActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPerformGroupActionParamsWithTimeout(timeout time.Duration) *PerformGroupActionParams {
	var ()
	return &PerformGroupActionParams{

		timeout: timeout,
	}
}

// NewPerformGroupActionParamsWithContext creates a new PerformGroupActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPerformGroupActionParamsWithContext(ctx context.Context) *PerformGroupActionParams {
	var ()
	return &PerformGroupActionParams{

		Context: ctx,
	}
}

// NewPerformGroupActionParamsWithHTTPClient creates a new PerformGroupActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPerformGroupActionParamsWithHTTPClient(client *http.Client) *PerformGroupActionParams {
	var ()
	return &PerformGroupActionParams{
		HTTPClient: client,
	}
}

/*PerformGroupActionParams contains all the parameters to send to the API endpoint
for the perform group action operation typically these are written to a http.Request
*/
type PerformGroupActionParams struct {

	/*ActionName
	  The action to perform

	*/
	ActionName string
	/*Body*/
	Body *models.MsaEntityActionRequestV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the perform group action params
func (o *PerformGroupActionParams) WithTimeout(timeout time.Duration) *PerformGroupActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform group action params
func (o *PerformGroupActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform group action params
func (o *PerformGroupActionParams) WithContext(ctx context.Context) *PerformGroupActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform group action params
func (o *PerformGroupActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform group action params
func (o *PerformGroupActionParams) WithHTTPClient(client *http.Client) *PerformGroupActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform group action params
func (o *PerformGroupActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the perform group action params
func (o *PerformGroupActionParams) WithActionName(actionName string) *PerformGroupActionParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the perform group action params
func (o *PerformGroupActionParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the perform group action params
func (o *PerformGroupActionParams) WithBody(body *models.MsaEntityActionRequestV2) *PerformGroupActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform group action params
func (o *PerformGroupActionParams) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformGroupActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action_name
	qrActionName := o.ActionName
	qActionName := qrActionName
	if qActionName != "" {
		if err := r.SetQueryParam("action_name", qActionName); err != nil {
			return err
		}
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
