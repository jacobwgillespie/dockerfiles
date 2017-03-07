FROM debian:8

ENV DEBIAN_FRONTEND noninteractive
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV I2P_DIR /usr/share/i2p

RUN \
  # Install dependencies
  apt-get update -qq && \
  apt-get upgrade -qy && \
  apt-get install -qy wget apt-transport-https locales && \

  # Configure locale
  sed -i 's/.*\(en_US\.UTF-8\)/\1/' /etc/locale.gen && \
  /usr/sbin/locale-gen && \
  /usr/sbin/update-locale LANG=${LANG} LANGUAGE=${LANGUAGE} && \

  # Install i2p
  wget -q -O - https://geti2p.net/_static/i2p-debian-repo.key.asc | apt-key add - && \
  echo "deb https://deb.i2p2.de/ jessie main" | tee -a /etc/apt/sources.list.d/i2p.list && \
  echo "deb-src https://deb.i2p2.de/ jessie main" | tee -a /etc/apt/sources.list.d/i2p.list && \
  apt-get update -qq && \
  apt-get install -qy i2p i2p-keyring && \

  # Configure i2p
  sed -i 's/127\.0\.0\.1/0.0.0.0/g' ${I2P_DIR}/i2ptunnel.config && \
  sed -i 's/::1,127\.0\.0\.1/0.0.0.0/g' ${I2P_DIR}/clients.config && \
  echo "i2cp.tcp.bindAllInterfaces=true" >> ${I2P_DIR}/router.config && \
  echo "i2np.ipv4.firewalled=true" >> ${I2P_DIR}/router.config && \
  echo "i2np.ntcp.ipv6=false" >> ${I2P_DIR}/router.config && \
  echo "i2np.udp.ipv6=false" >> ${I2P_DIR}/router.config && \
  echo "i2np.upnp.enable=false" >> ${I2P_DIR}/router.config && \

  # Cleanup
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/* /root/.cache/*

ADD run-i2p /usr/local/bin/run-i2p

##
# Expose some ports used by I2P
# Description at https://geti2p.net/ports
#
# Main ports:
# 2827 - BOB port
# 4444 — HTTP proxy
# 6668 — Proxy to Irc2P
# 7656 - SAM port
# 7657 — router console
# 7658 — self-hosted eepsite
# 7659 — SMTP proxy to smtp.postman.i2p
# 7660 — POP3 proxy to pop.postman.i2p
# 8998 — Proxy to mtn.i2p-projekt.i2p
##
EXPOSE 2827 7650 7654 7655 7656 7657 7658 7659 7660 7661 7662 4444 6668 8998

VOLUME /var/lib/i2p
ENTRYPOINT ["/usr/local/bin/run-i2p"]
CMD ["console"]
