package handler

import (
	"Product/model"
	"Product/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	qService service.ProductService
}

func NewProductHandler(qService service.ProductService) ProductHandler {
	return ProductHandler{qService: qService}
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	Products, err := h.qService.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Products})
}

func (h ProductHandler) GetProductsType(c *gin.Context) {
	genre := c.Param("Type")
	if genre == "A" || genre == "B" || genre == "C" || genre == "D" {
		Products, err := h.qService.GetProductsType(genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Products})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid types error"})
}

func (h ProductHandler) GetProductsCode(c *gin.Context) {
	code := c.Param("Code")
	Products, err := h.qService.GetProduct(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Products})
}

func (h ProductHandler) AddProduct(c *gin.Context) {
	var input model.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	Product, err := h.qService.AddProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": Product, "message": "Added"})
}

func (h ProductHandler) UpdateProduct(c *gin.Context) {
	code := c.Query("code")
	quantity := c.Query("val")
	value, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": model.ErrNotNumber})
		return
	}
	Product, err := h.qService.UpdateProduct(code, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Updated"})
}

func (h ProductHandler) SellProduct(c *gin.Context) {
	code := c.Query("code")
	quantity := c.Query("val")
	value, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": model.ErrNotNumber})
		return
	}
	Product, err := h.qService.SellProduct(code, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Updated"})
}

func (h ProductHandler) DeleteProduct(c *gin.Context) {
	Product, err := h.qService.DelistProduct(c.Param("Code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Deleted", "context": fmt.Sprintf("Product %v Deleted by Admin", Product.Code)})
}
