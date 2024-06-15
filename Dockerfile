FROM golang:1.22.3

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

RUN air ./cmd/main.go -b 0.0.0.0