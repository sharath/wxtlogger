package Barometer

import (
	"github.com/tarm/serial"
	"time"
)

type Device struct {
	Location     string `json:"location"`
	Baud         int    `json:"baud,string"`
	SerialNumber string
	Response     Response
	port         *serial.Port
}

func (dqt *Device) write(command string) {
	dqt.port.Write([]byte(command))
}

func (dqt *Device) read() string {
	buf := make([]byte, 1024)
	dqt.port.Read(buf)
	return string(buf)
}

func (dqt *Device) Configure() {
	dqt.port,_ = serial.OpenPort(&serial.Config{Name: dqt.Location, Baud: dqt.Baud, ReadTimeout: time.Millisecond*40})
	delay := time.Millisecond * 200
	time.Sleep(delay)
	dqt.port.Flush()
	dqt.write("")
	time.Sleep(delay)

}