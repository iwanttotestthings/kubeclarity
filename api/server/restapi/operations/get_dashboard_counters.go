// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDashboardCountersHandlerFunc turns a function with the right signature into a get dashboard counters handler
type GetDashboardCountersHandlerFunc func(GetDashboardCountersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDashboardCountersHandlerFunc) Handle(params GetDashboardCountersParams) middleware.Responder {
	return fn(params)
}

// GetDashboardCountersHandler interface for that can handle valid get dashboard counters params
type GetDashboardCountersHandler interface {
	Handle(GetDashboardCountersParams) middleware.Responder
}

// NewGetDashboardCounters creates a new http.Handler for the get dashboard counters operation
func NewGetDashboardCounters(ctx *middleware.Context, handler GetDashboardCountersHandler) *GetDashboardCounters {
	return &GetDashboardCounters{Context: ctx, Handler: handler}
}

/* GetDashboardCounters swagger:route GET /dashboard/counters getDashboardCounters

Get number of applications, resources, packages and vulnerabilities

*/
type GetDashboardCounters struct {
	Context *middleware.Context
	Handler GetDashboardCountersHandler
}

func (o *GetDashboardCounters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDashboardCountersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
