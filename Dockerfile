# Stage 1: The 'builder' stage to compile the Go application
FROM golang:1.24-alpine AS builder

# Install git for swag init
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Generate the Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the Go application into a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /itinerary-api ./main.go


# Stage 2: The 'final' stage to create the lightweight production image
FROM alpine:latest

# Install Chromium for PDF generation
RUN apk add --no-cache chromium udev ttf-freefont

# Set the working directory
WORKDIR /app

# Copy the required assets from the builder stage
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/docs ./docs

# Copy the compiled binary from the builder stage
COPY --from=builder /itinerary-api .

# Expose the port the application runs on
EXPOSE 8080

# The command to run the application
# We add the --no-sandbox flag which is required for running Chrome as a root user in a container
CMD ["/app/itinerary-api", "--no-sandbox"]