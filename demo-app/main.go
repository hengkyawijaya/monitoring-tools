package main

import (
	"net/http"

	"log"
	"math/rand"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "demoapp_counter",
			Help:      "This is demo counter",
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "golang",
			Name:      "demoapp_gauge",
			Help:      "This is demo gauge",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "golang",
			Name:      "demoapp_histogram",
			Help:      "This is demo histogram",
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "golang",
			Name:      "demoapp_summary",
			Help:      "This is demo summary",
		})
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create the client
	c, err := statsd.New("127.0.0.1:2113")
	if err != nil {
		log.Fatal(err)
	}
	// Prefix every metric with the app name
	c.Namespace = "golang."

	// Send the EC2 availability zone as a tag with every metric
	c.Tags = append(c.Tags, "indo-east-1a")
	err = c.Gauge("request.duration", 1.2, nil, 1)

	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			// Prometheus Metric Generator
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			// Datadog Metric Generator
			c.Count("demoapp_counter", rand.Int63()*5, []string{"counter"}, 1)
			c.Gauge("demoapp_gauge", rand.Float64()*15-5, []string{"gauge"}, 1)
			c.Histogram("demoapp_histogram", rand.Float64()*10, []string{"histogram"}, 1)
			c.Set("demoapp_set", generateSet(rand.Intn(1)), []string{"set"}, 1)

			time.Sleep(time.Second)
		}
	}()

	log.Println("server running on port :2113")
	http.ListenAndServe(":2113", nil)
}

func generateSet(index int) string {
	values := []string{"dilan", "milea"}
	return values[index]
}
