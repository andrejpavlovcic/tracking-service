# docker/Dockerfile

# Start from the official Go image
FROM golang:1.23.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=1 go build -o /client cmd/client/main.go

RUN ls -l /app

# Expose the application on port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/client"]