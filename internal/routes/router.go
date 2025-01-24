package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/handlers"
)

func SetupRouter(newsHandler *handlers.NewsHandler, healthHandler *handlers.HealthHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", healthHandler.Check)

	news := router.Group("/news")
	{
		news.POST("/", newsHandler.AddNews)
		news.GET("/", newsHandler.GetAllNews)
		news.GET("/search", newsHandler.SearchNews)
	}

	return router
}
