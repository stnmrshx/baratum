#!/usr/bin/env bash

p=`pwd`
for d in $(ls ./code); do
  echo "verifying code/$d"
  cd $p/code/$d
  go fmt
  golint
done
cd $p
