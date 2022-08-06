package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	//web
	GetProducts(c *gin.Context)
	GetProductsType(c *gin.Context)
	GetProductsCode(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	SellProduct(c *gin.Context)
	UpdateMultiProduct(c *gin.Context)
	//line
	Callback(c *gin.Context)
	Hello(c *gin.Context)
}
