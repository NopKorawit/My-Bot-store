package main

import (
	"store/handler"
	"store/repository"
	"store/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	handler.InitAll()
	//connect to database + auto migrate
	db := handler.ConnectDatabase()

	queueRepo := repository.NewGoodRepositoryDB(db)
	queueService := service.NewGoodService(queueRepo)
	queueHandler := handler.NewGoodHandler(queueService)

	route := gin.Default()
	route.Use(cors.Default())
	//Routes
	web := route.Group("/api/v1/web")

	{
		web.GET("/", queueHandler.Hello)
		web.GET("/All", queueHandler.GetGoods)
		web.GET("/:Type", queueHandler.GetGoodsType)
		web.GET("/code/:Code", queueHandler.GetGoodsCode)
		web.POST("/", queueHandler.AddGood)
		web.PUT("/", queueHandler.UpdateGood)
		web.DELETE("/:Code", queueHandler.DeleteGood)
	}

	line := route.Group("/api/v1/line")
	{
		line.GET("/", queueHandler.Hello)
		line.POST("/callback", queueHandler.Callback)
	}
	//Run Server
	route.Run()
}
