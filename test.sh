#!/bin/bash

SCRIPT=$1

echo "Running: ${SCRIPT}"
${SCRIPT}
echo

echo "Running: echo foo | ${SCRIPT}"
echo foo | ${SCRIPT}
echo

echo "Running: ${SCRIPT} </dev/null"
${SCRIPT} </dev/null
echo
