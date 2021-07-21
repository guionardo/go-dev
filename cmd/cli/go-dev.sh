#!/bin/bash
GO_DEV={GO_DEV}
CMD="go"
OPTIONS=""

output_file="$HOME/.go-dev.output.sh"

case $1 in
list | update | setup | help)
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
