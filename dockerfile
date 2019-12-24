# 构建
# MAINTAINER zs "810909753@qq.com"
FROM golang:1.12.1 as builder

WORKDIR /app
ENV GOPROXY   https://gocenter.io

ADD . /app
RUN go build  -mod=vendor -o /app/memo  main.go

# 运行
FROM haoyinqianzui/ubuntu:1.0
WORKDIR /app
ENV GO_SSO_WORKDIR  /app
RUN apt-get update
RUN apt-get install libc6
COPY --from=builder /app/memo /app

EXPOSE 8009
ENTRYPOINT ["./memo"]