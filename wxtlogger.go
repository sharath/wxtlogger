package main

import (
	wxt "wxtlogger/WeatherStation"
	conf "wxtlogger/Config"
	"time"
)

func sample(x *wxt.WeatherStation, i int) {
	for {
		x.UpdateResponse()
		time.Sleep(time.Second)
	}
}


func main() {
	wstations := conf.Load("config.json").Wxt;
	for i, w := range wstations {
		w.Configure()
		go sample(&w, i)
	}
	select {}
}