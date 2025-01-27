package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/db"
	"github.com/xarick/golang-mongodb-elasticsearch-example/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type NewsService struct {
	mongo   *db.MongoDB
	elastic *db.ElasticDB
}

func NewNewsService(mongo *db.MongoDB, elastic *db.ElasticDB) *NewsService {
	return &NewsService{mongo: mongo, elastic: elastic}
}

func (s *NewsService) AddNews(ctx *gin.Context, news models.News) error {
	news.CreatedAt = time.Now()

	// MongoDB'ga qo'shish
	if _, err := s.mongo.NewsCollection.InsertOne(ctx, news); err != nil {
		log.Printf("MongoDB xatosi: %v", err)
		return err
	}

	// Elasticsearch'ga qo'shish
	data, err := json.Marshal(news)
	if err != nil {
		log.Printf("JSON Marshal xatosi: %v", err)
		return err
	}

	res, err := s.elastic.Client.Index(s.elastic.Index, bytes.NewReader(data), s.elastic.Client.Index.WithContext(ctx))
	if err != nil {
		log.Printf("Elasticsearch xatosi: %v", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Elasticsearch xatosi: %s", res.Status())
		return fmt.Errorf("elasticsearch xato statusi: %s", res.Status())
	}

	return err
}

func (s *NewsService) GetAllNews(ctx *gin.Context) ([]models.News, error) {
	var news []models.News

	cursor, err := s.mongo.NewsCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &news); err != nil {
		log.Printf("MongoDB cursor.Decode xatosi: %v", err)
		return nil, err
	}

	// for cursor.Next(ctx) {
	// 	var item models.News
	// 	if err := cursor.Decode(&item); err != nil {
	// 		log.Printf("MongoDB cursor.Decode xatosi: %v", err)
	// 		continue
	// 	}
	// 	news = append(news, item)
	// }

	return news, nil
}

func (s *NewsService) SearchNews(query string) ([]models.News, error) {
	var results []models.News

	res, err := s.elastic.Client.Search(
		s.elastic.Client.Search.WithIndex(s.elastic.Index),
		s.elastic.Client.Search.WithQuery(query),
		s.elastic.Client.Search.WithSize(100), // Natijalarni sonini o'zgartirish
	)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch qidiruvida xato: %w", err)
	}
	defer res.Body.Close()

	var esRes struct {
		Hits struct {
			Hits []struct {
				Source models.News `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&esRes); err != nil {
		return nil, fmt.Errorf("elasticsearch natijalarini JSON dekodlashda xato: %w", err)
	}

	for _, hit := range esRes.Hits.Hits {
		results = append(results, hit.Source)
	}

	return results, nil
}
