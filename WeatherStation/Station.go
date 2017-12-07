package WeatherStation

import (
	"github.com/tarm/serial"
	"time"
	"strings"
)

type Station struct {
	Location string `json:"location"`
	Baud     int    `json:"baud,string"`
	Response Response
	port     *serial.Port
}

func (wxt *Station) write(command string) {
	wxt.port.Write([]byte(command))
}

func (wxt *Station) read() string {
	buf := make([]byte, 1024)
	wxt.port.Read(buf)
	return string(buf)
}

func (wxt *Station) Configure() {
	wxt.port, _ = serial.OpenPort(&serial.Config{Name: wxt.Location, Baud: wxt.Baud, ReadTimeout: time.Second})
	delay := time.Millisecond * 200
	time.Sleep(delay)
	wxt.port.Flush()
	wxt.write("0XU,M=P,C=3,B=4800,L=25\r\n")
	time.Sleep(delay)
	wxt.write("0WU,I=1,A=1,U=M,D=0,N=W,F=1\r\n")
	time.Sleep(delay)
	wxt.write("0WU,R=0100100001001000\r\n")
	time.Sleep(delay)
	wxt.write("0TU,I=1,P=H,T=F\r\n")
	time.Sleep(delay)
	wxt.write("0TU,R=1101000011010000\r\n")
	time.Sleep(delay)
	wxt.write("0SU,S=N,H=Y,I=5\r\n")
	time.Sleep(delay)
	wxt.write("0SU,R=1111000000000000\r\n")
	time.Sleep(delay)
	wxt.port.Flush()

	wxt.Response.WindUnits = "m/s"
	wxt.Response.PressureUnits = "hPa"
	wxt.Response.TempUnits = "F"

	wxt.Response.Time = time.Now()
	wxt.write("0R0\r\n")
}


func (wxt *Station) UpdateResponse() {
	resp := strings.Split(strings.TrimSpace(wxt.read()),",")
	wxt.Response.Parse(resp[1:])
	wxt.Response.Time = time.Now()
	wxt.write("0R0\r\n")
}