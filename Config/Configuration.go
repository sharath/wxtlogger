package Config

import (
	"github.com/json-iterator/go"
	"github.com/sharath/wxtlogger/WeatherStation"
	"io/ioutil"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Configuration struct {
	Wxt []WeatherStation.Station
}

func Load(file string) *Configuration {
	cf := new(Configuration)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.Wxt)
	return cf
}