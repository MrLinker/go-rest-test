// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindPetsByStatusHandlerFunc turns a function with the right signature into a find pets by status handler
type FindPetsByStatusHandlerFunc func(FindPetsByStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn FindPetsByStatusHandlerFunc) Handle(params FindPetsByStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// FindPetsByStatusHandler interface for that can handle valid find pets by status params
type FindPetsByStatusHandler interface {
	Handle(FindPetsByStatusParams, interface{}) middleware.Responder
}

// NewFindPetsByStatus creates a new http.Handler for the find pets by status operation
func NewFindPetsByStatus(ctx *middleware.Context, handler FindPetsByStatusHandler) *FindPetsByStatus {
	return &FindPetsByStatus{Context: ctx, Handler: handler}
}

/*FindPetsByStatus swagger:route GET /pet/findByStatus pet findPetsByStatus

Finds Pets by status

Multiple status values can be provided with comma separated strings

*/
type FindPetsByStatus struct {
	Context *middleware.Context
	Handler FindPetsByStatusHandler
}

func (o *FindPetsByStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindPetsByStatusParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
