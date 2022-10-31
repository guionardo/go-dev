#!/bin/env bash

BUILD_TIME=$(date +"%Y-%m-%dT%H:%M:%S%:z")
BUILD_RUNNER="$(whoami)@$(hostname)"
RELEASE="v0.0.0"

function check_gh {
    gh auth status -t 2>&1 >/dev/null | grep "✓ Token: gho_" -q
    if [ $? -ne 0 ]; then
        echo "ERROR: You must be logged in to GitHub with gh cli installed"
        exit 1
    fi
    echo "✓ GitHub is logged in"
}

function check_tests {
    echo "Running tests"
    echo "============="
    go test ./... -failfast | grep FAIL
    if [ $? -ne 1 ]; then
        echo "ERROR: Tests failed"
        exit 1
    fi
    echo "✓ Tests passed"
}

function get_release {
    LAST_RELEASE=$(git tag -l --sort=refname "v*" | tail -n 1 | sed -e 's/^v//')
    echo "✓ Last release: $LAST_RELEASE"
    if [ -z "$1" ]; then
        input_release
    else
        get_next_release $1
    fi
    RELEASE="v${VERSION}"
    echo "✓ Release version is ${RELEASE}"
}

function get_next_release {
    VERSION=$(update_version "$LAST_RELEASE" "$1")
}

function update_version() {
    version="$1"
    major=0
    minor=0
    build=0

    # break down the version number into it's components
    regex="([0-9]+).([0-9]+).([0-9]+)"
    if [[ $version =~ $regex ]]; then
        major="${BASH_REMATCH[1]}"
        minor="${BASH_REMATCH[2]}"
        build="${BASH_REMATCH[3]}"
    fi

    # check parameter to see which number to increment
    if [[ "$2" == "feature" ]]; then
        minor=$(echo $minor + 1 | bc)
        build=0
    elif [[ "$2" == "bug" ]]; then
        build=$(echo $build + 1 | bc)
    elif [[ "$2" == "major" ]]; then
        major=$(echo $major+1 | bc)
        build=0
        minor=0
    else
        echo "usage: ./version.sh version_number [major/feature/bug]"
        exit -1
    fi

    # echo the new version number
    echo "${major}.${minor}.${build}"
}

function input_release {
    echo "Input release version (e.g. 1.0.0):"
    read VERSION
    if [ -z "$VERSION" ]; then
        echo "ERROR: Release version is required"
        exit 1
    fi
    if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        echo "ERROR: Release version must be in format x.y.z"
        exit 1
    fi
    git tag -l --sort "v*" | grep "v${VERSION}" -q
    if [ $? -eq 0 ]; then
        echo "ERROR: Release version already exists"
        exit 1
    fi

}

function build {
    echo "====================="
    echo "Building distribution"
    echo "====================="
    echo "Build date:   ${BUILD_TIME}"
    echo "Build runner: ${BUILD_RUNNER}"
    echo "Release:      ${RELEASE}"
    jq --null-input \
        --arg t "$BUILD_TIME" \
        --arg r "$BUILD_RUNNER" \
        --arg v "$VERSION" \
        '{"build_time": $t, "build_runner": $r, "version": $v}' >pkg/consts/metadata.json
}

function create_release {
    echo "====================="
    echo "Adding metadata"
    git add pkg/consts/metadata.json
    git commit -m "Release ${RELEASE}"
    git tag -a "${RELEASE}" -m "Release ${RELEASE}"
    git push 
    gh release create $RELEASE
}

check_tests
check_gh
get_release $1
build
check_tests
create_release
