#!/bin/bash

touch /var/lib/tomcat7/logs/catalina.out

cd /var/lib/tomcat7
exec /usr/lib/jvm/java-8-oracle/jre/bin/java \
  -Djava.util.logging.config.file=/var/lib/tomcat7/conf/logging.properties \
  -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager \
  -Djava.awt.headless=true -Xmx128m -XX:+UseConcMarkSweepGC \
  -Djava.endorsed.dirs=/usr/share/tomcat7/endorsed \
  -classpath /usr/share/tomcat7/bin/bootstrap.jar:/usr/share/tomcat7/bin/tomcat-juli.jar \
  -Dcatalina.base=/var/lib/tomcat7 \
  -Dcatalina.home=/usr/share/tomcat7 \
  -Djava.io.tmpdir=/tmp/tomcat7-tomcat7-tmp \
  org.apache.catalina.startup.Bootstrap \
  start
