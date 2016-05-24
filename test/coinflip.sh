#!/bin/sh

action="$(echo ${1} | jq -r .action)"

case $action in
    init)
        echo '{"name": "bot1"}'
        ;;
    *)
        echo '{"error": "not implemented action: $action"}'
        ;;
esac
