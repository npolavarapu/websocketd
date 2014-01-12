#!/usr/bin/ruby

# Autoflush output
STDOUT.sync = true

# Simple example script that counts to 10 at ~2Hz, then stops.
(1..10).each do |count| 
	puts count
	sleep(0.5)
end
