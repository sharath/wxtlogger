package WeatherStation

import (
	"time"
	"strings"
	"strconv"
)

type Response struct {
	Time time.Time
	WindAvg float64
	WindDir int
	Temp float64
	Humidity float64
	Pressure float64
}

func cut(str string, start string, end string) string {
	return strings.Replace(strings.Replace(str, start, "", -1), end, "", -1)
}

func (storage *Response) Parse(resp []string) {
	for _, p := range resp {
		t := p
		t = strings.Replace(t, "\x00", "", -1)
		t = strings.Replace(t, "\r\n", "", -1)
		if strings.Contains(t, "Sm=") {
			storage.WindAvg, _ = strconv.ParseFloat(cut(t, "Sm=", "M"), 64)
			continue
		}
		if strings.Contains(t, "Dm=") {
			storage.WindDir, _ = strconv.Atoi(cut(t, "Dm=", "D"))
			continue
		}
		if strings.Contains(t, "Ta=") {
			storage.Temp, _ = strconv.ParseFloat(cut(t, "Ta=", "F"), 64)
			continue
		}
		if strings.Contains(t, "Ua=") {
			storage.Humidity, _ = strconv.ParseFloat(cut(t, "Ua=", "P"), 64)
			continue
		}
		if strings.Contains(t, "Pa=") {
			storage.Pressure, _ = strconv.ParseFloat(cut(t, "Pa=", "H"), 64)
			continue
		}
	}
}