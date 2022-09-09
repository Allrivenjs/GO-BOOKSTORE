package main

import (
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.New()
	r.Use(gin.Logger())
	//routes.RegisterBookStoreRoutes(r)
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	log.Fatal(r.Run(":" + port))

}
