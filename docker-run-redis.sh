#!/bin/bash

# Configuration
CONTAINER_NAME="weather-redis"
PORT="6379"

# Check if docker is installed
if ! command -v docker &> /dev/null
then
    echo "Error: docker could not be found. Please install docker first."
    exit 1
fi

# Stop and remove existing container if it exists
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    echo "Stopping and removing existing Redis container..."
    docker stop ${CONTAINER_NAME}
    docker rm ${CONTAINER_NAME}
fi

# Run Redis in Docker with port mapping
echo "Starting Redis in Docker on port ${PORT}..."
docker run --name ${CONTAINER_NAME} -p ${PORT}:6379 -d redis

# Verify it's running
if [ $? -eq 0 ]; then
    echo "Redis is now running in Docker!"
    echo "You can connect to it at localhost:${PORT}"
else
    echo "Failed to start Redis container."
fi
