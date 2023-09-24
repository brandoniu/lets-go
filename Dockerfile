# Use the official Golang image to create a build artifact.
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

###

# Start from a lightweight image.
FROM alpine:latest

# Copy the binary from builder.
COPY --from=builder /app/main /app/

# Port on which the service will be exposed.
EXPOSE 8080

# Command to run the application.
CMD ["/app/main"]

# Setting up the environment your application needs to run.
# Defining the build process for your application.
# Creating the final Docker image with the application's binary.
