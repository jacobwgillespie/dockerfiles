#!/bin/bash
umask 0000

while true; do
  TERM=xterm HOME=/config dbus-run-session dolphin
done
