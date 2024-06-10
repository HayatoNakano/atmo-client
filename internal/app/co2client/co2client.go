package co2client

import (
	"github.com/tarm/serial"
)

const dev = "/dev/ttyACM0"

func Open() (*serial.Port, error) {
	c := &serial.Config{Name: dev, Baud: 115200}
	return serial.OpenPort(c)
}

func read(output string) (co2 float64, hum float64, tmp float64, err error) {
	return 1, 1, 1, nil
}
