package WeatherStation

import (
	"github.com/tarm/serial"
	"time"
	"strings"
)

type Device struct {
	Location string `json:"location"`
	Baud     int    `json:"baud,string"`
	Id       int    `json:"id,string"`
	Response Response
	port     *serial.Port
}

func (wxt *Device) write(command string) {
	wxt.port.Write([]byte(command))
}

func (wxt *Device) read() string {
	buf := make([]byte, 1024)
	wxt.port.Read(buf)
	return string(buf)
}

func (wxt *Device) Configure() {
	wxt.port, _ = serial.OpenPort(&serial.Config{Name: wxt.Location, Baud: wxt.Baud, ReadTimeout: time.Second})
	commands := []string{"0XU,M=P,C=3,B=9600,L=25\r\n", "0WU,I=1,A=1,U=M,D=0,N=W,F=4\r\n", "0WU,R=0100100001001000\r\n",
		"0TU,I=1,P=H,T=F\r\n", "0TU,R=1101000011010000\r\n", "0SU,S=N,H=Y,I=5\r\n", "0SU,R=1111000000000000\r\n"}
	for _, cmd := range commands {
		time.Sleep(time.Millisecond * 100)
		wxt.write(cmd)
		time.Sleep(time.Millisecond * 100)
	}
	wxt.port.Flush()
	wxt.Response.Time = time.Now()
	wxt.write("0R0\r\n")
}

func (wxt *Device) UpdateResponse() {
	resp := strings.Split(strings.TrimSpace(wxt.read()), ",")
	wxt.Response.Parse(resp)
	wxt.Response.Time = time.Now()
	wxt.write("0R0\r\n")
}
