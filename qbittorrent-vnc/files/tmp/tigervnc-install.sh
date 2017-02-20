#!/bin/bash

# Install TigerVNC
curl -L https://dl.bintray.com/tigervnc/stable/tigervnc-1.7.0.x86_64.tar.gz | tar -xz --strip=1 -C /

# vncpasswd.py
mkdir -p /opt/vncpasswd
curl -L https://github.com/trinitronx/vncpasswd.py/archive/master.tar.gz | tar -zx --strip=1 -C /opt/vncpasswd

# Add service to init
cat <<'EOT' > /etc/init.d/vnc
#!/bin/bash
source /opt/default-values.sh
case "$1" in
  start)
    /usr/bin/sv start crashplan
    /usr/bin/sv start openbox
    ;;
  stop)
    /usr/bin/sv stop crashplan
    /usr/bin/sv stop openbox
    ;;
  restart)
    /usr/bin/sv restart crashplan
    /usr/bin/sv restart openbox
    ;;
  status)
    eval 'exec 6<>/dev/tcp/127.0.0.1/${VNC_PORT} && echo "running" || echo "stopped"' 2>/dev/null
    exec 6>&- # close output connection
    exec 6<&- # close input connection
    ;;
esac
EOT
chmod +x /etc/init.d/vnc
