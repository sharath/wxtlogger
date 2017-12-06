package main

import (
	wxt "wxtlogger/WeatherStation"
	conf "wxtlogger/Config"
	"time"
)

func sample(x *wxt.WeatherStation) {
	for {
		x.UpdateResponse()
		time.Sleep(time.Second)
	}
}


func main() {
	wstations := conf.Load("config.json").Wxt;
	for _, w := range wstations {
		w.Configure()
		go sample(&w)
	}
	select {}
}