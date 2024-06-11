package co2client

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

type values struct {
	co2 float64
	hum float64
	tmp float64
}

type client struct {
	dev string
	re  *regexp.Regexp
}

func (c *client) open() (*serial.Port, error) {
	return serial.OpenPort(&serial.Config{Name: c.dev, Baud: 115200})
}

func (c *client) Start() {
	s, err := c.open()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		c.read(string(line))
		time.Sleep(1 * time.Second)
	}

}

func (c *client) read(line string) (value *values, err error) {
	matches := c.re.FindStringSubmatch(line)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid format: %s", line)
	}

	var result [3]float64
	for i, m := range matches[1:4] {
		result[i], err = strconv.ParseFloat(m, 64)
		if err != nil {
			return nil, err
		}
	}

	return &values{result[0], result[1], result[2]}, nil
}
