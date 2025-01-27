package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/handlers"
)

func SetupRouter(newsHandler *handlers.NewsHandler, checkHandler *handlers.CheckHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/check", checkHandler.Check)

	news := router.Group("/news")
	{
		news.POST("/", newsHandler.AddNews)
		news.GET("/", newsHandler.GetAllNews)
		news.GET("/search", newsHandler.SearchNews)
	}

	return router
}
