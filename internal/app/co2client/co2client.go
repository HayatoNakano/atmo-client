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

func correct(raw values) (corrected *values) {
	// https://x.com/rynan4818/status/1627089985454366720
	tmp := raw.tmp - 4.5
	hum := raw.hum * raw.tmp * (tmp + 237.3) / (tmp * (raw.tmp + 237.3))
	return &values{raw.co2, hum, tmp}
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

		v, err := c.read(string(line))
		if err != nil {
			log.Fatal(err)
		}

		correct(*v)
		time.Sleep(1 * time.Second)
	}

}

func (c *client) open() (*serial.Port, error) {
	return serial.OpenPort(&serial.Config{Name: c.dev, Baud: 115200})
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
