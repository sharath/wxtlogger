package WeatherStation

import "github.com/tarm/serial"

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
	wxt.port, _ = serial.OpenPort(&serial.Config{Name: location, Baud: baud})
	return wxt
}

func (wxt *WeatherStation) write(command string) {
	wxt.port.Write([]byte(command))
}

func (wxt *WeatherStation) configure() {
	// wxt_comms_configure
		// send set_comm
	wxt.port.Write([]byte("0R0\r\n"))
	// wxt_wind_configure
		// send set_wind_conf
		// set wind_units in WXTResponse
		// send set_wind_parameters
	// wxt_ptu_configure
		// send set_ptu_conf
		// set pressure_units in WXTResponse
		// set temp_units in WXTResponse
		// send set_ptu_parameters
	// wxt_supervisor_configure
		// send set_super_conf
		// send set_super_parameters
}
