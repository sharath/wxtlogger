package main

import (
	"fmt"
	"github.com/tarm/serial"
	"time"
)


func main() {
	fmt.Println("starting wxtlogger")
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 4800}
	s, _ := serial.OpenPort(c)
	time.Sleep(2)
	_, _ = s.Write([]byte("0R0\r\n"))
	time.Sleep(2)
	buf := make([]byte, 128)
	_, _ = s.Read(buf)
	time.Sleep(2)
	fmt.Println(string(buf))
}