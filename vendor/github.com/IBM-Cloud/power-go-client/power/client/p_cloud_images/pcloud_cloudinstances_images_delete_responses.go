// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesImagesDeleteReader is a Reader for the PcloudCloudinstancesImagesDelete structure.
type PcloudCloudinstancesImagesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesImagesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesImagesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesImagesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesImagesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesImagesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesImagesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudinstancesImagesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesImagesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesDeleteOK creates a PcloudCloudinstancesImagesDeleteOK with default headers values
func NewPcloudCloudinstancesImagesDeleteOK() *PcloudCloudinstancesImagesDeleteOK {
	return &PcloudCloudinstancesImagesDeleteOK{}
}

/*
PcloudCloudinstancesImagesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesImagesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances images delete o k response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances images delete o k response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete o k response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images delete o k response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete o k response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesImagesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteBadRequest creates a PcloudCloudinstancesImagesDeleteBadRequest with default headers values
func NewPcloudCloudinstancesImagesDeleteBadRequest() *PcloudCloudinstancesImagesDeleteBadRequest {
	return &PcloudCloudinstancesImagesDeleteBadRequest{}
}

/*
PcloudCloudinstancesImagesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete bad request response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete bad request response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete bad request response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images delete bad request response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete bad request response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteUnauthorized creates a PcloudCloudinstancesImagesDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesImagesDeleteUnauthorized() *PcloudCloudinstancesImagesDeleteUnauthorized {
	return &PcloudCloudinstancesImagesDeleteUnauthorized{}
}

/*
PcloudCloudinstancesImagesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images delete unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteForbidden creates a PcloudCloudinstancesImagesDeleteForbidden with default headers values
func NewPcloudCloudinstancesImagesDeleteForbidden() *PcloudCloudinstancesImagesDeleteForbidden {
	return &PcloudCloudinstancesImagesDeleteForbidden{}
}

/*
PcloudCloudinstancesImagesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesImagesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete forbidden response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete forbidden response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete forbidden response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images delete forbidden response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete forbidden response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudCloudinstancesImagesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteNotFound creates a PcloudCloudinstancesImagesDeleteNotFound with default headers values
func NewPcloudCloudinstancesImagesDeleteNotFound() *PcloudCloudinstancesImagesDeleteNotFound {
	return &PcloudCloudinstancesImagesDeleteNotFound{}
}

/*
PcloudCloudinstancesImagesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesImagesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete not found response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete not found response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete not found response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images delete not found response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete not found response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudinstancesImagesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteGone creates a PcloudCloudinstancesImagesDeleteGone with default headers values
func NewPcloudCloudinstancesImagesDeleteGone() *PcloudCloudinstancesImagesDeleteGone {
	return &PcloudCloudinstancesImagesDeleteGone{}
}

/*
PcloudCloudinstancesImagesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudinstancesImagesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete gone response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete gone response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete gone response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images delete gone response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images delete gone response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteGone) IsCode(code int) bool {
	return code == 410
}

func (o *PcloudCloudinstancesImagesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesDeleteInternalServerError creates a PcloudCloudinstancesImagesDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesImagesDeleteInternalServerError() *PcloudCloudinstancesImagesDeleteInternalServerError {
	return &PcloudCloudinstancesImagesDeleteInternalServerError{}
}

/*
PcloudCloudinstancesImagesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images delete internal server error response has a 2xx status code
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images delete internal server error response has a 3xx status code
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images delete internal server error response has a 4xx status code
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images delete internal server error response has a 5xx status code
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances images delete internal server error response a status code equal to that given
func (o *PcloudCloudinstancesImagesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
