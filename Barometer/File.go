package Barometer

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type File struct {
	bMeters []Device
}

func Load(file string) []Device {
	cf := new(File)
	configFile,_ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.bMeters)
	return cf.bMeters
}