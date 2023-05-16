#!/bin/sh
go build -o solar-autogtk && ./solar-autogtk "$@" > /dev/null 2>&1 
