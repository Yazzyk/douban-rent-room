FROM alpine:3.17
WORKDIR /app

COPY ./build/douban-rent-room_linux_amd64/ /app

RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.4/main\nhttp://mirrors.aliyun.com/alpine/v3.4/community" >  /etc/apk/repositories \
    && apk update && apk add tzdata vim \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Shanghai/Asia" > /etc/timezone \
    && apk del tzdata

RUN chmod 777 /app/douban-rent-room

CMD ["./douban-rent-room"]
