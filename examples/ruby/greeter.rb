#!/usr/bin/ruby

# Autoflush output
STDOUT.sync = true

# For each line FOO received on STDIN, respond with "Hello FOO!".
while 1
  line = STDIN.readline.strip
  puts "Hello #{line}!"
end
