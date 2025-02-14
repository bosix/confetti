package providers

import (
	"github.com/confetti-framework/contract/inter"
	foundationMiddleware "github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/routing"
	"src/app/http/middleware"
	"src/routes"
)

var globalMiddlewares = []inter.HttpMiddleware{
	foundationMiddleware.RequestBodyDecoder{},
	middleware.RouteModelBinding{},
}

type RouteServiceProvider struct{}

// Define your router model bindings, pattern filters, etc.
func (r RouteServiceProvider) Boot(container inter.Container) inter.Container {
	collection := routing.NewRouteCollection()

	collection.Merge(routes.Api)
	collection.Merge(routes.Web)

	// Here you can make your adjustments that apply to all routes:
	// collection.WhereMulti(map[string]string{
	// 	"id": "[0-9]+",
	// })

	// These middlewares should be executed on all routes
	collection.Middleware(globalMiddlewares...)

	routing.DecorateRoutes(collection)
	container.Singleton("routes", collection)

	return container
}
