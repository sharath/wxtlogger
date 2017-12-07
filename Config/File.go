package Config

import (
	"github.com/json-iterator/go"
	"github.com/sharath/wxtlogger/WeatherStation"
	"io/ioutil"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type File struct {
	Wxt []WeatherStation.Station
}

func Load(file string) *File {
	cf := new(File)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.Wxt)
	return cf
}