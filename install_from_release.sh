#!/bin/bash

LATEST_RELEASE=https://github.com/guionardo/go-dev/releases/latest
#URL_RELEASE=https://github.com/guionardo/go-dev/releases/download/v0.3.0/go-dev-v0.3.0-linux-amd64.tar.gz
#tmp=$(mktemp)
#curl -s $LATEST_RELEASE | \
#  grep -io '<a href=['"'"'"][^"'"'"']*['"'"'"]' | \
#  sed -e 's/^<a href=["'"'"']//i' -e 's/["'"'"']$//i'
#curl \
#  -H "Accept: application/vnd.github.v3+json" \
#  https://api.github.com/repos/guionardo/go-dev/releases/latest