package Barometer

import (
	"time"
)

type Response struct {
	Time time.Time
	Pressure float64
}

func (storage *Response) Parse(resp []string) {

}