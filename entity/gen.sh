#!/usr/bin/env bash
xorm reverse mysql "root:root@(127.0.0.1:3306)/lql?charset=utf8" $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm .