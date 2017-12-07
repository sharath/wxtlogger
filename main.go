package main

import (
	wxt "github.com/sharath/infralogger/WeatherStation"
	"time"
	"fmt"
	"os"
	"path"
	"bufio"
)

func sampleW(x *wxt.Device, i int) {
	var index uint64 = 0

	folder := time.Now().Format("data-20060102")
	filename := fmt.Sprintf("WX%d-%s",i,time.Now().Format("20060102-030405.txt"))
	os.Mkdir(folder, 0777)
	file, _ := os.Create(path.Join(folder, filename))
	w := bufio.NewWriter(file)

	header := fmt.Sprintf("Model Number: WXT-510\n")
	header += fmt.Sprintf("Sample rate: 1 (Hz)\n")
	header += fmt.Sprintf("Wind speed units: m/s\n")
	header += fmt.Sprintf("Pressure units: hPa\n")
	header += fmt.Sprintf("Temperature units: F\n\n")
	header += fmt.Sprintf("Index, Hour, Minute, Second, Direction, Speed, Temp, Humidity, Pressure\n")
	header += fmt.Sprintf("_______________________________________________________________________\n\n")
	fmt.Print(header)
	fmt.Fprint(w, header)
	w.Flush()

	for {
		if time.Now().Second() != x.Response.Time.Second() {
			if time.Now().Day() != x.Response.Time.Day() {
				file.Close()
				folder = time.Now().Format("data-20060102")
				filename = fmt.Sprintf("WX%d-%s",i,time.Now().Format("20060102-030405.txt"))
				os.Mkdir(folder, 0777)
				file, _ = os.Create(path.Join(folder, filename))
				w = bufio.NewWriter(file)
				fmt.Print(header)
				fmt.Fprint(w, header)
				w.Flush()
			}
			x.UpdateResponse()
			line := fmt.Sprintf("%d, %s, %d, %.1f, %.1f, %.1f, %.1f\n",
				index, x.Response.Time.Format("15, 04, 05.999999"),
					x.Response.WindDir, x.Response.WindAvg,
						x.Response.Temp, x.Response.Humidity,
							x.Response.Pressure)
			fmt.Print(line)
			fmt.Fprint(w, line)
			w.Flush()
			index++
		}
	}
}


func main() {
	wStations := wxt.Load("wStations.json")
	for i, w := range wStations {
		w.Configure()
		go sampleW(&w, i+1)
	}
	select {}
}
