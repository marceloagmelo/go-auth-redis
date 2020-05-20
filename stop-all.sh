#!/usr/bin/env bash

source setenv.sh

# Mysql
echo "Finalizando o $MYSQL_HOSTNAME..."
docker rm -f $MYSQL_HOSTNAME

# Redis
echo "Finalizando o redis..."
docker rm -f redis

echo "Finalizando o $APP_NAME..."
docker rm -f $APP_NAME

# Remover rede
echo "Removendo a rede $DOCKER_NETWORK..."
docker network rm $DOCKER_NETWORK
