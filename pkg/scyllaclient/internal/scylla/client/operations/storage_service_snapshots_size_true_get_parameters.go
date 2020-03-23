// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceSnapshotsSizeTrueGetParams creates a new StorageServiceSnapshotsSizeTrueGetParams object
// with the default values initialized.
func NewStorageServiceSnapshotsSizeTrueGetParams() *StorageServiceSnapshotsSizeTrueGetParams {

	return &StorageServiceSnapshotsSizeTrueGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceSnapshotsSizeTrueGetParamsWithTimeout creates a new StorageServiceSnapshotsSizeTrueGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceSnapshotsSizeTrueGetParamsWithTimeout(timeout time.Duration) *StorageServiceSnapshotsSizeTrueGetParams {

	return &StorageServiceSnapshotsSizeTrueGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceSnapshotsSizeTrueGetParamsWithContext creates a new StorageServiceSnapshotsSizeTrueGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceSnapshotsSizeTrueGetParamsWithContext(ctx context.Context) *StorageServiceSnapshotsSizeTrueGetParams {

	return &StorageServiceSnapshotsSizeTrueGetParams{

		Context: ctx,
	}
}

// NewStorageServiceSnapshotsSizeTrueGetParamsWithHTTPClient creates a new StorageServiceSnapshotsSizeTrueGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceSnapshotsSizeTrueGetParamsWithHTTPClient(client *http.Client) *StorageServiceSnapshotsSizeTrueGetParams {

	return &StorageServiceSnapshotsSizeTrueGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceSnapshotsSizeTrueGetParams contains all the parameters to send to the API endpoint
for the storage service snapshots size true get operation typically these are written to a http.Request
*/
type StorageServiceSnapshotsSizeTrueGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) WithTimeout(timeout time.Duration) *StorageServiceSnapshotsSizeTrueGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) WithContext(ctx context.Context) *StorageServiceSnapshotsSizeTrueGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) WithHTTPClient(client *http.Client) *StorageServiceSnapshotsSizeTrueGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service snapshots size true get params
func (o *StorageServiceSnapshotsSizeTrueGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceSnapshotsSizeTrueGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}