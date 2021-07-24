#!/bin/bash
if [ -x "$(command -v docker)" ]; then
  export DOCKER_BUILDKIT=1
  docker build https://github.com/guionardo/go-dev.git#develop --target bin --output bin/ --platform local
  bin/go-dev install
  bin/go-dev help
  echo "Open a new console to load updated configurations"
  echo "Usage: dev <options>"
else
  echo "Docker is not installed"
fi