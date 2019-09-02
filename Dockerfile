FROM golang

WORKDIR /go/src

RUN export GOROOT=$HOME/go
RUN export PATH=$PATH:$GOROOT/bin

COPY ./src/ /go/src/

