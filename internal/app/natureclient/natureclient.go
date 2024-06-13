package natureclient

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hn-11/atmo-client/internal/pkg/db"
	"github.com/tenntenn/natureremo"
)

const (
	NATURE_TOKEN     = "NATURE_TOKEN"
	BUCKET_NAME      = "nature"
	MEASUREMENT_NAME = "nature"
)

type Client struct {
	naCli *natureremo.Client
	ctx   *context.Context
	dbCli *db.Client
}

func Init() *Client {
	naCli := natureremo.NewClient(os.Getenv(NATURE_TOKEN))
	ctx := context.Background()
	dbClient := db.Client{Bucket: BUCKET_NAME, Measurement: MEASUREMENT_NAME}
	dbClient.Connect()

	return &Client{naCli, &ctx, &dbClient}
}

func (c *Client) Start() {
	for {
		ds, err := c.naCli.DeviceService.GetAll(*c.ctx)
		if err != nil {
			log.Print(err)
			time.Sleep(1 * time.Minute)
		} else {
			for _, d := range ds {
				hu := d.NewestEvents[natureremo.SensorTypeHumidity]
				il := d.NewestEvents[natureremo.SensorTypeIllumination]
				mo := d.NewestEvents[natureremo.SensorTypeMovement]
				te := d.NewestEvents[natureremo.SensorTypeTemperature]

				c.dbCli.Write(map[string]interface{}{"Hu": hu.Value}, hu.CreatedAt)
				c.dbCli.Write(map[string]interface{}{"Il": il.Value}, il.CreatedAt)
				c.dbCli.Write(map[string]interface{}{"Mo": mo.Value}, mo.CreatedAt)
				c.dbCli.Write(map[string]interface{}{"Te": te.Value}, te.CreatedAt)
			}

			d := time.Until(c.naCli.LastRateLimit.Reset)
			time.Sleep(d / time.Duration(c.naCli.LastRateLimit.Remaining))
		}
	}
}

func (c *Client) Close() {
	c.dbCli.Close()
}
