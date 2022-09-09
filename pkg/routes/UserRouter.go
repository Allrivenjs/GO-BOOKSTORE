package routes

import (
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func(Routes *gin.Engine) {
	//Routes.Use(middleware.Authenticate())
	Routes.GET("/users", controllers.GetUsers())
	Routes.GET("/users/:user_id", controllers.GetUser())
}
