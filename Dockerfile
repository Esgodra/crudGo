FROM golang:latest AS builder

RUN #!/bin/bash
RUN apt-get install git
RUN go get -d github.com/ibmdb/go_ibm_db

RUN cd /home/../go/src/github.com/ibmdb/go_ibm_db/installer && go run setup.go

RUN export DB2HOME=/home/../go/src/github.com/ibmdb/go_ibm_db/installer/clidriver
RUN export CGO_CFLAGS=-I$DB2HOME/include
RUN export CGO_LDFLAGS=-L$DB2HOME/lib
RUN export LD_LIBRARY_PATH=/home/../go/src/github.com/ibmdb/go_ibm_db/installer/clidriver/lib

RUN cd /home && cd ../
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .
CMD ["/app/main"]