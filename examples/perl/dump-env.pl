#!/usr/bin/perl



use strict;

# Autoflush output
use IO::Handle;
STDOUT->autoflush(1);

# Standard CGI(ish) environment variables, as defined in
# http://tools.ietf.org/html/rfc3875
my @names = qw(
  AUTH_TYPE
  CONTENT_LENGTH
  CONTENT_TYPE
  GATEWAY_INTERFACE
  PATH_INFO
  PATH_TRANSLATED
  QUERY_STRING
  REMOTE_ADDR
  REMOTE_HOST
  REMOTE_IDENT
  REMOTE_PORT
  REMOTE_USER
  REQUEST_METHOD
  REQUEST_URI
  SCRIPT_NAME
  SERVER_NAME
  SERVER_PORT
  SERVER_PROTOCOL
  SERVER_SOFTWARE
  UNIQUE_ID
);

for my $name (@names) {
	my $value = $ENV{$name} || '<unset>';
	print "$name=$value\n";
}

# Additional HTTP headers
for my $name (keys(%ENV)) {
	if ($name =~ /^HTTP_/) {
		my $value = $ENV{$name};
		print "$name=$value\n";
	}
}
