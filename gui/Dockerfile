FROM jacobwgillespie/base

RUN \
  # Install dependencies
  echo "deb http://ppa.launchpad.net/webupd8team/java/ubuntu xenial main" | tee /etc/apt/sources.list.d/webupd8team-java.list && \
  echo "deb-src http://ppa.launchpad.net/webupd8team/java/ubuntu xenial main" | tee -a /etc/apt/sources.list.d/webupd8team-java.list && \
  apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys EEA14886 && \
  apt-get update -qq && \

  # Install window manager, X server, and Xrdp
  apt-get install -qy --no-install-recommends \
    libfuse2 \
    openbox \
    vnc4server \
    x11-xserver-utils \
    xfonts-100dpi \
    xfonts-75dpi \
    xfonts-base \
    xrdp && \

  # Install Java
  echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | debconf-set-selections && \
  apt-get install -qy oracle-java8-installer && \

  # Install Guacamole (web UI)
  apt-get install -qy --no-install-recommends \
    build-essential \
    libcairo2-dev \
    libfreerdp-dev \
    libjpeg62-turbo-dev \
    libossp-uuid-dev \
    libpng12-dev \
    libwebp-dev \
    tomcat8 && \
  wget -O /tmp/guacd.tar.gz "http://apache.org/dyn/closer.cgi?action=download&filename=incubator/guacamole/0.9.11-incubating/source/guacamole-server-0.9.11-incubating.tar.gz" && \
  cd /tmp && \
  tar zxfv /tmp/guacd.tar.gz && \
  cd guacamole-server-* && \
  ./configure && \
  make && \
  make install && \
  ldconfig && \
  mkdir -p /var/cache/tomcat8 && \
  mkdir -p /var/lib/guacamole/classpath && \
  mkdir -p /usr/share/tomcat8/.guacamole && \
  mkdir -p /usr/share/tomcat8-root/.guacamole && \
  mkdir -p /root/.guacamole && \
  wget -O /var/lib/tomcat8/webapps/guacamole.war "http://apache.org/dyn/closer.cgi?action=download&filename=incubator/guacamole/0.9.11-incubating/binary/guacamole-0.9.11-incubating.war" && \
  wget -O /var/lib/guacamole/classpath/guacamole-auth-noauth-0.9.11.jar "http://apache.org/dyn/closer.cgi?action=download&filename=incubator/guacamole/0.9.11-incubating/binary/guacamole-auth-noauth-0.9.11-incubating.tar.gz" && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat8/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat8-root/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /root/.guacamole/ && \
  rm -Rf /var/lib/tomcat8/webapps/ROOT && \
  ln -s /var/lib/tomcat8/webapps/guacamole.war /var/lib/tomcat8/webapps/ROOT.war && \

  # Cleanup
  apt-get remove -y build-essential && \
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/*

ADD /files/ /

ENV JAVA_HOME /usr/lib/jvm/java-8-oracle
