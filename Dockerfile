FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git tzdata
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /sanitizer ./cmd/server

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /
COPY --from=builder /sanitizer /sanitizer
EXPOSE 8080
ENV PORT=8080
ENV LOG_LEVEL=info
ENTRYPOINT ["/sanitizer"]
