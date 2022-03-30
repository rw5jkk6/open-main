#!/bin/bash

tmp=0
for i in 3 2 4 1
do 
  if [ $i -gt $tmp ];then
    tmp=$i
  fi
done
echo $tmp
