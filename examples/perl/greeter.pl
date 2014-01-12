#!/usr/bin/perl

use strict;

# Autoflush output
use IO::Handle;
STDOUT->autoflush(1);

# For each line FOO received on STDIN, respond with "Hello FOO!".
while (<>) {
  chomp; # remove \n
  print "Hello $_!\n";
}
