#!/bin/sh

while true
do
    goose postgres "postgres://${DB_USER}:${DB_PASSWORD}@${DB_URL}" up 2>/dev/null
    if [ $? -eq 0 ]
    then
        echo 'migration is done'
        break
    else
        sleep 1
    fi
done