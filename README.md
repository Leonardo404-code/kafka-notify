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

```shell
docker compose -f ./config/docker/docker-compose.yaml
```

Check if the Kafka container is running
```shell
docker ps
```

You should see all containers running, like:

![kafka containers](/docs/images/kafka-container-running.png)


### Download project dependencies

- Inside root project directory, run:

```go
go mod vendor
```

## Execute the project

- Inside the root project directory, run:

```go
go run cmd/producer/producer.go
```

You should see the log like:

![kafka containers](/docs/images/producer.png)

and in another terminal instance, run:

```go
go run cmd/consumer/consumer.go
```

You should see the log like:

![kafka containers](/docs/images/consumer.png)

This initializes the producer and consumer service respectively 

Application will be running locally on the port you configured! 🚀