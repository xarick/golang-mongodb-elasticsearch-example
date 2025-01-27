package main

import (
	"log"

	"github.com/xarick/golang-mongodb-elasticsearch-example/config"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/db"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/handlers"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/routes"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	mongo := db.NewMongoDB(cfg.MongoURL)
	elastic := db.NewElasticDB(cfg.ElasticURL, cfg.ElasticIndex)

	newsService := services.NewNewsService(mongo, elastic)
	newsHandler := handlers.NewNewsHandler(newsService)
	checkHandler := handlers.NewCheckHandler()

	r := routes.SetupRouter(newsHandler, checkHandler)

	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
