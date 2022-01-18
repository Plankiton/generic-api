FROM golang:1.17.6

ADD . /generic-api
WORKDIR /generic-api
RUN go mod tidy
RUN go build -o /bin/server .
