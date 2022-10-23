#!/bin/bash

metadata_file="./cmd/configuration/metadata.txt"

function get_key() {
  key="$1"
  echo "$(grep $key $metadata_file | cut -d'=' -f2)"
}

function set_key() {
  key="$1"
  value="$2"
  tmpfile=$(mktemp /tmp/go-dev-metadata.XXXXXX)
  grep -v $key $metadata_file > $tmpfile
  echo "$key=$value" >> $tmpfile
  cat $tmpfile > $metadata_file
  rm $tmpfile
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
  elif [[ "$2" == "bug" ]]; then
    build=$(echo $build + 1 | bc)
  elif [[ "$2" == "major" ]]; then
    major=$(echo $major+1 | bc)
  else
    echo "usage: ./version.sh version_number [major/feature/bug]"
    exit -1
  fi

  # echo the new version number
  echo "${major}.${minor}.${build}"

}

function change_version() {
  version=$(get_key version)
  new_version=$(update_version "$version" "$1")
  set_key version "$new_version"

}

case $1 in
  version_major)
    change_version "major"
    ;;
  version_feature)
    change_version "feature"
    ;;
  version_bug)
    change_version "bug"
    ;;
  release)
    change_version "feature"
    set_key build_date $(date +'%Y-%m-%dT%H:%M:%S')
    ;;
esac
#get_key "version"
#set_key "name" "go-dev-2"
set_key "builder_info" $USER@$(hostname)