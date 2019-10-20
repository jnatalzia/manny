#! /bin/bash

go build -o dist/main
server_file=server.js
start_go_cmd="DEBUG=t PROJECTROOT=./ ./dist/main >> /tmp/log.log &"
echo $start_go_cmd
 
./kill_gogram.sh
 
echo "building binary"
go build -o dist/main
echo "starting program"
eval "$start_go_cmd"  