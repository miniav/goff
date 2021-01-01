FROM golang:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache build-base tzdata ffmpeg-dev

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /app/src/github.com/miniav/goff

CMD "/bin/sh"
