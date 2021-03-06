// Code generated by go-swagger; DO NOT EDIT.

package jira

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

// NewAggregateJiraRemediationsParams creates a new AggregateJiraRemediationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateJiraRemediationsParams() *AggregateJiraRemediationsParams {
	return &AggregateJiraRemediationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateJiraRemediationsParamsWithTimeout creates a new AggregateJiraRemediationsParams object
// with the ability to set a timeout on a request.
func NewAggregateJiraRemediationsParamsWithTimeout(timeout time.Duration) *AggregateJiraRemediationsParams {
	return &AggregateJiraRemediationsParams{
		timeout: timeout,
	}
}

// NewAggregateJiraRemediationsParamsWithContext creates a new AggregateJiraRemediationsParams object
// with the ability to set a context for a request.
func NewAggregateJiraRemediationsParamsWithContext(ctx context.Context) *AggregateJiraRemediationsParams {
	return &AggregateJiraRemediationsParams{
		Context: ctx,
	}
}

// NewAggregateJiraRemediationsParamsWithHTTPClient creates a new AggregateJiraRemediationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateJiraRemediationsParamsWithHTTPClient(client *http.Client) *AggregateJiraRemediationsParams {
	return &AggregateJiraRemediationsParams{
		HTTPClient: client,
	}
}

/* AggregateJiraRemediationsParams contains all the parameters to send to the API endpoint
   for the aggregate jira remediations operation.

   Typically these are written to a http.Request.
*/
type AggregateJiraRemediationsParams struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate jira remediations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateJiraRemediationsParams) WithDefaults() *AggregateJiraRemediationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate jira remediations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateJiraRemediationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) WithTimeout(timeout time.Duration) *AggregateJiraRemediationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) WithContext(ctx context.Context) *AggregateJiraRemediationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) WithHTTPClient(client *http.Client) *AggregateJiraRemediationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateJiraRemediationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate jira remediations params
func (o *AggregateJiraRemediationsParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateJiraRemediationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
