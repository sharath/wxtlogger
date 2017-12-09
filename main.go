package main

import (
	wxt "github.com/sharath/infralogger/WeatherStation"
	"time"
	"fmt"
	"os"
	"path"
	"bufio"
)

type Sampler struct {
	Writer *bufio.Writer
	DataFile *os.File
	Index uint64
	Station *wxt.Device
}

func InitializeSampler(station *wxt.Device) *Sampler {
	s := new(Sampler)
	s.NewFile()
	s.Index = 0
	s. Station = station
	return s
}

func (s *Sampler) NewFile() {
	s.DataFile.Close()

	folder := time.Now().Format("data-20060102")
	filename := fmt.Sprintf("WX%d-%s",s.Station.Id,time.Now().Format("20060102-030405.txt"))
	os.Mkdir(folder, 0777)
	file, _ := os.Create(path.Join(folder, filename))
	s.Writer = bufio.NewWriter(file)

	header := fmt.Sprintf("Model Number: WXT-510\n")
	header += fmt.Sprintf("Sample rate: 1 (Hz)\n")
	header += fmt.Sprintf("Wind speed units: m/s\n")
	header += fmt.Sprintf("Pressure units: hPa\n")
	header += fmt.Sprintf("Temperature units: F\n\n")
	header += fmt.Sprintf("Index, Hour, Minute, Second, Direction, Speed, Temp, Humidity, Pressure\n")
	header += fmt.Sprintf("_______________________________________________________________________\n\n")
	fmt.Print(header)
	fmt.Fprint(s.Writer, header)
	s.Writer.Flush()
}

func (s *Sampler) poll() {
	if time.Now().Day() != s.Station.Response.Time.Day() {
		s.NewFile()
	}
	s.Station.UpdateResponse()
	line := fmt.Sprintf("%d, %s, %d, %.1f, %.1f, %.1f, %.1f\n",
		s.Index, s.Station.Response.Time.Format("15, 04, 05.999999"),
		s.Station.Response.WindDir, s.Station.Response.WindAvg,
		s.Station.Response.Temp, s.Station.Response.Humidity,
		s.Station.Response.Pressure)
	fmt.Print(line)
	fmt.Fprint(s.Writer, line)
	s.Writer.Flush()
	s.Index++
}


func main() {
	station := wxt.Load(os.Args[1])[0]

	station.Configure()
	sampler := InitializeSampler(&station)

	for range time.NewTicker(time.Second).C {
		sampler.poll()
	}
	select {}
}