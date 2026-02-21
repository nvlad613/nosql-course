# Stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nosql-app ./cmd/app

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/nosql-app .
EXPOSE 8080
CMD ["./nosql-app"]