#!/bin/bash

chown -R abc:abc /openbox/

if [[ -e /opt/startapp.sh ]]; then
  chown abc:abc /opt/startapp.sh
  chmod +x /opt/startapp.sh
fi
