FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install reflex
RUN go install github.com/cespare/reflex@latest

# Copy the source code
COPY . .

# Expose port
EXPOSE 4200

# Default command will be overridden by docker-compose
CMD ["go", "run", "/app/cmd/server/main.go"]