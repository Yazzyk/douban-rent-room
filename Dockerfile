FROM golang:1.20.5-buster as builder
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

RUN mkdir -p /app

COPY . /app

RUN cd /app && make build_linux


FROM unstable-slim
LABEL MAINTAINER='yazzyk<root@shroot.dev>'
WORKDIR /app

COPY --from=builder /app/build/douban_linux_amd64 /app

RUN apt update && apt install tzdata vim \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Shanghai/Asia" > /etc/timezone \
    && apk del tzdata

RUN chmod 777 /app/douban

CMD ["./douban"]
