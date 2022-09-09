package routes

import (
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var AuthRoutes = func(Routes *gin.Engine) {
	Routes.POST("users/signup", controllers.Signup())
	//Routes.POST("users/login", controllers.Login())
}
