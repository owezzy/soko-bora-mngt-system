// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owezzy/soko-bora-mngt-system/stores/storesclient/models"
)

// RemoveProductReader is a Reader for the RemoveProduct structure.
type RemoveProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveProductOK creates a RemoveProductOK with default headers values
func NewRemoveProductOK() *RemoveProductOK {
	return &RemoveProductOK{}
}

/*
RemoveProductOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveProductOK struct {
	Payload models.StorespbRemoveProductResponse
}

// IsSuccess returns true when this remove product o k response has a 2xx status code
func (o *RemoveProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove product o k response has a 3xx status code
func (o *RemoveProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove product o k response has a 4xx status code
func (o *RemoveProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove product o k response has a 5xx status code
func (o *RemoveProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove product o k response a status code equal to that given
func (o *RemoveProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove product o k response
func (o *RemoveProductOK) Code() int {
	return 200
}

func (o *RemoveProductOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProductOK %s", 200, payload)
}

func (o *RemoveProductOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProductOK %s", 200, payload)
}

func (o *RemoveProductOK) GetPayload() models.StorespbRemoveProductResponse {
	return o.Payload
}

func (o *RemoveProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveProductDefault creates a RemoveProductDefault with default headers values
func NewRemoveProductDefault(code int) *RemoveProductDefault {
	return &RemoveProductDefault{
		_statusCode: code,
	}
}

/*
RemoveProductDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveProductDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this remove product default response has a 2xx status code
func (o *RemoveProductDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove product default response has a 3xx status code
func (o *RemoveProductDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove product default response has a 4xx status code
func (o *RemoveProductDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove product default response has a 5xx status code
func (o *RemoveProductDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove product default response a status code equal to that given
func (o *RemoveProductDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove product default response
func (o *RemoveProductDefault) Code() int {
	return o._statusCode
}

func (o *RemoveProductDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProduct default %s", o._statusCode, payload)
}

func (o *RemoveProductDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/stores/products/{id}][%d] removeProduct default %s", o._statusCode, payload)
}

func (o *RemoveProductDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RemoveProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
