FROM golang:1.4

COPY . /remy.io/src/
COPY ./web /remy.io/web

RUN cd /remy.io/src \
    && go build

EXPOSE 8080

ENTRYPOINT ["/remy.io/src/remy.io"]
