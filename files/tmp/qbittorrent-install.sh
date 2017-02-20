#!/bin/bash

add-apt-repository -y ppa:qbittorrent-team/qbittorrent-stable
apt-get update

apt-get install -qy qbittorrent

cat <<'EOT' > /etc/init.d/qbittorrent
#!/bin/bash
source /opt/default-values.sh
case "$1" in
  start)
    ;;
  stop)
    ;;
  restart)
    ;;
  status)
    PID=$(ps aux | grep -i [q]bittorrent | awk {'print $2'})
    if [ -n "$PID" ]; then
      echo "running"
    else
      echo "stopped"
    fi
    ;;
esac
EOT
chmod +x /etc/init.d/qbittorrent
