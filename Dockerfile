FROM golang:1.4

COPY . /remy.io
COPY ./web /remy.io/web
COPY ./static /remy.io/web

RUN cd /remy.io \
    && go build

ENV DOCKER 1

EXPOSE 8080

ENTRYPOINT ["/remy.io/remy.io"]
