
FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o gateway-service ./cmd/main/main.go

FROM alpine:3.17
WORKDIR /app

COPY --from=builder /app/gateway-service /app/

EXPOSE 8080

ENTRYPOINT ["/app/gateway-service"]
