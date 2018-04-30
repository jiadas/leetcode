#!/usr/bin/env bash

file_name=$2".go"
file_name=`echo ${file_name}|sed 's/[ ][ ]*/_/g'|tr A-Z a-z`
echo ${file_name}
echo "package main" > $1/${file_name}
git add $1/${file_name}