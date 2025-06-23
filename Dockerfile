# Cache gogcc alpine
FROM golang:1.23-alpine AS gogcc
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev

# Build the binary
FROM gogcc AS builder

WORKDIR /app

# Download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

# Build /app/bin
COPY main.go .
# Make a .env file
RUN echo "DB_PATH=${DB_PATH}" > .env
RUN echo "DB_AUTH_TOKEN=${DB_AUTH_TOKEN}" >> .env

RUN go build -ldflags="-s -w" -o bin .

# Serve the binary with pb_public
FROM alpine:latest AS bin

# Build args for environment variables
ARG DB_PATH
ARG DB_AUTH_TOKEN

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev \
        ca-certificates

WORKDIR /app/

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=builder /app/bin .

# Create pb_data directory for PocketBase data
RUN mkdir -p pb_data && \
    chown -R appuser:appgroup /app

# Set environment variables from build args
ENV DB_PATH=${DB_PATH}
ENV DB_AUTH_TOKEN=${DB_AUTH_TOKEN}

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["/app/bin", "serve", "--http=0.0.0.0:8080"]
