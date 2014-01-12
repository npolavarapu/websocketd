#!/bin/bash

# For each line FOO received on STDIN, respond with "Hello FOO!".
while read LINE
do
	echo "Hello $LINE!"
done
