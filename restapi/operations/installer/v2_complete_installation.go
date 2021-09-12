// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2CompleteInstallationHandlerFunc turns a function with the right signature into a v2 complete installation handler
type V2CompleteInstallationHandlerFunc func(V2CompleteInstallationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2CompleteInstallationHandlerFunc) Handle(params V2CompleteInstallationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2CompleteInstallationHandler interface for that can handle valid v2 complete installation params
type V2CompleteInstallationHandler interface {
	Handle(V2CompleteInstallationParams, interface{}) middleware.Responder
}

// NewV2CompleteInstallation creates a new http.Handler for the v2 complete installation operation
func NewV2CompleteInstallation(ctx *middleware.Context, handler V2CompleteInstallationHandler) *V2CompleteInstallation {
	return &V2CompleteInstallation{Context: ctx, Handler: handler}
}

/*V2CompleteInstallation swagger:route POST /v2/clusters/{cluster_id}/actions/complete-installation installer v2CompleteInstallation

Agent API to mark a finalizing installation as complete.

*/
type V2CompleteInstallation struct {
	Context *middleware.Context
	Handler V2CompleteInstallationHandler
}

func (o *V2CompleteInstallation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewV2CompleteInstallationParams()

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
