FROM golang:1.20.5-buster as builder
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

RUN mkdir -p /app

COPY . /app

RUN cd /app && make build_linux


FROM debian:12.0-slim
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

COPY --from=builder /app/build/douban_linux_amd64 /app

RUN sudo timedatectl set-timezone Asia/Shanghai

RUN chmod 777 /app/douban

CMD ["./douban"]
