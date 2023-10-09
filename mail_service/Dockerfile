FROM golang:1.21-alpine

WORKDIR /mail_service

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -v -o app ./main.go

CMD ["./app"]