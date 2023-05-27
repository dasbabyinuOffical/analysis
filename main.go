package main

import (
	"analysis/biz/hot"
	"analysis/config"
)

func main() {
	config.Init()
	hot.SyncHotPairs()
}
