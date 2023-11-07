FROM golang:1.21.3-bookworm as builder
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

RUN mkdir -p /app

COPY . /app

RUN cd /app && make build_linux


FROM ubuntu:20.04
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

COPY --from=builder /app/build/douban_linux_amd64 /app

# 国内源
#RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && sed -i 's|security.debian.org|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list && apt update
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list && apt update

RUN apt install -y apt-transport-https ca-certificates vim tzdata

# 修改时区为中国
ENV TZ=Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN chmod 777 /app/douban

CMD ["./douban"]
