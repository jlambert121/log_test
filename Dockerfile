FROM golang:alpine

ADD . /go/src/github.com/jlambert121/log_test
RUN go install github.com/jlambert121/log_test

ENTRYPOINT /go/bin/log_test
