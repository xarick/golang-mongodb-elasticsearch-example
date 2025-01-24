package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client         *mongo.Client
	NewsCollection *mongo.Collection
}

func NewMongoDB(uri string) *MongoDB {
	// MongoDB ulanishi uchun sozlamalarni yaratish
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB mijozini ulash
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB ulanishda xato: %v", err)
	}

	// Ulash muvaffaqiyatli ekanligini tekshirish
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB serveriga ping jo'natishda xato: %v", err)
	}

	db := client.Database("newsDB")
	return &MongoDB{
		Client:         client,
		NewsCollection: db.Collection("news"),
	}
}
