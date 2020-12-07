#!/bin/bash

if [ ! -d "day$1" ]
then
    mkdir day$1
fi

if [ ! -f "day$1/main.go" ] 
then
    cd day$1
    touch main.go
    curl -O --cookie "session=<enter cookie>" https://adventofcode.com/2020/day/$1/input 
fi
