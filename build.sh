#!/bin/sh

# build all resources and store them in resources.go
python3 resources/build.py

# rebuild the app
go build

# run 
./talkapply --port=54321