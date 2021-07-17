#!/bin/bash
DEV_FOLDER_GO=/home/guionardo/dev/lab/dev/dev_go/dev_go

function list_folders() {
  ($DEV_FOLDER_GO list)
}

function update_folders() {
  ($DEV_FOLDER_GO update)
}

function setup_folders() {
  ($DEV_FOLDER_GO setup $2)
}

if [[ "$1" == "" ]]; then
  ($DEV_FOLDER_GO)
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
    RESULT=$($DEV_FOLDER_GO go $@ | tee >(cat - >&5))
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
