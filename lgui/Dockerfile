# dockergui
FROM phusion/baseimage:0.9.19

#########################################
##        ENVIRONMENTAL CONFIG         ##
#########################################
# Set correct environment variables
ENV LC_ALL=C.UTF-8 LANG=en_US.UTF-8 LANGUAGE=en_US.UTF-8 TERM=xterm


# Use baseimage-docker's init system
CMD ["/sbin/my_init"]

#########################################
##         RUN INSTALL SCRIPT          ##
#########################################
COPY ./install/ /tmp/
RUN \
  export DEBIAN_FRONTEND="noninteractive" && \
  usermod -u 99 nobody && \
  usermod -g 100 nobody && \
  usermod -m -d /nobody nobody && \
  usermod -s /bin/bash nobody && \
  usermod -a -G adm,sudo nobody && \
  echo "nobody:PASSWD" | chpasswd && \

  # Disable SSH
  rm -rf /etc/service/sshd /etc/my_init.d/00_regen_ssh_host_keys.sh && \

  #########################################
  ##    REPOSITORIES AND DEPENDENCIES    ##
  #########################################

  # Repositories
  echo 'deb mirror://mirrors.ubuntu.com/mirrors.txt xenial main universe restricted' > /etc/apt/sources.list && \
  echo 'deb mirror://mirrors.ubuntu.com/mirrors.txt xenial-updates main universe restricted' >> /etc/apt/sources.list && \
  echo 'deb http://ppa.launchpad.net/webupd8team/java/ubuntu xenial main' >> /etc/apt/sources.list && \
  apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys EEA14886 && \

  # Install Dependencies
  apt-get update -qq && \
  # Install general
  apt-get install -qy --no-install-recommends wget \
                              				unzip && \

  # Install window manager and x-server
  apt-get install -qy --no-install-recommends \
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
    x11-apps \
    openbox

RUN \

  # x11rdp install
  dpkg -i /tmp/x11rdp/x11rdp_0.9.0+devel-1_amd64.deb && \

  # xrdp needs to be installed seperately
  dpkg -i /tmp/x11rdp/xrdp_0.9.0+devel_amd64.deb && \

  # Install Guac
  echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections && \
  apt-get install -y --force-yes oracle-java8-installer && \
  apt-get install -y --force-yes oracle-java8-set-default && \
  apt-get install -qy --no-install-recommends libossp-uuid-dev \
                                                          libpng12-dev \
                                                          libfreerdp-dev \
                                                          libcairo2-dev \
                                                          tomcat7

RUN \


  #########################################
  ##  FILES, SERVICES AND CONFIGURATION  ##
  #########################################

  # User directory
  mkdir /nobody && \
  mkdir -p /nobody/.config/openbox && \
  mkdir /nobody/.cache && \

  chmod -R +x /etc/service/ /etc/my_init.d/ && \

  #########################################
  ##             INSTALLATION            ##
  #########################################

  # Make needed directories
  mkdir -p /var/cache/tomcat7 && \
  mkdir -p /var/lib/guacamole/classpath && \
  mkdir -p /usr/share/tomcat7/.guacamole && \
  mkdir -p /usr/share/tomcat7-root/.guacamole && \
  mkdir -p /root/.guacamole && \

  # Install guacd
  dpkg -i /tmp/guacamole/guacamole-server_0.9.7_amd64.deb && \
  ldconfig && \

  # Configure tomcat
  cp /tmp/guacamole/guacamole-0.9.7.war /var/lib/tomcat7/webapps/guacamole.war && \
  cp /tmp/guacamole/guacamole-auth-noauth-0.9.7.jar /var/lib/guacamole/classpath && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat7/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /usr/share/tomcat7-root/.guacamole/ && \
  ln -s /etc/guacamole/guacamole.properties /root/.guacamole/ && \

  # Fix tomcat webroot
  rm -Rf /var/lib/tomcat7/webapps/ROOT && \
  ln -s /var/lib/tomcat7/webapps/guacamole.war /var/lib/tomcat7/webapps/ROOT.war && \

  ### Compensate for GUAC-513
  mkdir -p /usr/lib/x86_64-linux-gnu/freerdp/ && \
  ln -s /usr/local/lib/freerdp/guacsnd.so /usr/lib/x86_64-linux-gnu/freerdp/ && \
  ln -s /usr/local/lib/freerdp/guacdr.so /usr/lib/x86_64-linux-gnu/freerdp/ && \

  apt-get install -qy libfreerdp-dev && \

  # pulseauido rdp
  # cp /tmp/x11rdp/module-xrdp* /usr/lib/pulse-4.0/modules && \
  # chown -R 777 /usr/lib/pulse-4.0/modules && \

  #########################################
  ##                 CLEANUP             ##
  #########################################

  # Clean APT install files
  apt-get autoremove -y && \
  apt-get clean -y && \
  rm -rf /var/lib/apt/lists/* /var/cache/* /var/tmp/* /tmp/guacamole /tmp/openbox /tmp/x11rdp

ADD /files/ /
