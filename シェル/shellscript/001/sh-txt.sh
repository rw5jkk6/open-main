#!/bin/bash

case "$1" in
  *.sh)
    touch $1
    chmod +x $1
    vim $1
    ;;
  *.txt)
    vim $1
    ;;
esac
