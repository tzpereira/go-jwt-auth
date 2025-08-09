# Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 9001
CMD ["./app"]