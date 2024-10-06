FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o tickets-data-api ./cmd/main.go


EXPOSE 8080

CMD ["./tickets-data-api"]