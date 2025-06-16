FROM golang:1.24-alpine

RUN apk update && apk add --no-cache sqlite sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o snippet-sharing ./cmd/main.go

EXPOSE 8080

CMD ["./snippet-sharing"]
