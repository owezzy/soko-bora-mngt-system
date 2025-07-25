// Code generated by go-swagger; DO NOT EDIT.

package customer

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

	"github.com/owezzy/soko-bora-mngt-system/customers/customersclient/models"
)

// NewDisableCustomerParams creates a new DisableCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableCustomerParams() *DisableCustomerParams {
	return &DisableCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableCustomerParamsWithTimeout creates a new DisableCustomerParams object
// with the ability to set a timeout on a request.
func NewDisableCustomerParamsWithTimeout(timeout time.Duration) *DisableCustomerParams {
	return &DisableCustomerParams{
		timeout: timeout,
	}
}

// NewDisableCustomerParamsWithContext creates a new DisableCustomerParams object
// with the ability to set a context for a request.
func NewDisableCustomerParamsWithContext(ctx context.Context) *DisableCustomerParams {
	return &DisableCustomerParams{
		Context: ctx,
	}
}

// NewDisableCustomerParamsWithHTTPClient creates a new DisableCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableCustomerParamsWithHTTPClient(client *http.Client) *DisableCustomerParams {
	return &DisableCustomerParams{
		HTTPClient: client,
	}
}

/*
DisableCustomerParams contains all the parameters to send to the API endpoint

	for the disable customer operation.

	Typically these are written to a http.Request.
*/
type DisableCustomerParams struct {

	// Body.
	Body models.CustomersServiceDisableCustomerBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableCustomerParams) WithDefaults() *DisableCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable customer params
func (o *DisableCustomerParams) WithTimeout(timeout time.Duration) *DisableCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable customer params
func (o *DisableCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable customer params
func (o *DisableCustomerParams) WithContext(ctx context.Context) *DisableCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable customer params
func (o *DisableCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable customer params
func (o *DisableCustomerParams) WithHTTPClient(client *http.Client) *DisableCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable customer params
func (o *DisableCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the disable customer params
func (o *DisableCustomerParams) WithBody(body models.CustomersServiceDisableCustomerBody) *DisableCustomerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the disable customer params
func (o *DisableCustomerParams) SetBody(body models.CustomersServiceDisableCustomerBody) {
	o.Body = body
}

// WithID adds the id to the disable customer params
func (o *DisableCustomerParams) WithID(id string) *DisableCustomerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the disable customer params
func (o *DisableCustomerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DisableCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
