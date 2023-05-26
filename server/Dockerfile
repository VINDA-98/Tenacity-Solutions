FROM golang:alpine as builder

WORKDIR /go/src/github.com/VINDA-98/Tenacity-Solutions/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/github.com/VINDA-98/Tenacity-Solutions/server

COPY --from=0 /go/src/github.com/VINDA-98/Tenacity-Solutions/server/server ./
COPY --from=0 /go/src/github.com/VINDA-98/Tenacity-Solutions/server/resource ./resource/
COPY --from=0 /go/src/github.com/VINDA-98/Tenacity-Solutions/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
