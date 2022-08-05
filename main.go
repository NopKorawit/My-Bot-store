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

	queueRepo := repository.NewProductRepositoryDB(db)
	queueService := service.NewProductService(queueRepo)
	queueHandler := handler.NewProductHandler(queueService)

	route := gin.Default()
	route.Use(cors.Default())
	//Routes
	web := route.Group("/api/v1/web")

	{
		web.GET("/", queueHandler.Hello)
		web.GET("/All", queueHandler.GetProducts)
		web.GET("/:Type", queueHandler.GetProductsType)
		web.GET("/code/:Code", queueHandler.GetProductsCode)
		web.POST("/", queueHandler.AddProduct)
		web.PUT("/update", queueHandler.UpdateProduct)
		web.PUT("/sell", queueHandler.SellProduct)
		web.DELETE("/:Code", queueHandler.DeleteProduct)
	}

	line := route.Group("/api/v1/line")
	{
		line.GET("/", queueHandler.Hello)
		line.POST("/callback", queueHandler.Callback)
	}
	//Run Server
	route.Run()
}
