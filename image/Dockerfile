FROM playlist/base
MAINTAINER Jacob Gillespie <jacob@playlist.com>

ENV HOME /root

# Add redis repo
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 5862E31D
RUN echo deb http://ppa.launchpad.net/rwky/redis/ubuntu trusty main > /etc/apt/sources.list.d/redis.list

# Install redis
RUN apt-get install -y redis-server libhiredis-dev
RUN mkdir /etc/service/redis
ADD support/runit/redis /etc/service/redis/run
ADD support/config/redis.conf /etc/redis/redis.conf

# Clean up
RUN apt-get clean
RUN rm -rf /tmp/* /var/tmp/*

# Use playlist/base's init script
CMD ["/sbin/my_init"]
