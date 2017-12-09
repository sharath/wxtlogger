package main

import (
	wxt "github.com/sharath/wxtlogger/WeatherStation"
	"time"
	"fmt"
	"os"
	"path"
	"bufio"
)

type Sampler struct {
	Writer   *bufio.Writer
	Index    uint64
	Station  *wxt.Device
	Datafile *os.File
}

func InitializeSampler(station *wxt.Device) *Sampler {
	s := new(Sampler)
	s.Index = 0
	s.Station = station
	s.Datafile, _ = os.Create("dummyfile")
	s.NewFile()
	return s
}

func (s *Sampler) NewFile() {
	s.Datafile.Close()

	folder := time.Now().Format("data-20060102")
	filename := fmt.Sprintf("WX%d-%s", s.Station.Id, time.Now().Format("20060102-030405.txt"))
	os.Mkdir(folder, 0777)
	s.Datafile, _ = os.Create(path.Join(folder, filename))
	s.Writer = bufio.NewWriter(s.Datafile)
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
	stations := wxt.Load(os.Args[1])
	samplers := [2]*Sampler{}
	for i := 0; i < len(stations); i++ {
		stations[i].Configure()
		samplers[i] = InitializeSampler(&stations[i])
	}
	for _ = range time.NewTicker(time.Second).C {
		for i := 0; i < len(samplers); i++ {
			go samplers[i].poll()
		}
	}
	select {}
}
