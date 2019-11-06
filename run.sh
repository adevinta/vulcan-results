#!/bin/sh

if [ -f "$1" ]; then
  cat $1 | sed 's/LogFile *=.*/LogFile = ""/g' > config.toml
else
  echo "ERROR: Expected config file"
  echo "Usage: $0 config.toml"
fi

# Force logs to STDOUT
/vulcan-results config.toml
