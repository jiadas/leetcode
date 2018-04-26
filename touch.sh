#!/usr/bin/env bash

file_name=$1".go"
file_name=`echo ${file_name}|sed 's/[ ][ ]*/_/g'|tr A-Z a-z`
echo ${file_name}
echo "package main" > ${file_name}
git add ${file_name}