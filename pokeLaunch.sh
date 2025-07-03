#!/bin/bash
if [[ "$TERM" == "linux" ]]
then
    exit
else
    cd ~/Documents/goPokeImgQuery/
    go run main.go responseFeilds.go && ./cmdlnImageV2 CPoke.png && cat pokeName.txt
fi
