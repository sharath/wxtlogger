package Config

import (
	"wxtlogger/WeatherStation"
	"io/ioutil"
	"encoding/json"
)

type Configuration struct {
	Wxt []WeatherStation.WeatherStation
}

func Load(file string) *Configuration {
	cf := new(Configuration)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.Wxt)
	return cf
}