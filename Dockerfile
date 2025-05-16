# Use an official Go image as the base image
FROM docker.io/library/golang:1.24 AS builder

# Add metadata to the image
LABEL maintainer="cgkaldan, sgougoul, saggelak"
LABEL version="1.0"
LABEL description="Dockerized ASCII Art Generator server"

# Set the working directory inside the container
WORKDIR /app

# Copy project files into the container
COPY . .

# Build the Go server
RUN go build -o ascii-art-generator main.go

# Use a minimal base image to run the application
FROM alpine:latest

# Add metadata to the final image
LABEL maintainer="cgkaldan, sgougoul, saggelak "
LABEL version="1.0"
LABEL description="Minimal runtime container for ASCII Art Generator server"

# Install glibc compatibility layer
RUN apk add --no-cache libc6-compat

# Set the working directory for the runtime container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app /app

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["/app/ascii-art-generator"]
