#!/bin/bash

echo "Running: ./demo"
./demo
echo

echo "Running: echo foo | ./demo"
echo foo | ./demo
echo

echo "Running: ./demo </dev/null"
./demo </dev/null
echo
