// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesConsoleGetReader is a Reader for the PcloudPvminstancesConsoleGet structure.
type PcloudPvminstancesConsoleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsoleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesConsoleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesConsoleGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesConsoleGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesConsoleGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesConsoleGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesConsoleGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesConsoleGetOK creates a PcloudPvminstancesConsoleGetOK with default headers values
func NewPcloudPvminstancesConsoleGetOK() *PcloudPvminstancesConsoleGetOK {
	return &PcloudPvminstancesConsoleGetOK{}
}

/*
PcloudPvminstancesConsoleGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesConsoleGetOK struct {
	Payload *models.ConsoleLanguages
}

// IsSuccess returns true when this pcloud pvminstances console get o k response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances console get o k response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get o k response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console get o k response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console get o k response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesConsoleGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetOK) GetPayload() *models.ConsoleLanguages {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleLanguages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetBadRequest creates a PcloudPvminstancesConsoleGetBadRequest with default headers values
func NewPcloudPvminstancesConsoleGetBadRequest() *PcloudPvminstancesConsoleGetBadRequest {
	return &PcloudPvminstancesConsoleGetBadRequest{}
}

/*
PcloudPvminstancesConsoleGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesConsoleGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console get bad request response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console get bad request response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get bad request response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console get bad request response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console get bad request response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesConsoleGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetUnauthorized creates a PcloudPvminstancesConsoleGetUnauthorized with default headers values
func NewPcloudPvminstancesConsoleGetUnauthorized() *PcloudPvminstancesConsoleGetUnauthorized {
	return &PcloudPvminstancesConsoleGetUnauthorized{}
}

/*
PcloudPvminstancesConsoleGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesConsoleGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console get unauthorized response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console get unauthorized response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get unauthorized response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console get unauthorized response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console get unauthorized response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetForbidden creates a PcloudPvminstancesConsoleGetForbidden with default headers values
func NewPcloudPvminstancesConsoleGetForbidden() *PcloudPvminstancesConsoleGetForbidden {
	return &PcloudPvminstancesConsoleGetForbidden{}
}

/*
PcloudPvminstancesConsoleGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesConsoleGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console get forbidden response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console get forbidden response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get forbidden response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console get forbidden response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console get forbidden response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPvminstancesConsoleGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetNotFound creates a PcloudPvminstancesConsoleGetNotFound with default headers values
func NewPcloudPvminstancesConsoleGetNotFound() *PcloudPvminstancesConsoleGetNotFound {
	return &PcloudPvminstancesConsoleGetNotFound{}
}

/*
PcloudPvminstancesConsoleGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesConsoleGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console get not found response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console get not found response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get not found response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console get not found response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console get not found response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPvminstancesConsoleGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetInternalServerError creates a PcloudPvminstancesConsoleGetInternalServerError with default headers values
func NewPcloudPvminstancesConsoleGetInternalServerError() *PcloudPvminstancesConsoleGetInternalServerError {
	return &PcloudPvminstancesConsoleGetInternalServerError{}
}

/*
PcloudPvminstancesConsoleGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsoleGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console get internal server error response has a 2xx status code
func (o *PcloudPvminstancesConsoleGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console get internal server error response has a 3xx status code
func (o *PcloudPvminstancesConsoleGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console get internal server error response has a 4xx status code
func (o *PcloudPvminstancesConsoleGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console get internal server error response has a 5xx status code
func (o *PcloudPvminstancesConsoleGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances console get internal server error response a status code equal to that given
func (o *PcloudPvminstancesConsoleGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
