FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend_it ./cmd

FROM alpine:latest                      

WORKDIR /app/

COPY --from=builder /app/backend_it .

COPY --from=builder /app/assets ./assets

EXPOSE 8080

CMD ["./backend_it"]
