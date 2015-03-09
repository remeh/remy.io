FROM golang:1.4

COPY . /go/src/github.com/remeh/remy.io

RUN cd /go/src/github.com/remeh/remy.io \
    && go build

EXPOSE 8080

ENTRYPOINT ["/go/src/github.com/remeh/remy.io/remy.io"]
