#!/usr/bin/env bash
cd phoenix
go build -v
./phoenix -plugins lol
cd ..
