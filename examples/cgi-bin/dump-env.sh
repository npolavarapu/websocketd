#!/bin/bash

# Standard CGI(ish) environment variables, as defined in
# http://tools.ietf.org/html/rfc3875
NAMES="""
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
"""

echo "Content-type: text/plain"
echo

for NAME in ${NAMES}
do
	echo ${NAME}=${!NAME:-<unset>}
done

# Additional HTTP headers
env | grep '^HTTP_'
