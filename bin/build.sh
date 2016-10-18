#!/usr/bin/env bash

p=`pwd`
for d in $(ls ./code); do
  echo "building code/$d"
  cd $p/code/$d
  env GOOS=linux GOARCH=386 go build
done
cd $p