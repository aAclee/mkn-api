#!/usr/bin/env bash

chmod +x ./script/permissions.sh

INSTALLED=$true

GO_WHICH=$(which go)
if [[ -z $GO_WHICH ]]; then
  echo go installation not found
  echo follow these insructions: https://golang.org/doc/install
  INSTALLED=$false
else
  echo go installed already
fi

echo ---------------------------

BREW_WHICH=$(which brew)
if [[ -z $BREW_WHICH ]]; then
  echo brew installation not found
  echo follow these insructions: https://brew.sh
  INSTALLED=$false
else
  echo brew installed already
fi

echo ---------------------------

DOCKER_WHICH=$(which docker)
if [[ -z $DOCKER_WHICH ]]; then
  echo docker installation not found
  echo follow these insructions: https://docs.docker.com/docker-for-mac/install/
  INSTALLED=$false
else
  echo docker installed already
fi

echo ---------------------------

DOCKER_COMPOSE_WHICH=$(which docker-compose)
if [[ -z $DOCKER_COMPOSE_WHICH ]]; then
  echo docker-compose installation not found
  echo follow these insructions: https://docs.docker.com/compose/install/
  INSTALLED=$false
else
  echo docker-compose installed already
fi

echo ---------------------------

GO_MIGRATE_WHICH=$(which migrate)
if [[ -z $GO_MIGRATE_WHICH ]]; then
  echo golang-migrate installation not found
  echo follow these insructions: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
  INSTALLED=$false
else
  echo golang-migrate installed already
fi

echo ---------------------------

GO_DEP_WHICH=$(which dep)
if [[ -z $GO_DEP_WHICH ]]; then
  echo go dep installation not found
  echo follow these insructions: https://golang.github.io/dep/docs/installation.html
  INSTALLED=$false
else
  echo go dep installed already
fi

echo ---------------------------

if [ "$INSTALLED" = false ]; then
  echo Not everything was installed, exiting script
  exit 1
fi

echo Running setup...

docker-compose up -d psql_mkn &&
# Sleep 30 seconds to wait for images to be set up
sleep 10 &&
echo still here &&
sleep 10 &&
echo working on it &&
sleep 5 &&
echo few more seconds &&
migrate -database "postgres://mkn_psql:password@localhost:5432/mkn_db?sslmode=disable" -path ./db/migrations up &&
sleep 3 &&
go run cmd/seed/main.go

echo \"go run cmd/server/main.go\" to start server