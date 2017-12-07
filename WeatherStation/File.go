package WeatherStation

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type File struct {
	Wxt []Station
}

func Load(file string) *File {
	cf := new(File)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.Wxt)
	return cf
}