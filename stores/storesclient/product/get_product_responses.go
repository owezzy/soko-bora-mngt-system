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

// GetProductReader is a Reader for the GetProduct structure.
type GetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductOK creates a GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {
	return &GetProductOK{}
}

/*
GetProductOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetProductOK struct {
	Payload *models.StorespbGetProductResponse
}

// IsSuccess returns true when this get product o k response has a 2xx status code
func (o *GetProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get product o k response has a 3xx status code
func (o *GetProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get product o k response has a 4xx status code
func (o *GetProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get product o k response has a 5xx status code
func (o *GetProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get product o k response a status code equal to that given
func (o *GetProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get product o k response
func (o *GetProductOK) Code() int {
	return 200
}

func (o *GetProductOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/products/{id}][%d] getProductOK %s", 200, payload)
}

func (o *GetProductOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/products/{id}][%d] getProductOK %s", 200, payload)
}

func (o *GetProductOK) GetPayload() *models.StorespbGetProductResponse {
	return o.Payload
}

func (o *GetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorespbGetProductResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductDefault creates a GetProductDefault with default headers values
func NewGetProductDefault(code int) *GetProductDefault {
	return &GetProductDefault{
		_statusCode: code,
	}
}

/*
GetProductDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetProductDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get product default response has a 2xx status code
func (o *GetProductDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get product default response has a 3xx status code
func (o *GetProductDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get product default response has a 4xx status code
func (o *GetProductDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get product default response has a 5xx status code
func (o *GetProductDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get product default response a status code equal to that given
func (o *GetProductDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get product default response
func (o *GetProductDefault) Code() int {
	return o._statusCode
}

func (o *GetProductDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/products/{id}][%d] getProduct default %s", o._statusCode, payload)
}

func (o *GetProductDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/products/{id}][%d] getProduct default %s", o._statusCode, payload)
}

func (o *GetProductDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
