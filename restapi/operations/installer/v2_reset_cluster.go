// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2ResetClusterHandlerFunc turns a function with the right signature into a v2 reset cluster handler
type V2ResetClusterHandlerFunc func(V2ResetClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2ResetClusterHandlerFunc) Handle(params V2ResetClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2ResetClusterHandler interface for that can handle valid v2 reset cluster params
type V2ResetClusterHandler interface {
	Handle(V2ResetClusterParams, interface{}) middleware.Responder
}

// NewV2ResetCluster creates a new http.Handler for the v2 reset cluster operation
func NewV2ResetCluster(ctx *middleware.Context, handler V2ResetClusterHandler) *V2ResetCluster {
	return &V2ResetCluster{Context: ctx, Handler: handler}
}

/*V2ResetCluster swagger:route POST /v2/clusters/{cluster_id}/actions/reset installer v2ResetCluster

Resets a failed installation.

*/
type V2ResetCluster struct {
	Context *middleware.Context
	Handler V2ResetClusterHandler
}

func (o *V2ResetCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV2ResetClusterParams()

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
