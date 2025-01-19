FROM golang:1.23.4 AS base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download Go dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Install air for hot-reloading during development
RUN go install github.com/air-verse/air@latest

# Copy the source code into the container
COPY . .

# Start the application with air for hot reloading
CMD ["air"]