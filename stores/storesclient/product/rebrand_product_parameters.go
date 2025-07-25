// Code generated by go-swagger; DO NOT EDIT.

package product

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

	"github.com/owezzy/soko-bora-mngt-system/stores/storesclient/models"
)

// NewRebrandProductParams creates a new RebrandProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRebrandProductParams() *RebrandProductParams {
	return &RebrandProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRebrandProductParamsWithTimeout creates a new RebrandProductParams object
// with the ability to set a timeout on a request.
func NewRebrandProductParamsWithTimeout(timeout time.Duration) *RebrandProductParams {
	return &RebrandProductParams{
		timeout: timeout,
	}
}

// NewRebrandProductParamsWithContext creates a new RebrandProductParams object
// with the ability to set a context for a request.
func NewRebrandProductParamsWithContext(ctx context.Context) *RebrandProductParams {
	return &RebrandProductParams{
		Context: ctx,
	}
}

// NewRebrandProductParamsWithHTTPClient creates a new RebrandProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewRebrandProductParamsWithHTTPClient(client *http.Client) *RebrandProductParams {
	return &RebrandProductParams{
		HTTPClient: client,
	}
}

/*
RebrandProductParams contains all the parameters to send to the API endpoint

	for the rebrand product operation.

	Typically these are written to a http.Request.
*/
type RebrandProductParams struct {

	// Body.
	Body *models.RebrandProductParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rebrand product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebrandProductParams) WithDefaults() *RebrandProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rebrand product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebrandProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rebrand product params
func (o *RebrandProductParams) WithTimeout(timeout time.Duration) *RebrandProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebrand product params
func (o *RebrandProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebrand product params
func (o *RebrandProductParams) WithContext(ctx context.Context) *RebrandProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebrand product params
func (o *RebrandProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebrand product params
func (o *RebrandProductParams) WithHTTPClient(client *http.Client) *RebrandProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebrand product params
func (o *RebrandProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rebrand product params
func (o *RebrandProductParams) WithBody(body *models.RebrandProductParamsBody) *RebrandProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rebrand product params
func (o *RebrandProductParams) SetBody(body *models.RebrandProductParamsBody) {
	o.Body = body
}

// WithID adds the id to the rebrand product params
func (o *RebrandProductParams) WithID(id string) *RebrandProductParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rebrand product params
func (o *RebrandProductParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RebrandProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
