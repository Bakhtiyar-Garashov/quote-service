# Stage 1: Build the Go application
FROM golang:1.20 AS build

WORKDIR /app

# Copy the Go application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Stage 2: Create the final minimal image
FROM busybox:1.34

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .

# Expose any necessary ports
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]