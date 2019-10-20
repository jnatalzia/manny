#! /bin/bash

trap "./kill_gogram.sh" INT

./build.sh

fswatch -o ./ -e ".*" -i "\\.go$" -i "\\.csv$" | xargs -n1 './build.sh'