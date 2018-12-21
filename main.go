package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"scrapy/config"
	"scrapy/logic"
	"time"
)

func main() {
	config.InitConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		config.LoadConfig()
		logic.Check()
	})
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
