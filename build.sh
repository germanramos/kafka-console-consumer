#!/bin/sh
source env.sh
echo "Building..."
CGO_ENABLED=0 go build -v -a -installsuffix cgo main
ERROR=$?
if [ -n "$GO_UID" ]; then
  echo "Setting $GO_UID:$GO_GID to main..."
  chown $GO_UID:$GO_GID main
fi
echo "Exiting with code $ERROR"
exit $ERROR
