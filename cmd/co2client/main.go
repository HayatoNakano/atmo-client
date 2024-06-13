package main

import (
	"log"

	"github.com/hn-11/atmo-client/internal/app/co2client"
)

func main() {
	c, err := co2client.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	err = c.Start()
	if err != nil {
		c.Close()
		log.Fatal(err)
	}
}
