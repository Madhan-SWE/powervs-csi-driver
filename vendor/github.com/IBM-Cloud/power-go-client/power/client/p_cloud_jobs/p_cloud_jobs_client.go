// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudCloudinstancesJobsDelete(params *PcloudCloudinstancesJobsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsDeleteOK, error)

	PcloudCloudinstancesJobsGet(params *PcloudCloudinstancesJobsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsGetOK, error)

	PcloudCloudinstancesJobsGetall(params *PcloudCloudinstancesJobsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsGetallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PcloudCloudinstancesJobsDelete deletes a cloud instance job
*/
func (a *Client) PcloudCloudinstancesJobsDelete(params *PcloudCloudinstancesJobsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesJobsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.jobs.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesJobsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudinstancesJobsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudinstances.jobs.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudinstancesJobsGet lists the detail of a job
*/
func (a *Client) PcloudCloudinstancesJobsGet(params *PcloudCloudinstancesJobsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesJobsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.jobs.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesJobsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudinstancesJobsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudinstances.jobs.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PcloudCloudinstancesJobsGetall lists up to the last 5 jobs initiated by the cloud instance
*/
func (a *Client) PcloudCloudinstancesJobsGetall(params *PcloudCloudinstancesJobsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudCloudinstancesJobsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudCloudinstancesJobsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.cloudinstances.jobs.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudCloudinstancesJobsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudCloudinstancesJobsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.cloudinstances.jobs.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
