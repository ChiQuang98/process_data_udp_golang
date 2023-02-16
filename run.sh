#!/bin/bash
nohup go run main.go &

#command to get source
#tcpdump -i any -n -c 5 ip | awk '{ print gensub(/(.*)\..*/,"\\1","g",$3), $4, gensub(/(.*)\..*/,"\\1","g",$5) }'
