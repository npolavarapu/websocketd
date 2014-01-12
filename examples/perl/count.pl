#!/usr/bin/perl

use strict;

# Autoflush output
use IO::Handle;
STDOUT->autoflush(1);

use Time::HiRes qw(sleep);

# Simple example script that counts to 10 at ~2Hz, then stops.
for my $count (1 .. 10) {
	print "$count\n";
	sleep 0.5;
}
