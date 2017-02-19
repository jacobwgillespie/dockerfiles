#!/bin/bash
umask 0000

export SWT_GTK3=0

i=0
until [ "$(/etc/init.d/qbittorrent status)" == "running" ]; do
  sleep 1
  let i+=1
  if [ $i -gt 10 ]; then
    break
  fi
done

mkdir -p /root/.config/qBittorrent

qbittorrent \
   > /root/.config/qBittorrent/desktop_output.log \
  2> /root/.config/qBittorrent/desktop_error.log
