package main

import (
	"log"

	"github.com/hn-11/atmo-client/internal/app/co2client"
)

func main() {
	c := co2client.Client{}

	err := c.Connect()
	defer c.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = c.Start()
	if err != nil {
		c.Close()
		log.Fatal(err)
	}
}
