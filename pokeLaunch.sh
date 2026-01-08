#!/bin/bash
if [[ "$TERM" == "linux" ]]
then
    exit
else
    cd ~/Documents/goPokeImgQuery/ || return
    go run main.go responseFeilds.go && ./cmdlnImageV2 CPoke.png
	 if [ "$RANDOM" -lt 3277 ]; then {
		echo "	------------ditto-------------"	> pokeName.txt
	 } fi
	 cat pokeName.txt
	 touch IRAN.txt
fi
