FROM golang:1.21.3-alpine3.18 as builder
WORKDIR /app

COPY internal ./internal
COPY pkg/ ./pkg/
COPY config.yml go.mod go.sum main.go ./

RUN go mod download && \
    go build -o main

FROM alpine:3.18
WORKDIR /app/petition_service

RUN apk add openssl && \
    adduser -D user && chown -R user /app

COPY --from=builder /app/main /app/config.yml ./

USER user

EXPOSE 50050

CMD ["./main"]