# MunchKin API

# Requirements

## Folder Structure
For the packages to import properly ensure that the folder structure looks like below:

```
$GOPATH/src/github.com/aaclee
└── mkn-api/
    ├── cmd/
    ├── config/
    ├── db/
    ├── pkg/
    └── script/
```

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

### Using the `setup.sh`

Using the `setup.sh` script will check if all the dependancies are availible.

```shell
# At the top level folder

# Give permission to setup.sh
$ chmod +x ./script/setup.sh

# Run setup script
$ script/setup.sh

# Running the server
$ go run cmd/server/main.go
```

### Manually

```shell
# Starting the database
$ docker-compose up

# Running migrations
$ migrate -database "postgres://mkn_psql:password@localhost:5432/mkn_db?sslmode=disable" -path ./db/migrations up

# Seed database
$ go run cmd/seed/main.go

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

### Connecting Manually

```shell
$ psql -h localhost -p 5432 -d mkn_db -U mkn_psql
```