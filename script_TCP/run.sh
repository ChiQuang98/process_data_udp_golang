#!/bin/bash

APP_NAME="TCP_Packet"
APP_BIN="./${APP_NAME}"
PID_FILE="${APP_NAME}.pid"
mkdir -p /u01/app_mobileid/script_TCP/levelDB/ips
case "$1" in
  start)
    if [ -f $PID_FILE ]; then
      echo "Error: $APP_NAME is already running."
      exit 1
    fi
    nohup $APP_BIN > /dev/null 2>&1 &
    echo $! > $PID_FILE
    echo "$APP_NAME started."
    ;;
  stop)
    if [ ! -f $PID_FILE ]; then
      echo "Error: $APP_NAME is not running."
      exit 1
    fi
    kill $(cat $PID_FILE)
    rm $PID_FILE
    echo "$APP_NAME stopped."
    ;;
  status)
    if [ -f $PID_FILE ]; then
      echo "$APP_NAME is running with PID $(cat $PID_FILE)."
    else
      echo "$APP_NAME is not running."
    fi
    ;;
  *)
    echo "Usage: $0 {start|stop|status}"
    exit 1
esac

exit 0


#command to get source
#tcpdump -i any -n -c 5 ip | awk '{ print gensub(/(.*)\..*/,"\\1","g",$3), $4, gensub(/(.*)\..*/,"\\1","g",$5) }'
#curl -X POST -H "Content-Type: application/json" -d '{"ips":["42.1.66.19"]}' 0.0.0.0:7001/update/v1/ips
#curl -X DELETE -H "Content-Type: application/json" -d '{"ips":["8.8.8.8","103.199.79.26","42.1.64.19","8.8.8.8"]}' 0.0.0.0:7001/delete/v1/ips
#htop -p $(pgrep -d',' -f "go run main.go")
#
#YHA:
#    "SrcUDP1": "10.51.46.101",
#    "SrcUDP2": "10.51.46.103",

