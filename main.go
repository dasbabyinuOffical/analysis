package main

import (
	"analysis/biz/hot"
	"analysis/config"
	"time"
)

func main() {
	config.Init()
	for {
		time.Sleep(time.Minute)
		hot.SyncHotPairs()
	}
}
