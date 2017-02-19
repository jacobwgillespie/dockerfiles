#!/bin/bash

. /opt/default-values.sh

# VNC credentials
if [ ! -f "${VNC_CREDENTIALS}" -a -n "${VNC_PASSWD}" ]; then
  /opt/vncpasswd/vncpasswd.py -f "${VNC_CREDENTIALS}" -e "${VNC_PASSWD}"
fi
