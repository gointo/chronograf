package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostSourcesHandlerFunc turns a function with the right signature into a post sources handler
type PostSourcesHandlerFunc func(context.Context, PostSourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSourcesHandlerFunc) Handle(ctx context.Context, params PostSourcesParams) middleware.Responder {
	return fn(ctx, params)
}

// PostSourcesHandler interface for that can handle valid post sources params
type PostSourcesHandler interface {
	Handle(context.Context, PostSourcesParams) middleware.Responder
}

// NewPostSources creates a new http.Handler for the post sources operation
func NewPostSources(ctx *middleware.Context, handler PostSourcesHandler) *PostSources {
	return &PostSources{Context: ctx, Handler: handler}
}

/*PostSources swagger:route POST /sources postSources

Create new data source

*/
type PostSources struct {
	Context *middleware.Context
	Handler PostSourcesHandler
}

func (o *PostSources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostSourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}