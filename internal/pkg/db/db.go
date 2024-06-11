package db

import (
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	INFLUX_TOKEN = "INFLUX_TOKEN"
	INFLUX_ORG   = "myorg"
	INFLUX_URL   = "http://localhost:8086"
)

type Client struct {
	client      influxdb2.Client
	writeAPI    api.WriteAPI
	Bucket      string
	Measurement string
}

func (c *Client) Connect() {
	c.client = influxdb2.NewClient(INFLUX_URL, os.Getenv(INFLUX_TOKEN))
}

func (c *Client) Close() {
	c.writeAPI.Flush()
	c.client.Close()
}

func (c *Client) Write(fields map[string]interface{}) {
	c.writeAPI = c.client.WriteAPI(INFLUX_ORG, c.Bucket)
	p := influxdb2.NewPoint(c.Measurement, nil, fields, time.Now())
	c.writeAPI.WritePoint(p)
}
