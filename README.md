# MunchKin API

# Requirements

## Go

### Link

https://golang.org/doc/install

## Brew

### Link

https://brew.sh

## golang-migrate

### Link

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


### Installation

```shell
# macOS

$ brew install golang-migrate
```

## DEP

### Link

https://golang.github.io/dep/docs/installation.html

### Installation

```shell
# macOS

$ brew install dep
$ brew upgrade dep
```

User `dep` to install vendor files

## Docker

### Link

https://docs.docker.com/docker-for-mac/install/

## Docker Compose

### Link

https://docs.docker.com/compose/install/


## Getting Started


```shell
# Starting the database
$ docker-compose up
```

```shell
# Running migrations
$ migrate -database "postgres://mkn_psql:password@localhost:5432/mkn_db?sslmode=disable" -path ./db/migrations up
```

```shell
# Running the server
$ go run cmd/server/main.go
```

## Tests

```shell
$ go test ./...
```

## Migrations 

```shell
# Creating a migration
$ migrate create -ext sql -dir db/migrations create_sample_table
```