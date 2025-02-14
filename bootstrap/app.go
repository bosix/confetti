package bootstrap

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	net "net/http"
	"src/app/console"
	"src/app/http"
	"src/app/http/decorator"
)

var bootContainer inter.Container

func init() {

	/*
		|--------------------------------------------------------------------------
		| Create boot container
		|--------------------------------------------------------------------------
		|
		| The first thing we will do is create a new boot container instance
		| which is the IoC container for the system binding all binding that is
		| equal for each request. This container is readonly after init.
		|
	*/

	bootContainer = decorator.Bootstrap(foundation.NewContainer())
}

func NewAppFromBoot() inter.App {

	/*
		|--------------------------------------------------------------------------
		| Create The Application
		|--------------------------------------------------------------------------
		|
		| The second thing we will do is create a new Confetti application instance
		| which serves as the "glue" for all the components of Confetti, and is
		| the IoC container for the request binding all of the various parts.
		|
	*/

	container := foundation.NewContainerByBoot(bootContainer)
	app := foundation.NewApp()
	app.SetContainer(container)

	app.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(app),
	)

	app.Singleton(
		(*inter.ConsoleKernel)(nil),
		console.NewKernel(app),
	)

	app.Bind(
		(*net.HandlerFunc)(nil),
		HandleHttpKernel,
	)

	return app
}
