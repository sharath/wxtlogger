package WeatherStation

type WXTResponse struct {
	WindUnits string
	PressureUnits string
	TempUnits string
	WindAvg float64
	WindDir int
	Temp float64
	Humidity float64
	Pressure float64
}