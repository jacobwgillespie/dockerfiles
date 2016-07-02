FROM golang:1.6
MAINTAINER Jacob Gillespie <jacobwgillespie@gmail.com>

RUN \
  apt-get update -y && \
  apt-get install --no-install-recommends -y curl build-essential git bzr \
    ca-certificates mercurial && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

# Install
RUN \
  cd /tmp && \
  wget http://download.redis.io/redis-stable.tar.gz && \
  tar zxfv redis-stable.tar.gz && \
  cd redis-stable && \
  make && \
  make install && \
  cp -f src/redis-sentinel /usr/local/bin && \
  mkdir -p /etc/redis && \
  cp -f *.conf /etc/redis && \
  rm -rf /tmp/redis-stable* && \
  sed -i 's/^\(bind .*\)$/# \1/' /etc/redis/redis.conf && \
  sed -i 's/^\(daemonize .*\)$/# \1/' /etc/redis/redis.conf && \
  sed -i 's/^\(dir .*\)$/# \1\ndir \/data/' /etc/redis/redis.conf && \
  sed -i 's/^\(logfile .*\)$/# \1/' /etc/redis/redis.conf

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

VOLUME ["/data"]
ENTRYPOINT ["go-wrapper", "run"]
EXPOSE 6379 26379
