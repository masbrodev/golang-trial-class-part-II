package main

import (
	"api-mini-ecommerce/config"
	"api-mini-ecommerce/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnect()
	s := gin.Default()
	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Welcome",
		})
	})

	s.GET("/product", controller.HandlerGetProduct)
	s.POST("orders", controller.HandlerPostProduct)
	s.GET("/orders", controller.HandlerGetOrder)
	s.Run()
}
