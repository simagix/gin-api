FROM golang:1.7.3

ENV GOPATH=/go
ADD . /go/src/github.com/simagix/gin-api
WORKDIR /go/src/github.com/simagix/gin-api
RUN go get gopkg.in/mgo.v2 gopkg.in/gin-gonic/gin.v1 github.com/bitly/go-simplejson \
    && go build -o /gin-api github.com/simagix/gin-api \
    && rm -rf /go/src/github.com/simagix/gin-api

WORKDIR /
ENTRYPOINT ["/gin-api"]
