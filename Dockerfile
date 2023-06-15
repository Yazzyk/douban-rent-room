FROM golang:1.20.5-buster as builder
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

RUN mkdir -p /app

COPY . /app

RUN cd /app && make build_linux


FROM alpine:3.17
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

COPY --from=builder /app/build/douban_linux_amd64 /app

RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.4/main\nhttp://mirrors.aliyun.com/alpine/v3.4/community" >  /etc/apk/repositories \
    && apk update && apk add tzdata vim \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Shanghai/Asia" > /etc/timezone \
    && apk del tzdata

RUN chmod 777 /app/douban

CMD ["./douban"]
