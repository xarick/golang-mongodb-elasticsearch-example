package models

import "time"

type News struct {
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Author    string    `json:"author" bson:"author"`
	Source    string    `json:"source" bson:"source"`
	Published time.Time `json:"published" bson:"published"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
