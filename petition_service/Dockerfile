FROM golang:1.21-alpine

WORKDIR /app/petition_service

COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY config.yml ./
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go mod download && \
    go build -o main

EXPOSE 50050

CMD ["./main"]