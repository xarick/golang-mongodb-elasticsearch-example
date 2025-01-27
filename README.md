# golang-mongodb-elasticsearch-example

- Running Docker Compose (docker-compose.yml is in the project)
```bash
docker-compose up -d
```

- Connecting to the MongoDB server (2 different methods)
```bash
docker exec -it mongodb bash
docker exec -it mongodb mongosh -u admin -p admin123 --authenticationDatabase admin
```

-  Elasticsearchda indeksni qo'lda qo'shish uchun
```bash
curl -X PUT "http://localhost:9200/news" -H 'Content-Type: application/json' -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "properties": {
      "title": { "type": "text" },
      "content": { "type": "text" },
      "author": { "type": "keyword" },
      "source": { "type": "keyword" },
      "published": { "type": "date" },
      "created_at": { "type": "date" }
    }
  }
}'
```

- Elasticsearch buyruqlari
```bash
curl -X GET "http://localhost:9200/_cat/indices?v"          # Indekslar ro'yxatini ko'rish
curl -X GET "http://localhost:9200/news"                    # news indeksini ko'rish
curl -X GET "http://localhost:9200/news/_mapping?pretty"    # Maydon turlarini ko'rish
curl -X GET "http://localhost:9200/news/_count"             # new indeksida hujjat sonini ko'rish
```

- news indeksida birinchi 10 ta hujjatni ko'rish
```bash
curl -X GET "http://localhost:9200/news/_search?pretty" -H 'Content-Type: application/json' -d '{
  "size": 10,
  "query": {
    "match_all": {}
  }
}'
```
