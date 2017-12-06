package Logging

import (wxt "wxtlogger/WeatherStation"
	"bufio"
	"fmt"
)

type WxtLog struct {
	 Header string
	 Filename string
	 WStation *wxt.WeatherStation
	 Index uint64
	 curfile bufio.Writer
}

func NewWxtLog(wstation *wxt.WeatherStation) *WxtLog {
	r := new(WxtLog)
	r.WStation = wstation
	r.formHeader()
	return r
}

func (log *WxtLog) WriteLine() {

}

func (log *WxtLog) formHeader() {
	fmt.Sprintf(log.Header, "Model Number: WXT-510\n")
	fmt.Sprintf(log.Header, "Sample rate: 1 (Hz)\n")
	fmt.Sprintf(log.Header, "Wind speed units: %s", log.WStation.Response.WindUnits)
	fmt.Sprintf(log.Header, "Pressure units: %s", log.WStation.Response.PressureUnits)
	fmt.Sprintf(log.Header, "Temperature units: %s", log.WStation.Response.TempUnits)
	fmt.Println(log.Header)
}