# Use the official Golang image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files
# COPY go.mod go.sum ./
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the app
RUN go build -o main .

# Expose port 8085
EXPOSE 8085

# Run the app
CMD ["./main"]