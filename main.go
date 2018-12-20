package main

import (
	"github.com/spf13/viper"
	"scrapy/config"
	"scrapy/logic"
	"time"
)

func main() {
	config.InitConfig()
	logic.Init()
	logic.Check()
	interval := viper.GetDuration("watchInterval")
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C: // 检测定时器
			logic.Check()
		}
	}
}
