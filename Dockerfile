FROM    golang:1.12 as bulider

MAINTAINER  Gaotiefeng "1096392101@qq.com"

ENV GOPROXY https://goproxy.io
ENV GO111MODULE=on

WORKDIR $GOPATH/src/go-gin

ADD . .
ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . $GOPATH/src/go-gin

EXPOSE  8080

ENTRYPOINT go run app/main.go