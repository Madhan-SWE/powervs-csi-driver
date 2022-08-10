// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_disaster_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudLocationsDisasterrecoveryGetReader is a Reader for the PcloudLocationsDisasterrecoveryGet structure.
type PcloudLocationsDisasterrecoveryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudLocationsDisasterrecoveryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudLocationsDisasterrecoveryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudLocationsDisasterrecoveryGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudLocationsDisasterrecoveryGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudLocationsDisasterrecoveryGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudLocationsDisasterrecoveryGetOK creates a PcloudLocationsDisasterrecoveryGetOK with default headers values
func NewPcloudLocationsDisasterrecoveryGetOK() *PcloudLocationsDisasterrecoveryGetOK {
	return &PcloudLocationsDisasterrecoveryGetOK{}
}

/* PcloudLocationsDisasterrecoveryGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudLocationsDisasterrecoveryGetOK struct {
	Payload *models.DisasterRecoveryLocation
}

func (o *PcloudLocationsDisasterrecoveryGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetOK  %+v", 200, o.Payload)
}
func (o *PcloudLocationsDisasterrecoveryGetOK) GetPayload() *models.DisasterRecoveryLocation {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DisasterRecoveryLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetUnauthorized creates a PcloudLocationsDisasterrecoveryGetUnauthorized with default headers values
func NewPcloudLocationsDisasterrecoveryGetUnauthorized() *PcloudLocationsDisasterrecoveryGetUnauthorized {
	return &PcloudLocationsDisasterrecoveryGetUnauthorized{}
}

/* PcloudLocationsDisasterrecoveryGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudLocationsDisasterrecoveryGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetNotFound creates a PcloudLocationsDisasterrecoveryGetNotFound with default headers values
func NewPcloudLocationsDisasterrecoveryGetNotFound() *PcloudLocationsDisasterrecoveryGetNotFound {
	return &PcloudLocationsDisasterrecoveryGetNotFound{}
}

/* PcloudLocationsDisasterrecoveryGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudLocationsDisasterrecoveryGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudLocationsDisasterrecoveryGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetInternalServerError creates a PcloudLocationsDisasterrecoveryGetInternalServerError with default headers values
func NewPcloudLocationsDisasterrecoveryGetInternalServerError() *PcloudLocationsDisasterrecoveryGetInternalServerError {
	return &PcloudLocationsDisasterrecoveryGetInternalServerError{}
}

/* PcloudLocationsDisasterrecoveryGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudLocationsDisasterrecoveryGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
