#!/usr/bin/python

from sys import stdout
from time import sleep

# Simple example script that counts to 10 at ~2Hz, then stops.
for count in range(0, 10):
  print count + 1
  sleep(0.5)
  stdout.flush() # Remember to flush
