package db

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticDB struct {
	Client *elasticsearch.Client
	Index  string
}

func NewElasticDB(uri, index string) *ElasticDB {
	// Proxy sozlamalari (agar sizda proxy bo'lmasa, proxy sozlamalarini olib tashlang)
	// proxy, err := url.Parse("")
	// if err != nil {
	// 	log.Fatalf("Proxy URL'ni tahlil qilishda xatolik: %v", err)
	// }

	// Elasticsearch sozlamalari
	cfg := elasticsearch.Config{
		Addresses: []string{uri}, // Elasticsearch server manzili
		// Username:  "",     // Foydalanuvchi nomi
		// Password:  "",    // Parol
		Transport: &http.Transport{
			// Proxy: http.ProxyURL(proxy),
			Proxy: nil, // Proxyni sozlash (agar sizda proxy bo'lmasa ushbu sozlashni olib tashlang)
		},
	}

	// Elasticsearch mijozini yaratish
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Elasticsearch client yaratishda xato: %v", err)
	}

	// Elasticsearch serveriga so‘rov yuborish
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Elasticsearch serveriga ulanishda xatolik: %v", err)
	}
	defer res.Body.Close()

	fmt.Printf("Elasticsearch server haqida ma'lumot: %v\n", res)

	return &ElasticDB{
		Client: client,
		Index:  index,
	}
}
