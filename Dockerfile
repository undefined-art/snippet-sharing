FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o snippet-sharing ./cmd/main.go

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/snippet-sharing .
COPY --from=builder /app/env ./env

EXPOSE 8080

CMD ["./snippet-sharing"]
