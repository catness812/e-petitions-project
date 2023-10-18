FROM golang:1.21.1-alpine3.17 as builder

WORKDIR /app

COPY internal ./internal
COPY model ./model
COPY config.yml go.mod go.sum main.go rbac.yml ./

RUN go mod download && \
    go build -o main

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/main /app/config.yml /app/rbac.yml ./

EXPOSE 1337

CMD ["./main"]

#docker build -t e-petitions-gateway:1.0 .
#docker run --name gateway -it e-petitions-gateway
