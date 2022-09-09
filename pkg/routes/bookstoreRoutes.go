package routes

import (
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterBookStoreRoutes sett routes for book controller
var RegisterBookStoreRoutes = func(router *gin.Engine) {
	router.POST("/book/", controllers.CreateBook())
	router.GET("/book/", controllers.GetBook())
	router.GET("/book/{bookId}", controllers.GetBookById())
	router.PUT("/book/{bookId}", controllers.UpdateBook())
	router.DELETE("/book/{bookId}", controllers.DeleteBook())

}
