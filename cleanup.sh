#!/bin/sh

echo "Docker Cleanup Script (Project-Specific)"
echo "========================================="

echo "Performing full cleanup for the project..."

# Stop and remove project-specific containers
echo "Stopping and removing project-specific containers..."
docker container ls -a --filter "name=ascii-art-container"
docker container stop $(docker container ls -aq --filter "name=ascii-art-container")
docker container rm $(docker container ls -aq --filter "name=ascii-art-container")
echo "Project-specific containers have been stopped and removed."

# Remove project-specific images
echo "Removing project-specific images..."
docker image ls --filter "reference=ascii-art*"
docker image rm -f $(docker image ls -aq --filter "reference=ascii-art*")
echo "Project-specific images have been removed."

echo "Cleanup completed for the project."
