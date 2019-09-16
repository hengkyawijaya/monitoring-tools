package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// A counter is a cumulative metric that represents a single monotonically increasing counter whose value can only increase or be reset to zero on restart.
	// For example, you can use a counter to represent the number of requests served, tasks completed, or errors.
	counterPrometheus = promauto.NewCounter(prometheus.CounterOpts{
		Name: "simpleapp_counter",
		Help: "counter demo",
	})

	// A gauge is a metric that represents a single numerical value that can arbitrarily go up and down. Example of gauge includes the heap memory used or CPU usage.
	gaugePrometheus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "simpleapp_gauge",
		Help: "gauge demo",
	})

	// A histogram samples observations (usually things like request durations or response sizes) and counts them in configurable buckets.
	// It also provides a sum of all observed values.
	histogramPrometheus = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "simpleapp_histrogram",
		Help:    "histogram demo",
		Buckets: []float64{0.0001, 0.001, 0.01, 0.1, 1, 10},
	})

	// Similar to a histogram, a summary samples observations (usually things like request durations and response sizes).
	// While it also provides a total count of observations and a sum of all observed values,
	// it calculates configurable quantiles over a sliding time window.
	summaryPrometheus = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "simpleapp_summary",
		Help: "summary demo",
	})

	c *statsd.Client
)

func main() {

	// Create the client
	c, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	// Prefix every metric with the app name
	c.Namespace = "golang."

	// Send global Tags
	c.Tags = append(c.Tags, "indo-east-1a")

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/api/hello", metrics(sayHelloHandler))
	log.Println("server running on port :2112")
	http.ListenAndServe(":2112", nil)
}

func metrics(handler http.HandlerFunc) http.HandlerFunc {
	start := time.Now()

	// Prometheus implementation
	defer counterPrometheus.Add(1)
	defer histogramPrometheus.Observe(time.Since(start).Seconds())

	// Datadog implementation
	defer c.Count("simpleapp_counter", 1, []string{"counter"}, 1)
	defer c.Histogram("simpleapp_histogram", time.Since(start).Seconds(), []string{"histogram"}, 1)
	return handler
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
