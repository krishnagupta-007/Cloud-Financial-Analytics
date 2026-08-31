# Stage 1: Build binary using official Go alpine image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency definition files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Stage 2: Create lightweight runtime image
FROM alpine:latest

WORKDIR /app

# Copy executable from builder stage
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
