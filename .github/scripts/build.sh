#!/bin/env bash

echo "Building distribution"
echo "====================="
BUILD_TIME=$(date +%Y-%m-%dT%H:%M:%S)
BUILD_HOST=$(hostname)
gh_release=${GITHUB_REF#/refs/*/}
RELEASE=$(git tag -l --sort=refname "v*" | tail -n 1 | sed -e 's/^v//')
BUILD_RUNNER="$(whoami)@$(hostname)"

echo "Build date:   ${BUILD_TIME}"
echo "Build runner: ${BUILD_RUNNER}"
echo "Release:      ${RELEASE}"
repo="github.com/guionardo/go-dev"
flags="-X main.build_date=${BUILD_TIME} -X main.build_runner=${BUILD_RUNNER} -X main.release=${RELEASE}"
echo "Flags:      ${flags}"
go build -ldflags "$flags" .
