# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Calvin Lee <clee421@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/aaclee/mkn-api

# Copy go mod and sum files
# COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o build/server/main cmd/server/main.go

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["build/server/main", "-dbhost=host.docker.internal"]