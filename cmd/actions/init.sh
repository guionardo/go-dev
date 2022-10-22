#!/bin/bash
function dev() {
  GO_DEV={GO_DEV}
  CMD="go"
  OPTIONS=""

  output_file="{GO_OUTPUT}"
  
  case $1 in
  {GO_CMDS})
  # list | update | setup | help | install)
    CMD="$1"
    shift 1
    ;;


  esac

  # shellcheck disable=SC2068
  $GO_DEV --output=$output_file "$CMD" $@ "$OPTIONS"

  if [[ -f $output_file ]]; then
    source $output_file
    rm "$output_file"
  fi
}