FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY products.json .

EXPOSE 8080

CMD [ "./server" ]
