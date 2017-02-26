#!/bin/bash
VNC_PORT="${VNC_PORT:-5900}"
eval 'exec 6<>/dev/tcp/127.0.0.1/${VNC_PORT} && echo "running" || echo "stopped"' 2>/dev/null
exec 6>&- # close output connection
exec 6<&- # close input connection
