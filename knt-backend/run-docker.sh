#!/bin/sh
cd "${0%/*}"

# Build docker image
docker build -t docker-knt-backend:multistage -f Dockerfile.multistage .

# Stop and remove docker container
docker stop knt-backend
docker rm knt-backend

# Run docker image
docker run -p 5000:5000 --name knt-backend -v "$(pwd)"/database:/database -v "$(pwd)"/logs:/logs docker-knt-backend:multistage
