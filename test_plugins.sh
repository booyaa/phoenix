#!/bin/env bash

COMMAND=ping

curl -XPOST 'http://localhost:8080/?token=abcdefg&team_id=1001&channel_id=C12345&channel_name=foo_channel&timestamp=1355517523.000005&user_id=1234&user_name=test&text=phoenix%20'${COMMAND}'&trigger_word=phoenix'
