# Kafka Notify

Basic real-time notification system using Kafka in Go.

## API Documentation

See [swagger.json](./docs/swagger.json)

## Requirements

- Golang 18+
- Docker

## How to run

### Setup Kafka

- Install Docker

Docker is a software platform that allows developers to create, deploy, and run applications in containers. See the official documentation for install the version compatible with your OS: [Install Docker Engine](https://docs.docker.com/engine/install/)

- Run Docker compose

Inside root directory, run:

```docker
docker compose -f ./config/docker/docker-compose.yaml
```

## Execute the project

- Inside the root project directory, run:

```go
go run cmd/producer/producer.go
```

and

```go
go run cmd/consumer/consumer.go
```

This initializes the producer and consumer service respectively 

Application will be running locally on the port you configured! 🚀
