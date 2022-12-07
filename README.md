# url-shortner

## Development

1. clone this repository
2. start up application
    - execute: `docker-compose up`


#### Quick Docker Reference

```bash

# shutdown stack and keep volumes
docker-compose down

# force rebuild and start all services in background
docker-compose up --build -d

# view and tail logs of all services
docker-compose logs -f

```

This Microservice is responsible for creating short url for an input url
The file UrlDatabase.csv is used as Database for storing urls
