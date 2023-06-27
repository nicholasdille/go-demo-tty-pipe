#!/bin/bash

STDIN=0
STDOUT=1
STDERR=2

if test -t ${STDIN}; then
    echo "stdin is attached to terminal" >&2
else
    echo "stdin is not attached to terminal" >&2
fi
if test -p /dev/stdin; then
    echo "stdin is a pipe" >&2
else
    echo "stdin is not a pipe" >&2
fi

if test -t ${STDOUT}; then
    echo "stdout is attached to terminal" >&2
else
    echo "stdout is not attached to terminal" >&2
fi
if test -p /dev/stdout; then
    echo "stdout is a pipe" >&2
else
    echo "stdout is not a pipe" >&2
fi