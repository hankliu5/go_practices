package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	myDB           = "mydb"
	instanceSize   = 10000
	metricSize     = 10
	startTime      = int64(1530316800)
	sampleInterval = 6 * 10
	dayInMillis    = 24 * 60 * 60
	batchSize      = 500
)

func main() {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})

	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wg := sync.WaitGroup{}

	start := time.Now()
	for timestamp := startTime; timestamp < (startTime + dayInMillis); timestamp += sampleInterval {
		wg.Add(1)
		go func(timestamp int64) {
			defer wg.Done()
			bp, err := client.NewBatchPoints(client.BatchPointsConfig{
				Database: myDB,
			})
			if err != nil {
				log.Fatal(err)
			}
			for i := 1; i <= instanceSize; i++ {
				instanceName := fmt.Sprintf("%s%d", "instance_", i)
				tags := map[string]string{"instance": instanceName}
				fields := map[string]interface{}{
					"metric_1":  rand.Float64(),
					"metric_2":  rand.Float64(),
					"metric_3":  rand.Float64(),
					"metric_4":  rand.Float64(),
					"metric_5":  rand.Float64(),
					"metric_6":  rand.Float64(),
					"metric_7":  rand.Float64(),
					"metric_8":  rand.Float64(),
					"metric_9":  rand.Float64(),
					"metric_10": rand.Float64(),
				}
				pt, err := client.NewPoint("test", tags, fields, time.Unix(timestamp, 0))
				if err != nil {
					log.Fatal(err)
				}

				bp.AddPoint(pt)
			}
			// Write the batch
			if err := c.Write(bp); err != nil {
				log.Fatal(err)
			}
		}(timestamp)
		wg.Wait()
	}

	elapsed := time.Since(start)
	fmt.Printf("%s: %d", "elapsed time", elapsed)
	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
