# Use the official Golang image as the base image
FROM golang:latest
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod file
COPY go.mod ./
# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Install ping utility
RUN apt-get update && apt-get install -y iputils-ping && rm -rf /var/lib/apt/lists/*

RUN go get github.com/lib/pq

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"]