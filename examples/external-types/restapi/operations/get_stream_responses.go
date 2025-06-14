// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/istforks/go-swagger/examples/external-types/models"
)

/*
GetStreamDefault Uses an external definition for an interface (e.g. io.Reader)

No validation is expected on binary format.

swagger:response getStreamDefault
*/
type GetStreamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.MyStreamer `json:"body,omitempty"`
}

// NewGetStreamDefault creates GetStreamDefault with default headers values
func NewGetStreamDefault(code int) *GetStreamDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStreamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stream default response
func (o *GetStreamDefault) WithStatusCode(code int) *GetStreamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stream default response
func (o *GetStreamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get stream default response
func (o *GetStreamDefault) WithPayload(payload models.MyStreamer) *GetStreamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stream default response
func (o *GetStreamDefault) SetPayload(payload models.MyStreamer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStreamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
