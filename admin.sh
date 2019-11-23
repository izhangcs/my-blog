#!/bin/bash

SERVER=blog
BASE_DIR=`pwd`
INTERVAL=2

ARGS=""

function start()
{
    if [ "`pgrep ${SERVER} -u $UID`" != "" ]; then
        echo "${SERVER} is alrealy running."
        exit 1
    fi
    
    nohup ${BASE_DIR}/${SERVER} ${ARGS} &>/dev/null &
    
    echo "sleeping" && sleep ${INTERVAL}

    if [ "`pgrep ${SERVER} -u ${UID}`" == "" ]; then
        echo "${SERVER} start failed."
        exit 1
    fi
    
    echo "${SERVER} start sucessfully."
}

function status()
{
    if [ "`pgrep ${SERVER} -u $UID`" != "" ];then 
        echo "${SERVER} is alrealy running."
    else
        echo "${SERVER} is not running."
    fi
}

function stop()
{
    if [ "`pgrep ${SERVER} -u ${UID}`" != "" ]; then 
        kill -9 `pgrep ${SERVER} -u ${UID}`
    else
        echo "${SERVER} is not running."
        exit 1
    fi

    echo "sleeping" && sleep ${INTERVAL}

    if [ "`pgrep ${SERVER} -u ${UID}`" != "" ]; then
        echo "${SERVER} stop failed."
        exit 1
    fi

    echo "${SERVER} stop successfully."
}


case $1 in
    'start')
        start 
    ;;
    'stop')
        stop
    ;;
    'status')
        status
    ;;
    'restart')
        stop && start
    ;; 
    *)
        echo "usage: $0 {start|stop|restart|status}"
	    exit 1
    ;;
esac