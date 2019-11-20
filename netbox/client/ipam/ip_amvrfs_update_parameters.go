// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewIPAMVrfsUpdateParams creates a new IPAMVrfsUpdateParams object
// with the default values initialized.
func NewIPAMVrfsUpdateParams() *IPAMVrfsUpdateParams {
	var ()
	return &IPAMVrfsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVrfsUpdateParamsWithTimeout creates a new IPAMVrfsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVrfsUpdateParamsWithTimeout(timeout time.Duration) *IPAMVrfsUpdateParams {
	var ()
	return &IPAMVrfsUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMVrfsUpdateParamsWithContext creates a new IPAMVrfsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVrfsUpdateParamsWithContext(ctx context.Context) *IPAMVrfsUpdateParams {
	var ()
	return &IPAMVrfsUpdateParams{

		Context: ctx,
	}
}

// NewIPAMVrfsUpdateParamsWithHTTPClient creates a new IPAMVrfsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVrfsUpdateParamsWithHTTPClient(client *http.Client) *IPAMVrfsUpdateParams {
	var ()
	return &IPAMVrfsUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMVrfsUpdateParams contains all the parameters to send to the API endpoint
for the ipam vrfs update operation typically these are written to a http.Request
*/
type IPAMVrfsUpdateParams struct {

	/*Data*/
	Data *models.WritableVRF
	/*ID
	  A unique integer value identifying this VRF.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) WithTimeout(timeout time.Duration) *IPAMVrfsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) WithContext(ctx context.Context) *IPAMVrfsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) WithHTTPClient(client *http.Client) *IPAMVrfsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) WithData(data *models.WritableVRF) *IPAMVrfsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) SetData(data *models.WritableVRF) {
	o.Data = data
}

// WithID adds the id to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) WithID(id int64) *IPAMVrfsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs update params
func (o *IPAMVrfsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVrfsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}