FROM golang:latest

MAINTAINER leo@scalingo.com

ADD . /go/src/github.com/Scalingo/acadock-monitoring
RUN go get github.com/tools/godep
RUN cd /go/src/github.com/Scalingo/acadock-monitoring/server && \
    godep go install && \
    cd /go/src/github.com/Scalingo/acadock-monitoring/runner/net && \
    godep go install

ENV RUNNER_DIR=/go/bin

CMD /go/bin/server

EXPOSE 4244
