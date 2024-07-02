#!/bin/bash

docker run -d --name golang_class \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=Secret \
    -e POSTGRES_DB=golang_class \
    -p 3333:5432 \
    bitnami/postgresql:latest