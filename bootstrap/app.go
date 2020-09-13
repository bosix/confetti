package bootstrap

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"lanvard/app/console"
	"lanvard/app/exception"
	"lanvard/app/http"
	"lanvard/app/http/decorator"
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
	// bootContainer.BindPathsInContainer(config.App.BasePath)
	// bootContainer.Instance("env", config.App.Env)
}

func NewAppFromBoot() inter.App {

	/*
		|--------------------------------------------------------------------------
		| Create The Application
		|--------------------------------------------------------------------------
		|
		| The second thing we will do is create a new Lanvard application instance
		| which serves as the "glue" for all the components of Lanvard, and is
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

	app.Singleton(
		(*inter.ExceptionHandler)(nil),
		exception.NewHandler(app),
	)

	return app
}
