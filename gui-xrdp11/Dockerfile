FROM jacobwgillespie/base

RUN \
  apt-get update -qq && \

  # Install window manager and x-server
  apt-get install -qy --no-install-recommends \
    libfuse2 \
    openbox \
    python-xdg \
    x11-xserver-utils \
    xbase-clients \
    xfonts-100dpi \
    xfonts-75dpi \
    xfonts-base \
    xkb-data \

    apt-utils \
    dialog \

    vnc4server \
    x11-xserver-utils \
    openbox \
    xfonts-base \
    xfonts-100dpi \
    xfonts-75dpi \
    libfuse2 \

    xrdp \

    default-jre \
    libossp-uuid-dev \
    libpng12-dev \
    libfreerdp-dev \
    libcairo2-dev \
    tomcat7 \
    nano

  # # Install TigerVNC
  # curl -L https://dl.bintray.com/tigervnc/stable/tigervnc-1.7.1.x86_64.tar.gz | tar -xz --strip=1 -C / && \
  #
  # # Install vncpasswd.py
  # mkdir -p /opt/vncpasswd && \
  # curl -L https://github.com/trinitronx/vncpasswd.py/archive/master.tar.gz | tar -zx --strip=1 -C /opt/vncpasswd && \

RUN \
  echo 'deb http://ppa.launchpad.net/webupd8team/java/ubuntu trusty main' >> /etc/apt/sources.list && \
  apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys EEA14886 && \

  # Install Dependencies
  apt-get update -qq && \
  /bin/bash -c 'echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections' && \
  apt-get install -y --force-yes oracle-java8-installer && \
  apt-get install -y --force-yes oracle-java8-set-default

ADD /vendor/ /

RUN \
  apt-get install -qy \
    x11-xserver-utils \
    libxrandr2 \
    libfuse2 \
    xutils \
    libxfixes3 \
    libx11-dev \
    libxml2 \
    zlib1g \
    fuse \
    autocutsel \
    pulseaudio \
    x11-apps && \

    # x11rdp install
    dpkg -i /tmp/x11rdp_0.9.0+devel-1_amd64.deb && \

    # xrdp needs to be installed seperately
    dpkg -i /tmp/xrdp_0.9.0+devel_amd64.deb

RUN \
  mkdir -p /var/cache/tomcat7 && \
  mkdir -p /var/lib/guacamole/classpath && \
  mkdir -p /usr/share/tomcat7/.guacamole && \
  mkdir -p /usr/share/tomcat7-root/.guacamole && \
  mkdir -p /root/.guacamole && \

  # Install guacd
  dpkg -i /tmp/guacamole-server_0.9.9_amd64.deb && \
  ldconfig && \

  # Configure tomcat
  cp /tmp/guacamole-0.9.9.war /var/lib/tomcat7/webapps/guacamole.war && \
  cp /tmp/guacamole-auth-noauth-0.9.9.jar /var/lib/guacamole/classpath && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat7/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat7-root/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /root/.guacamole/ && \

  # Fix tomcat webroot
  rm -Rf /var/lib/tomcat7/webapps/ROOT && \
  ln -s /var/lib/tomcat7/webapps/guacamole.war /var/lib/tomcat7/webapps/ROOT.war  && \

  ### Compensate for GUAC-513
  mkdir -p /usr/lib/x86_64-linux-gnu/freerdp && \
  ln -s /usr/local/lib/freerdp/guacsnd.so /usr/lib/x86_64-linux-gnu/freerdp/ && \
  ln -s /usr/local/lib/freerdp/guacdr.so /usr/lib/x86_64-linux-gnu/freerdp/

RUN \
  usermod -u 99 nobody && \
  usermod -g 100 nobody && \
  usermod -m -d /nobody nobody && \
  usermod -s /bin/bash nobody && \
  usermod -a -G adm,sudo nobody && \
  echo "nobody:PASSWD" | chpasswd && \

  mkdir /nobody && \
  mkdir -p /nobody/.config/openbox && \
  mkdir /nobody/.cache

RUN \
  # Cleanup
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/*

ADD /files/ /

EXPOSE 5900 6080 8080
