FROM alpine:edge

RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
  apk update && apk upgrade \
  && apk add --no-cache mdocml-apropos build-base coreutils ca-certificates autoconf automake libtool \
  bc tree vim git fish dialog less tzdata ffmpeg-dev go

RUN sed -i "s/bin\/ash/usr\/bin\/fish/" /etc/passwd

RUN echo "set mouse-=a" >> ~/.vimrc

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV SHELL /usr/bin/fish

RUN mkdir -p /root/.config/fish && \
  echo "set -gx GOPATH /app" >> /root/.config/fish/config.fish && \
  echo "set -gx PATH {\$PATH} /app/bin" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE on" >> /root/.config/fish/config.fish

WORKDIR /app/src/golang/miniav/goff

VOLUME /app

CMD ["/usr/bin/fish"]
