#!/bin/bash

# Set variables for image and container names
IMAGE_NAME="ascii-art-generator"
IMAGE_TAG="1.0"
CONTAINER_NAME="ascii-art-container"
DOCKERFILE="Dockerfile"
PORT=8080

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker image build -f $DOCKERFILE -t $IMAGE_NAME:$IMAGE_TAG .
if [ $? -ne 0 ]; then
    echo "Failed to build the Docker image. Exiting."
    exit 1
fi

# Step 2: Stop and remove any existing container with the same name
if docker ps -a --format '{{.Names}}' | grep -Eq "^${CONTAINER_NAME}\$"; then
    echo "Stopping and removing existing container..."
    docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME
fi

# Step 3: Run the Docker container
echo "Running the Docker container..."
docker run -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME:$IMAGE_TAG
if [ $? -ne 0 ]; then
    echo "Failed to start the Docker container. Exiting."
    exit 1
fi

# Step 4: Confirm the container is running
if docker ps --format '{{.Names}}' | grep -Eq "^${CONTAINER_NAME}\$"; then
    echo "Container '${CONTAINER_NAME}' is up and running."
    echo "Access the application at: http://localhost:$PORT"
else
    echo "Container failed to start. Check the logs for more details."
    exit 1
fi
