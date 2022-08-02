package handler

import (
	"fmt"
	"net/http"
	"store/model"
	"store/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type goodHandler struct {
	qService service.GoodService
}

func NewGoodHandler(qService service.GoodService) GoodHandler {
	return goodHandler{qService: qService}
}

func (h goodHandler) GetGoods(c *gin.Context) {
	goods, err := h.qService.GetGoods()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": goods})
}

func (h goodHandler) GetGoodsType(c *gin.Context) {
	genre := c.Param("Type")
	if genre == "A" || genre == "B" || genre == "C" || genre == "D" {
		goods, err := h.qService.GetGoodsType(genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": goods})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid types error"})
}

func (h goodHandler) GetGoodsCode(c *gin.Context) {
	code := c.Param("Code")
	goods, err := h.qService.GetGood(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": goods})
}

func (h goodHandler) AddGood(c *gin.Context) {
	var input model.StoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	good, err := h.qService.AddGood(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": good, "message": "Added"})
}

func (h goodHandler) UpdateGood(c *gin.Context) {
	code := c.Query("code")
	quantity := c.Query("val")
	value, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": model.ErrNotNumber})
		return
	}
	good, err := h.qService.UpdateGood(code, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": good, "message": "Updated"})
}

func (h goodHandler) DeleteGood(c *gin.Context) {
	good, err := h.qService.DelistGood(c.Param("Code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": good, "message": "Deleted", "context": fmt.Sprintf("Good %v Deleted by Admin", good.Code)})
}
