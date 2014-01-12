#!/usr/bin/python

from sys import stdin, stdout

# For each line FOO received on STDIN, respond with "Hello FOO!".
while True:
  line = stdin.readline().strip()
  print 'Hello %s!' % line
  stdout.flush() # Remember to flush
