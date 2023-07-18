package main

import (
	"github.com/oktarian/crud-restapi-go/controllers/productcontroller"
	"github.com/oktarian/crud-restapi-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//connection in setup models
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}