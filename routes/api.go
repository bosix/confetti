package routes

import (
	. "github.com/confetti-framework/foundation/http/routing"
	"src/app/http/controllers"
	"src/app/http/middleware"
)

/*
	|---------------------------------------------------------------------------
	| API routes
	|---------------------------------------------------------------------------
	|
	| Here is where you can register API routes for your application. By default
	| this is loaded in a group. The group is assigned to the "Api" middleware
	| and is placed with "/api" prefix. Feel free to remove the prefix if you
	| are using a subdomain for your API (which is recommended).
	| Enjoy building your API!
	|
*/

var Api = Group(
	Get("/ping", controllers.Ping),
).Prefix("/api").Middleware(middleware.Api...)
