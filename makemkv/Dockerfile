FROM jacobwgillespie/gui

ADD /files/ /

RUN \
  /tmp/install.sh && \
  rm /tmp/install.sh

RUN \
  chown -R abc:abc /openbox
