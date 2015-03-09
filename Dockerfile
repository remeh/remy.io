FROM golang:1.4

COPY . /remy.io
COPY ./web /remy.io/web

RUN cd /remy.io \
    && go build

EXPOSE 8080

ENTRYPOINT ["/remy.io/remy.io"]
