FROM ruby:2.3

ENV GITHUB_INSTALL_SCRIPTS=https://raw.githubusercontent.com/r888888888/danbooru/master/script/install

RUN \
  echo "deb http://www.deb-multimedia.org jessie main non-free" | tee -a /etc/apt/sources.list.d/deb-multimedia.list && \
  echo "deb-src http://www.deb-multimedia.org jessie main non-free" | tee -a /etc/apt/sources.list.d/deb-multimedia.list && \
  apt-get update -qq && \
  apt-get install -qy --force-yes deb-multimedia-keyring && \
  apt-get update -qq && \
  apt-get -qy --no-install-recommends install \
    automake \
    bison \
    build-essential \
    coreutils \
    curl \
    ffmpeg \
    flex \
    git \
    imagemagick \
    libcurl4-openssl-dev \
    libmagickcore-dev \
    libmagickwand-dev \
    libmemcached-dev \
    libpq-dev \
    libreadline-dev \
    libssl-dev \
    libxml2-dev \
    libxslt-dev \
    memcached \
    mkvtoolnix \
    ncurses-dev \
    nginx \
    postgresql \
    postgresql-contrib \
    postgresql-server-dev-all \
    ragel \
    sendmail \
    sendmail-bin \
    ssh \
    sudo \
    supervisor


RUN \
  # Create directories
  mkdir -p /config /logs /app && \

  # Create user account
  useradd -d /config -s /bin/bash danbooru && \

  # Install gems
  gem install --no-ri --no-rdoc bundler && \

  # Set up PostgreSQL
  export PG_VERSION=`pg_config --version | egrep -o '[0-9]{1,}\.[0-9]{1,}'` && \
  curl -L -s $GITHUB_INSTALL_SCRIPTS/postgresql_hba_conf -o /etc/postgresql/$PG_VERSION/main/pg_hba.conf && \
  /etc/init.d/postgresql start && \
  sudo -u postgres createuser -s danbooru && \
  sudo -u danbooru createdb danbooru2 && \

  # Set up nginx
  curl -L -s $GITHUB_INSTALL_SCRIPTS/nginx.danbooru.conf -o /etc/nginx/sites-enabled/danbooru.conf && \

  # Set up logrotate
  curl -L -s $GITHUB_INSTALL_SCRIPTS/danbooru_logrotate_conf -o /etc/logrotate.d/danbooru.conf && \

  # Install app
  git clone --depth=1 git://github.com/r888888888/danbooru.git /app && \
  cd /app && \
  bundle install && \

  # mkdir -p /var/www/danbooru2/shared/config && \
  # mkdir -p /var/www/danbooru2/shared/data && \
  # mkdir -p /var/www/danbooru2/shared/data/preview && \
  # mkdir -p /var/www/danbooru2/shared/data/sample && \
  # chown -R danbooru:danbooru /var/www/danbooru2 && \
  curl -L -s $GITHUB_INSTALL_SCRIPTS/database.yml.templ -o /app/config/database.yml && \
  curl -L -s $GITHUB_INSTALL_SCRIPTS/danbooru_local_config.rb.templ -o /app/config/danbooru_local_config.rb

ADD /files/ /

EXPOSE 80

CMD ["/app/start.sh"]
