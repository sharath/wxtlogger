package WeatherStation

type WeatherStation struct {
	Location string
	Response WXTResponse
}


func (wxt *WeatherStation) configure() {
	// wxt_comms_configure
		// send set_comm
	// wxt_wind_configure
		// send set_wind_conf
		// set wind_units in WXTResponse
		// send set_wind_parameters
	// wxt_ptu_configure
		// send set_ptu_conf

	// wxt_supervisor_configure
}