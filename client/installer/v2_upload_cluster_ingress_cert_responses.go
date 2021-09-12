// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UploadClusterIngressCertReader is a Reader for the V2UploadClusterIngressCert structure.
type V2UploadClusterIngressCertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UploadClusterIngressCertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2UploadClusterIngressCertCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UploadClusterIngressCertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UploadClusterIngressCertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UploadClusterIngressCertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UploadClusterIngressCertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UploadClusterIngressCertMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UploadClusterIngressCertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2UploadClusterIngressCertServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UploadClusterIngressCertCreated creates a V2UploadClusterIngressCertCreated with default headers values
func NewV2UploadClusterIngressCertCreated() *V2UploadClusterIngressCertCreated {
	return &V2UploadClusterIngressCertCreated{}
}

/*V2UploadClusterIngressCertCreated handles this case with default header values.

Success.
*/
type V2UploadClusterIngressCertCreated struct {
}

func (o *V2UploadClusterIngressCertCreated) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertCreated ", 201)
}

func (o *V2UploadClusterIngressCertCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2UploadClusterIngressCertBadRequest creates a V2UploadClusterIngressCertBadRequest with default headers values
func NewV2UploadClusterIngressCertBadRequest() *V2UploadClusterIngressCertBadRequest {
	return &V2UploadClusterIngressCertBadRequest{}
}

/*V2UploadClusterIngressCertBadRequest handles this case with default header values.

Error.
*/
type V2UploadClusterIngressCertBadRequest struct {
	Payload *models.Error
}

func (o *V2UploadClusterIngressCertBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertBadRequest  %+v", 400, o.Payload)
}

func (o *V2UploadClusterIngressCertBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadClusterIngressCertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertUnauthorized creates a V2UploadClusterIngressCertUnauthorized with default headers values
func NewV2UploadClusterIngressCertUnauthorized() *V2UploadClusterIngressCertUnauthorized {
	return &V2UploadClusterIngressCertUnauthorized{}
}

/*V2UploadClusterIngressCertUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2UploadClusterIngressCertUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2UploadClusterIngressCertUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UploadClusterIngressCertUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UploadClusterIngressCertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertForbidden creates a V2UploadClusterIngressCertForbidden with default headers values
func NewV2UploadClusterIngressCertForbidden() *V2UploadClusterIngressCertForbidden {
	return &V2UploadClusterIngressCertForbidden{}
}

/*V2UploadClusterIngressCertForbidden handles this case with default header values.

Forbidden.
*/
type V2UploadClusterIngressCertForbidden struct {
	Payload *models.InfraError
}

func (o *V2UploadClusterIngressCertForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertForbidden  %+v", 403, o.Payload)
}

func (o *V2UploadClusterIngressCertForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UploadClusterIngressCertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertNotFound creates a V2UploadClusterIngressCertNotFound with default headers values
func NewV2UploadClusterIngressCertNotFound() *V2UploadClusterIngressCertNotFound {
	return &V2UploadClusterIngressCertNotFound{}
}

/*V2UploadClusterIngressCertNotFound handles this case with default header values.

Error.
*/
type V2UploadClusterIngressCertNotFound struct {
	Payload *models.Error
}

func (o *V2UploadClusterIngressCertNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertNotFound  %+v", 404, o.Payload)
}

func (o *V2UploadClusterIngressCertNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadClusterIngressCertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertMethodNotAllowed creates a V2UploadClusterIngressCertMethodNotAllowed with default headers values
func NewV2UploadClusterIngressCertMethodNotAllowed() *V2UploadClusterIngressCertMethodNotAllowed {
	return &V2UploadClusterIngressCertMethodNotAllowed{}
}

/*V2UploadClusterIngressCertMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2UploadClusterIngressCertMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2UploadClusterIngressCertMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UploadClusterIngressCertMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadClusterIngressCertMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertInternalServerError creates a V2UploadClusterIngressCertInternalServerError with default headers values
func NewV2UploadClusterIngressCertInternalServerError() *V2UploadClusterIngressCertInternalServerError {
	return &V2UploadClusterIngressCertInternalServerError{}
}

/*V2UploadClusterIngressCertInternalServerError handles this case with default header values.

Error.
*/
type V2UploadClusterIngressCertInternalServerError struct {
	Payload *models.Error
}

func (o *V2UploadClusterIngressCertInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UploadClusterIngressCertInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadClusterIngressCertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadClusterIngressCertServiceUnavailable creates a V2UploadClusterIngressCertServiceUnavailable with default headers values
func NewV2UploadClusterIngressCertServiceUnavailable() *V2UploadClusterIngressCertServiceUnavailable {
	return &V2UploadClusterIngressCertServiceUnavailable{}
}

/*V2UploadClusterIngressCertServiceUnavailable handles this case with default header values.

Unavailable.
*/
type V2UploadClusterIngressCertServiceUnavailable struct {
	Payload *models.Error
}

func (o *V2UploadClusterIngressCertServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/uploads/ingress-cert][%d] v2UploadClusterIngressCertServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2UploadClusterIngressCertServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadClusterIngressCertServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
