#!/bin/sh

git pull origin main

cd deploy

docker stop helloNode helloGolang helloNginx 

docker rm helloNginx helloNode helloGolang

docker rmi deploy_hello-node:latest deploy_hello-golang:latest deploy_hello-nginx:latest

docker compose up -d

docker ps