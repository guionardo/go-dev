#!/bin/bash
DEV_FOLDER_PY=/home/guionardo/dev/lab/dev/dev_folder.py
DEV_FOLDER_GO=/home/guionardo/dev/lab/dev/dev_go/dev_go

SAIR=0
if [[ "$1" == "list" ]]; then
    ($DEV_FOLDER_GO list)
    SAIR=1
else 
    if [[ "$1" == "" ]]; then
        echo "Informe uma expressão de busca"
        SAIR=1
    else 
        if [[ "$1" == "update" ]]; then
            ($DEV_FOLDER_GO update)
            SAIR=1
        fi
    fi
fi

if [[ $SAIR -eq 0 ]]; then
    # RESULT=$($DEV_FOLDER_GO $@)
    exec 5>&1
    RESULT=$($DEV_FOLDER_GO go $@|tee >(cat - >&5))
    for ROW in ${RESULT[@]}
    do
        RESULT_CD="$ROW"
    done
    RESULT=$RESULT_CD

    # RESULT=$($DEV_FOLDER_PY go $@)
    if [[ -d "$RESULT" ]]; then
        echo "cd $RESULT"
        cd $RESULT
    else
        echo $RESULT
    fi
fi
