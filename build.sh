#!/bin/sh

# build all resources and store them in resources.go
python3 resources/build.py

# run go fmt
go fmt

# rebuild the app
go build

# run 
./talkapply --port=54321 --seconds=30 --file="talkapply.backup"