package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSourcesIDUsersUserIDExplorationsExplorationIDHandlerFunc turns a function with the right signature into a get sources ID users user ID explorations exploration ID handler
type GetSourcesIDUsersUserIDExplorationsExplorationIDHandlerFunc func(context.Context, GetSourcesIDUsersUserIDExplorationsExplorationIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSourcesIDUsersUserIDExplorationsExplorationIDHandlerFunc) Handle(ctx context.Context, params GetSourcesIDUsersUserIDExplorationsExplorationIDParams) middleware.Responder {
	return fn(ctx, params)
}

// GetSourcesIDUsersUserIDExplorationsExplorationIDHandler interface for that can handle valid get sources ID users user ID explorations exploration ID params
type GetSourcesIDUsersUserIDExplorationsExplorationIDHandler interface {
	Handle(context.Context, GetSourcesIDUsersUserIDExplorationsExplorationIDParams) middleware.Responder
}

// NewGetSourcesIDUsersUserIDExplorationsExplorationID creates a new http.Handler for the get sources ID users user ID explorations exploration ID operation
func NewGetSourcesIDUsersUserIDExplorationsExplorationID(ctx *middleware.Context, handler GetSourcesIDUsersUserIDExplorationsExplorationIDHandler) *GetSourcesIDUsersUserIDExplorationsExplorationID {
	return &GetSourcesIDUsersUserIDExplorationsExplorationID{Context: ctx, Handler: handler}
}

/*GetSourcesIDUsersUserIDExplorationsExplorationID swagger:route GET /sources/{id}/users/{user_id}/explorations/{exploration_id} getSourcesIdUsersUserIdExplorationsExplorationId

Returns the specified data exploration session

A data exploration session specifies query information.


*/
type GetSourcesIDUsersUserIDExplorationsExplorationID struct {
	Context *middleware.Context
	Handler GetSourcesIDUsersUserIDExplorationsExplorationIDHandler
}

func (o *GetSourcesIDUsersUserIDExplorationsExplorationID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetSourcesIDUsersUserIDExplorationsExplorationIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}