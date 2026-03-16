package routes

import (
	"web_app/app/controllers"

	"github.com/labstack/echo/v4"
)

func RoutesRegister(e *echo.Echo) {
	// root route
	e.GET("", controllers.RedirectToInfo)

	// ================ View routes =================
	pageRoute := e.Group("/page")
	pageRoute.GET("/home", controllers.ShowPage)
	// ================ API routes =================
	apiRoute := e.Group("/api")
	v1 := apiRoute.Group("/v1")
	// Redirect to information about the API
	v1.GET("", controllers.RedirectToInfo)
	// Create a users group
	users := v1.Group("/users")
	// Redirect route for users
	users.GET("", controllers.RedirectToInfo)
	// Information about the users API
	users.GET("/info", controllers.Information)
	// CRUD routes
	users.GET("/all", controllers.GetUsers)          // list
	users.POST("", controllers.CreateUser)           // create
	users.GET("/:id", controllers.GetUser)           // get by id
	users.PUT("/:id", controllers.UpdateUser)        // full update
	users.PATCH("/:id", controllers.UpdateUser)      // partial update (same handler)
	users.DELETE("/:id", controllers.DeleteUser)     // delete
	users.GET("/hello", controllers.HelloWorld)      // demo
	// create error handler for error routes
	e.HTTPErrorHandler = controllers.ErrorHandler
}

