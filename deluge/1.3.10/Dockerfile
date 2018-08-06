FROM linuxserver/baseimage

RUN \
    apt-get update -q && \
    apt-get install -qy \
      python python-twisted python-openssl python-setuptools \
      intltool python-xdg python-chardet geoip-database python-libtorrent \
      python-notify python-pygame python-glade2 librsvg2-common \
      xdg-utils python-mako wget \
      && \
    wget -O deluge-1.3.10.tar.gz http://download.deluge-torrent.org/source/deluge-1.3.10.tar.gz && \
    tar zxfv deluge-1.3.10.tar.gz && \
    cd deluge-1.3.10 && \
    python setup.py build && \
    python setup.py install --install-layout=deb && \
    cd / && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /deluge-1.3.10.tar.gz /deluge-1.3.10

ADD services/ /etc/service/
RUN chmod -v +x /etc/service/*/run /etc/my_init.d/*.sh

VOLUME /config /downloads
EXPOSE 8112 58846 58946 58946/udp
