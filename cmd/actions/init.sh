#!/bin/bash
# shellcheck disable=SC2068
function _dev() {
  if [[ ! -f "{GO_OUTPUT}" ]]; then
    echo "No output file found: {GO_OUTPUT}"
    exit 1
  fi
  source "{GO_OUTPUT}"
  rm "{GO_OUTPUT}"
}

function dev() {
  {GO_DEV} go --output="{GO_OUTPUT}" $@ && _dev
}

function devdbg() {
  {GO_DEV} --debug go --output="{GO_OUTPUT}" $@ && _dev
}

echo "go-dev is ready to use (dev, devdbg)"
