// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/istforks/go-swagger/examples/contributed-templates/stratoscale/models"
)

// OrderCreateOKCode is the HTTP code returned for type OrderCreateOK
const OrderCreateOKCode int = 200

/*
OrderCreateOK successful operation

swagger:response orderCreateOK
*/
type OrderCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Order `json:"body,omitempty"`
}

// NewOrderCreateOK creates OrderCreateOK with default headers values
func NewOrderCreateOK() *OrderCreateOK {

	return &OrderCreateOK{}
}

// WithPayload adds the payload to the order create o k response
func (o *OrderCreateOK) WithPayload(payload *models.Order) *OrderCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the order create o k response
func (o *OrderCreateOK) SetPayload(payload *models.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OrderCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// OrderCreateBadRequestCode is the HTTP code returned for type OrderCreateBadRequest
const OrderCreateBadRequestCode int = 400

/*
OrderCreateBadRequest Invalid Order

swagger:response orderCreateBadRequest
*/
type OrderCreateBadRequest struct {
}

// NewOrderCreateBadRequest creates OrderCreateBadRequest with default headers values
func NewOrderCreateBadRequest() *OrderCreateBadRequest {

	return &OrderCreateBadRequest{}
}

// WriteResponse to the client
func (o *OrderCreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
