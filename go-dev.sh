#!/bin/bash
GO_DEV=/home/guionardo/dev/github.com/guionardo/go-dev/go-dev

function list_folders() {
  ($GO_DEV list)
}

function update_folders() {
  ($GO_DEV update)
}

function setup_folders() {
  ($GO_DEV setup $2)
}

if [[ "$1" == "" ]]; then
  ($GO_DEV)
else
  case $1 in
  list)
    list_folders
    ;;
  update)
    update_folders
    ;;
  setup)
    setup_folders "$2"
    ;;

  *)
    exec 5>&1
    RESULT=$($GO_DEV go $@ | tee >(cat - >&5))
    for ROW in ${RESULT[@]}; do
      LAST_ROW="$ROW"
    done
    if [[ -f $LAST_ROW ]]; then
      source $LAST_ROW
      rm $LAST_ROW
    fi
    ;;
  esac
fi
