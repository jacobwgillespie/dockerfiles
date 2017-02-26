#!/bin/bash
umask 0000

export SWT_GTK3=0

while true; do
  HOME=/config qbittorrent \
     > /config/desktop_output.log \
    2> /config/desktop_error.log
done
