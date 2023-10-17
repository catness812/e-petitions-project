FROM golang:1.21.1-alpine3.17 as builder

WORKDIR /app
COPY internal ./internal
COPY pkg/ ./pkg/
COPY config.yml go.mod go.sum main.go ./
RUN go mod download && \
    go build -o main

FROM alpine:3.17
WORKDIR /app/petition_service
COPY --from=builder /app/main /app/config.yml ./
EXPOSE 50050

CMD ["./main"]


#docker build -t e-petitions-petition:1.0 .