# Use a specific version of the golang base image for predictable builds
FROM golang:1.18.0-buster AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files first to leverage Docker cache layers, as dependencies change less frequently than code
COPY go.mod go.sum ./

# Download dependencies for caching purposes
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app as a static binary so it doesn't require C libraries at runtime
# This improves security and reduces the container size
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a specific version of the Alpine image for runtime to have a stable and secure base
FROM alpine:3.13

# Add ca-certificates and tzdata for SSL connections and time zone support
RUN apk --no-cache add ca-certificates tzdata

# Set the working directory in the container to /root
# This is where we'll run our app from
WORKDIR /root/

# Copy the statically-linked binary from the builder stage
COPY --from=builder /app/main .

# Inform Docker that the container is listening on the specified port at runtime.
# Though this is optional, it's a good practice.
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
