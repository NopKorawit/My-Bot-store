package handler

import (
	"Product/model"
	"Product/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h productHandler) GetProducts(c *gin.Context) {
	Products, err := h.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Products})
}

func (h productHandler) GetProductsType(c *gin.Context) {
	genre := c.Param("Type")
	if genre == "A" || genre == "B" || genre == "C" || genre == "D" {
		Products, err := h.productService.GetProductsType(genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Products})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid types error"})
}

func (h productHandler) GetProductsCode(c *gin.Context) {
	code := c.Param("Code")
	Products, err := h.productService.GetProduct(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Products})
}

func (h productHandler) AddProduct(c *gin.Context) {
	var input model.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	Product, err := h.productService.AddProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": Product, "message": "Added"})
}

func (h productHandler) UpdateMultiProduct(c *gin.Context) {
	var input []model.MultiProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	Product, err := h.productService.UpdateMultiProducts(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": Product, "message": "Added"})
}

func (h productHandler) UpdateProduct(c *gin.Context) {
	code := c.Query("code")
	quantity := c.Query("val")
	value, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": model.ErrNotNumber})
		return
	}
	Product, err := h.productService.UpdateProduct(code, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Updated"})
}

func (h productHandler) SellProduct(c *gin.Context) {
	code := c.Query("code")
	quantity := c.Query("val")
	value, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": model.ErrNotNumber})
		return
	}
	Product, err := h.productService.SellProduct(code, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Updated"})
}

func (h productHandler) DeleteProduct(c *gin.Context) {
	Product, err := h.productService.DelistProduct(c.Param("Code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Product, "message": "Deleted", "context": fmt.Sprintf("Product %v Deleted by Admin", Product.Code)})
}
