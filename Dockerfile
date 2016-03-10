FROM golang:1.6-alpine

RUN mkdir /opt; mkdir /opt/service

WORKDIR /opt/service

ADD main /opt/service

ENTRYPOINT ["./main"]
