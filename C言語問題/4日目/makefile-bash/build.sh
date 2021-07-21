#!/bin/bash

gcc -c -o calc.o calc.c
gcc -c -o main.o main.c
gcc -o main calc.o main.o
