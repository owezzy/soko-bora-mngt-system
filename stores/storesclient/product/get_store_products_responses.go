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

// GetStoreProductsReader is a Reader for the GetStoreProducts structure.
type GetStoreProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoreProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStoreProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStoreProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoreProductsOK creates a GetStoreProductsOK with default headers values
func NewGetStoreProductsOK() *GetStoreProductsOK {
	return &GetStoreProductsOK{}
}

/*
GetStoreProductsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetStoreProductsOK struct {
	Payload *models.StorespbGetCatalogResponse
}

// IsSuccess returns true when this get store products o k response has a 2xx status code
func (o *GetStoreProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get store products o k response has a 3xx status code
func (o *GetStoreProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get store products o k response has a 4xx status code
func (o *GetStoreProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get store products o k response has a 5xx status code
func (o *GetStoreProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get store products o k response a status code equal to that given
func (o *GetStoreProductsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get store products o k response
func (o *GetStoreProductsOK) Code() int {
	return 200
}

func (o *GetStoreProductsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/{storeId}/products][%d] getStoreProductsOK %s", 200, payload)
}

func (o *GetStoreProductsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/{storeId}/products][%d] getStoreProductsOK %s", 200, payload)
}

func (o *GetStoreProductsOK) GetPayload() *models.StorespbGetCatalogResponse {
	return o.Payload
}

func (o *GetStoreProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorespbGetCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoreProductsDefault creates a GetStoreProductsDefault with default headers values
func NewGetStoreProductsDefault(code int) *GetStoreProductsDefault {
	return &GetStoreProductsDefault{
		_statusCode: code,
	}
}

/*
GetStoreProductsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetStoreProductsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this get store products default response has a 2xx status code
func (o *GetStoreProductsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get store products default response has a 3xx status code
func (o *GetStoreProductsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get store products default response has a 4xx status code
func (o *GetStoreProductsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get store products default response has a 5xx status code
func (o *GetStoreProductsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get store products default response a status code equal to that given
func (o *GetStoreProductsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get store products default response
func (o *GetStoreProductsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoreProductsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/{storeId}/products][%d] getStoreProducts default %s", o._statusCode, payload)
}

func (o *GetStoreProductsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/stores/{storeId}/products][%d] getStoreProducts default %s", o._statusCode, payload)
}

func (o *GetStoreProductsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetStoreProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
