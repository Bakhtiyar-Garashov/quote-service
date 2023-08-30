FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]