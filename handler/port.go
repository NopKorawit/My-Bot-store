package handler

import "github.com/gin-gonic/gin"

type GoodHandler interface {
	//web
	GetGoods(c *gin.Context)
	GetGoodsType(c *gin.Context)
	GetGoodsCode(c *gin.Context)
	AddGood(c *gin.Context)
	UpdateGood(c *gin.Context)
	DeleteGood(c *gin.Context)
	//line
	Callback(c *gin.Context)
	Hello(c *gin.Context)
}
