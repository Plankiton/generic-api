FROM golang:1.17.6

RUN go mod tidy -v
