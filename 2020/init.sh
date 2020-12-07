#!/bin/bash

if [ ! -d "day$1" ]
then
    mkdir day$1
fi

touch day$1/main.go
touch day$1/input
