#!/bin/bash

count=0
while [ $count != 5 ]
do 
  ping -c 1 example.com
  sleep 3
  count=$(($count + 1))
done
