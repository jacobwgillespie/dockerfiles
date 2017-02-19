#!/bin/bash

PID=$(ps aux | grep -i [s]startapp.sh | awk {'print $2'})

if [ -n "$PID" ]; then
  kill $PID
  sleep 1
fi

PID=$(ps aux | grep -i [s]startapp.sh | awk {'print $2'})

if [ -n "$PID" ]; then
  kill -9 $PID
fi

PID=$(ps aux | grep -i [q]bittorrent | awk {'print $2'})

if [ -n "$PID" ]; then
  kill -9 $PID
fi
