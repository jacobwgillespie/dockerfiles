# Load default values if empty
VNC_RESOLUTION="${VNC_RESOLUTION:-1920x1080}"
VNC_CREDENTIALS=/nobody/.vnc_passwd
VNC_PORT="${VNC_PORT:-5900}"

if [[ -n $VNC_PASSWD ]]; then
  VNC_SECURITY="SecurityTypes TLSVnc,VncAuth -PasswordFile ${VNC_CREDENTIALS}"
else
  VNC_SECURITY="SecurityTypes None"
fi
