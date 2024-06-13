package main

import (
	"github.com/hn-11/atmo-client/internal/app/natureclient"
)

func main() {
	c := natureclient.Init()
	defer c.Close()
	c.Start()
}
