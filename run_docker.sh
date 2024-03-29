#!/bin/bash

# Define your image name
IMAGE_NAME="urlshortener"

# Build the Docker image
docker build -t $IMAGE_NAME .

# Run the Docker container
docker run -d -p 8080:8080 --name $IMAGE_NAME $IMAGE_NAME

echo "URL Shortener service is running on port 8080"

