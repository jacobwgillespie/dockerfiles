FROM phusion/baseimage:0.9.19
MAINTAINER Jacob Gillespie <jacobwgillespie@gmail.com>

ENV \
  USER_ID="0" \
  GROUP_ID="0" \
  TERM="xterm" \
  QBITTORRENT_VERSION="3.3.10"

CMD ["/sbin/my_init"]

ADD ./files /files
RUN sync && /bin/bash /files/tmp/install.sh

EXPOSE 5900
