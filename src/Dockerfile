# Stage 1: Build stage
FROM golang:1.22.0-bullseye AS Build

# Set the working directory inside the container
WORKDIR /go/src

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download
    
# Copy the project files into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .

# Stage 2: Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the final stage
COPY --from=Build /bin/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/bin/app"]