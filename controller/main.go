package controller

import (
	"api-mini-ecommerce/config"
	"api-mini-ecommerce/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductId    int    `json:"product_id"`
}

func HandlerGetProduct(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func HandlerPostProduct(ctx *gin.Context) {
	var orderBody OrderRequest
	err := ctx.ShouldBindJSON(&orderBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var product entity.Product
	result := config.DB.Where("ID = ?", orderBody.ProductId).First(&product)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Message": "Product Not Found",
		})
		return
	}

	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductID:    uint(orderBody.ProductId),
		OrderTime:    time.Now(),
	}
	err = config.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Message": "Error",
		})
	}
	ctx.JSON(http.StatusOK, "Success created order")

}
func HandlerGetOrder(ctx *gin.Context) {
	var orders []entity.Order
	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
