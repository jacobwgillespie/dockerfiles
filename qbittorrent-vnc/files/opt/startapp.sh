#!/bin/bash
umask 0000

export SWT_GTK3=0

i=0
until [ "$(/etc/init.d/vnc status)" == "running" ]; do
  sleep 1
  let i+=1
  if [ $i -gt 10 ]; then
    break
  fi
done

killall qbittorrent

while true; do
  su - abc -s/bin/bash -c "qbittorrent \
     > /config/desktop_output.log \
    2> /config/desktop_error.log"
done
