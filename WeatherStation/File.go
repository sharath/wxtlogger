package WeatherStation

import (
	"io/ioutil"
	"encoding/json"
)

type File struct {
	wStations []Device
}

func Load(file string) []Device {
	cf := new(File)
	configFile, _ := ioutil.ReadFile(file)
	json.Unmarshal(configFile, &cf.wStations)
	return cf.wStations
}
