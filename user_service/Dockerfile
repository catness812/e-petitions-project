FROM golang:1.21-alpine3.17 as builder

COPY internal /app/internal/
COPY pkg /app/pkg/
COPY config.yml go.mod go.sum main.go /app/

RUN cd /app/ && \
    go mod download && \
    go build -o main

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/main /app/config.yml ./

EXPOSE 50052

CMD ["./main"]

#docker build -t e-petitions-user:1.0 .
#docker run --name user-service -it e-petitions-user:1.0