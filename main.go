package main

import (
	"scrapy/logic"
	"time"
)

func main() {
	logic.Check()
	t := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-t.C: // 检测定时器
			logic.Check()
		}
	}
}
