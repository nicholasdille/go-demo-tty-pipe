#!/bin/bash

STDIN=0
STDOUT=1
STDERR=2

if test -t ${STDIN}; then
    echo "stdin is attached to terminal"
else
    echo "stdin is not attached to terminal"
fi

if test -p /dev/stdin; then
    echo "stdin is a pipe"
else
    echo "stdin is not a pipe"
fi