#!/usr/bin/env bash

curl -s http://www.cmbc.com.cn/jrms/msdt/yjbg/index.htm | grep -P 'href="(.*pdf")' -o | sed -n "s/^href=\"\(.*\)\"/http:\/\/www.cmbc.com.cn\1/p" | xargs wget -ic -nc