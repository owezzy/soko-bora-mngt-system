// Code generated by go-swagger; DO NOT EDIT.

package store

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

// NewCreateStoreParams creates a new CreateStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStoreParams() *CreateStoreParams {
	return &CreateStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStoreParamsWithTimeout creates a new CreateStoreParams object
// with the ability to set a timeout on a request.
func NewCreateStoreParamsWithTimeout(timeout time.Duration) *CreateStoreParams {
	return &CreateStoreParams{
		timeout: timeout,
	}
}

// NewCreateStoreParamsWithContext creates a new CreateStoreParams object
// with the ability to set a context for a request.
func NewCreateStoreParamsWithContext(ctx context.Context) *CreateStoreParams {
	return &CreateStoreParams{
		Context: ctx,
	}
}

// NewCreateStoreParamsWithHTTPClient creates a new CreateStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStoreParamsWithHTTPClient(client *http.Client) *CreateStoreParams {
	return &CreateStoreParams{
		HTTPClient: client,
	}
}

/*
CreateStoreParams contains all the parameters to send to the API endpoint

	for the create store operation.

	Typically these are written to a http.Request.
*/
type CreateStoreParams struct {

	// Body.
	Body *models.StorespbCreateStoreRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoreParams) WithDefaults() *CreateStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create store params
func (o *CreateStoreParams) WithTimeout(timeout time.Duration) *CreateStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create store params
func (o *CreateStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create store params
func (o *CreateStoreParams) WithContext(ctx context.Context) *CreateStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create store params
func (o *CreateStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create store params
func (o *CreateStoreParams) WithHTTPClient(client *http.Client) *CreateStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create store params
func (o *CreateStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create store params
func (o *CreateStoreParams) WithBody(body *models.StorespbCreateStoreRequest) *CreateStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create store params
func (o *CreateStoreParams) SetBody(body *models.StorespbCreateStoreRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
