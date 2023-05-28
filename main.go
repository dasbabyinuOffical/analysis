package main

import (
	"analysis/biz/hot"
	"analysis/config"
	"fmt"
	"time"
)

func main() {
	config.Init()
	ch := time.Tick(time.Minute)
	for {
		select {
		case <-ch:
			fmt.Println("hot.SyncHotPairs")
			err := hot.SyncHotPairs()
			if err != nil {
				fmt.Println("hot.SyncHotPairs,err: ", err.Error())
			}
		}
	}
}
