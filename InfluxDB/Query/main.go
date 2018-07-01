package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	myDB           = "mydb"
	instanceSize   = 1000
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

	start := time.Now()
	// queryAllByInstance(c, "test3")
	// queryAllByInstanceAndMetric(c, "test5")
	// queryOneMetricByInstance(c, "test3")
	queryOneMetricByInstanceAndMetric(c, "test5")
	elapsed := time.Since(start)

	fmt.Printf("elapsed time: %d\n", elapsed)
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func queryAllByInstance(c client.Client, measurement string) {
	for i := 1; i <= instanceSize; i++ {
		selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE instance='instance_%d'", measurement, i)
		q := client.NewQuery(selectQuery, myDB, "rfc3339")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
		}
	}
}

func queryAllByInstanceAndMetric(c client.Client, measurement string) {
	for i := 1; i <= instanceSize; i++ {
		for j := 1; j <= metricSize; j++ {
			selectQuery := fmt.Sprintf("SELECT value FROM %s WHERE instance='instance_%d' and metric='metric_%d'", measurement, i, j)
			q := client.NewQuery(selectQuery, myDB, "rfc3339")
			if response, err := c.Query(q); err == nil && response.Error() == nil {
			}
		}
	}
}

func queryOneMetricByInstance(c client.Client, measurement string) {
	for i := 1; i <= instanceSize; i++ {
		selectQuery := fmt.Sprintf("SELECT metric_1 FROM %s WHERE instance='instance_%d'", measurement, i)
		q := client.NewQuery(selectQuery, myDB, "rfc3339")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
		}
	}
}

func queryOneMetricByInstanceAndMetric(c client.Client, measurement string) {
	for i := 1; i <= instanceSize; i++ {
		selectQuery := fmt.Sprintf("SELECT value FROM %s WHERE instance='instance_%d' and metric='metric_1'", measurement, i)
		q := client.NewQuery(selectQuery, myDB, "rfc3339")
		if response, err := c.Query(q); err == nil && response.Error() == nil {
		}
	}
}
