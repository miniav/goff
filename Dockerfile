FROM alpine:3.11

RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk update && apk upgrade \
  && apk add --no-cache mdocml-apropos build-base coreutils ca-certificates \
  tree vim git fish less tzdata ffmpeg-dev go

RUN sed -i "s/bin\/ash/usr\/bin\/fish/" /etc/passwd

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV SHELL /usr/bin/fish

RUN mkdir -p /root/.config/fish && \
  echo "set -gx GOPATH /app" >> /root/.config/fish/config.fish && \
  echo "set -gx PATH {\$PATH} /app/bin" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE on" >> /root/.config/fish/config.fish && \
  echo "set -gx GOPROXY https://goproxy.cn,direct" >> /root/.config/fish/config.fish

WORKDIR /app/src/github.com/miniav/goff

VOLUME /app/src/github.com/miniav/goff

CMD "/usr/bin/fish"
