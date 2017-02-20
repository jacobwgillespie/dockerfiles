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
  kill $PID
fi

i=0
until [ "$(/etc/init.d/qbittorrent status)" == "stopped" ]; do
  sleep 1
  let i+=1
  if [ $i -gt 10 ]; then
    break
  fi
done

PID=$(ps aux | grep -i [q]bittorrent | awk {'print $2'})

if [ -n "$PID" ]; then
  kill -9 $PID
fi
