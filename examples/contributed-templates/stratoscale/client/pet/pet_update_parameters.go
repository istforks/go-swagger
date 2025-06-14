// Code generated by go-swagger; DO NOT EDIT.

package pet

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

	"github.com/istforks/go-swagger/examples/contributed-templates/stratoscale/models"
)

// NewPetUpdateParams creates a new PetUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPetUpdateParams() *PetUpdateParams {
	return &PetUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPetUpdateParamsWithTimeout creates a new PetUpdateParams object
// with the ability to set a timeout on a request.
func NewPetUpdateParamsWithTimeout(timeout time.Duration) *PetUpdateParams {
	return &PetUpdateParams{
		timeout: timeout,
	}
}

// NewPetUpdateParamsWithContext creates a new PetUpdateParams object
// with the ability to set a context for a request.
func NewPetUpdateParamsWithContext(ctx context.Context) *PetUpdateParams {
	return &PetUpdateParams{
		Context: ctx,
	}
}

// NewPetUpdateParamsWithHTTPClient creates a new PetUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPetUpdateParamsWithHTTPClient(client *http.Client) *PetUpdateParams {
	return &PetUpdateParams{
		HTTPClient: client,
	}
}

/*
PetUpdateParams contains all the parameters to send to the API endpoint

	for the pet update operation.

	Typically these are written to a http.Request.
*/
type PetUpdateParams struct {

	/* Body.

	   Pet object that needs to be added to the store
	*/
	Body *models.Pet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pet update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PetUpdateParams) WithDefaults() *PetUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pet update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PetUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pet update params
func (o *PetUpdateParams) WithTimeout(timeout time.Duration) *PetUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pet update params
func (o *PetUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pet update params
func (o *PetUpdateParams) WithContext(ctx context.Context) *PetUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pet update params
func (o *PetUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pet update params
func (o *PetUpdateParams) WithHTTPClient(client *http.Client) *PetUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pet update params
func (o *PetUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pet update params
func (o *PetUpdateParams) WithBody(body *models.Pet) *PetUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pet update params
func (o *PetUpdateParams) SetBody(body *models.Pet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PetUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
