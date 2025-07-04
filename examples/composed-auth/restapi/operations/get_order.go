// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/istforks/go-swagger/examples/composed-auth/models"
)

// GetOrderHandlerFunc turns a function with the right signature into a get order handler
type GetOrderHandlerFunc func(GetOrderParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOrderHandlerFunc) Handle(params GetOrderParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetOrderHandler interface for that can handle valid get order params
type GetOrderHandler interface {
	Handle(GetOrderParams, *models.Principal) middleware.Responder
}

// NewGetOrder creates a new http.Handler for the get order operation
func NewGetOrder(ctx *middleware.Context, handler GetOrderHandler) *GetOrder {
	return &GetOrder{Context: ctx, Handler: handler}
}

/*
	GetOrder swagger:route GET /order/{orderID} getOrder

retrieves an order

Only registered customers should be able to retrieve orders
*/
type GetOrder struct {
	Context *middleware.Context
	Handler GetOrderHandler
}

func (o *GetOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetOrderParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
