FROM buildpack-deps:stable-scm

ARG VERSION=1.12

WORKDIR /root

COPY go${VERSION}.linux-amd64.tar.gz /root

RUN tar -C /usr/local/ -zxvf /root/go${VERSION}.linux-amd64.tar.gz && \
  rm -f /root/go${VERSION}.linux-amd64.tar.gz

RUN sed -i 's/deb.debian.org/mirrors.163.com/g' /etc/apt/sources.list && \
  sed -i 's/security.debian.org/mirrors.163.com/g' /etc/apt/sources.list

RUN apt-get update && apt-get upgrade -y && \
  apt-get install -y --no-install-recommends \
  tree vim bash-completion apt-utils fish \
  build-essential dialog less cmake locales nasm pkg-config && \
  apt-get clean

RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && \
  echo "zh_CN.UTF-8 UTF-8" >> /etc/locale.gen && locale-gen

RUN sed -i "s/bin\/bash/usr\/bin\/fish/" /etc/passwd

RUN mkdir -p /root/.config/fish && \
  echo "set -gx GOPATH /root/gocode" >> /root/.config/fish/config.fish && \
  echo "set -gx PATH {\$PATH} /usr/local/go/bin /root/gocode/bin" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE on" >> /root/.config/fish/config.fish

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p /root/gocode/bin

VOLUME /root/gocode

CMD ["/usr/bin/fish"]
