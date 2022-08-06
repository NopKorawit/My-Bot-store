package main

import (
	"Product/handler"
	"Product/repository"
	"Product/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	handler.InitAll()
	//connect to database + auto migrate
	db := handler.ConnectDatabase()

	storeRepo := repository.NewProductRepositoryDB(db)
	storeService := service.NewProductService(storeRepo)
	storeHandler := handler.NewProductHandler(storeService)

	route := gin.Default()
	route.Use(cors.Default())
	//Routes
	web := route.Group("/api/v1/web")

	{
		web.GET("/", storeHandler.Hello)
		web.GET("/All", storeHandler.GetProducts)
		web.GET("/:Type", storeHandler.GetProductsType)
		web.GET("/code/:Code", storeHandler.GetProductsCode)
		web.POST("/", storeHandler.AddProduct)
		web.PUT("/update", storeHandler.UpdateProduct)
		web.PUT("/sell", storeHandler.SellProduct)
		web.DELETE("/:Code", storeHandler.DeleteProduct)
	}

	line := route.Group("/api/v1/line")
	{
		line.GET("/", storeHandler.Hello)
		line.POST("/callback", storeHandler.Callback)
	}
	//Run Server
	route.Run()
}
