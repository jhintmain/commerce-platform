FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/worker ./cmd/worker
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/scheduler ./cmd/scheduler

####  duilder 배포
FROM ubuntu:24.04

WORKDIR /app

RUN apt-get update \
    && apt-get install -y ca-certificates tzdata curl \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin /app/bin

EXPOSE 8080

CMD ["/app/bin/api"]