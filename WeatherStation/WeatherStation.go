package WeatherStation

import (
	"github.com/tarm/serial"
	"time"
)

type WeatherStation struct {
	Location string
	Baud int
	Response WXTResponse
	port *serial.Port
}

func NewWeatherStation(location string, baud int) *WeatherStation {
	wxt := new(WeatherStation)
	wxt.Location = location
	wxt.Baud = baud
	wxt.port, _ = serial.OpenPort(&serial.Config{Name: location, Baud: baud, ReadTimeout: time.Second})
	return wxt
}

func (wxt *WeatherStation) write(command string) {
	wxt.port.Write([]byte(command))
}

func (wxt *WeatherStation) Read() string {
	buf := make([]byte, 1024)
	wxt.port.Read(buf)
	return string(buf)
}

func (wxt *WeatherStation) Configure() {
	// set units
	wxt.Response.WindUnits = "m/s"
	wxt.Response.PressureUnits = "hPa"
	wxt.Response.TempUnits = "F"
	wxt.port.Flush()
	// send set_comm
	wxt.write("0XU,M=P,C=3,B=4800,L=25\r\n")
	// send set_wind_conf
	wxt.write("0WU,I=1,A=1,U=M,D=0,N=W,F=1\r\n")
	// send set_wind_parameters
	wxt.write("0WU,R=0100100001001000\r\n")
	// send set_ptu_conf
	wxt.write("0TU,I=1,P=H,T=F\r\n")
	// send set_ptu_parameters
	wxt.write("0TU,R=1101000011010000\r\n")
	// send set_super_conf
	wxt.write("0SU,S=N,H=Y,I=5\r\n")
	// send set_super_parameters
	wxt.write("0SU,R=1111000000000000\r\n")
	wxt.port.Flush()
}