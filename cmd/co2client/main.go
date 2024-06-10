package main

import (
	"log"

	"github.com/hn-11/atmo-client/internal/app/co2client"
)

func main() {
	s, err := co2client.Open()
	if err != nil {
		log.Fatal(err)
	}

}
