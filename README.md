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
