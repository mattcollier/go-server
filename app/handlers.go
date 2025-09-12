package app

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCount atomic.Uint64

type metrics struct {
	requests prometheus.GaugeFunc
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		requests: prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Namespace: "my_go_server",
			Name:      "request_counter",
			Help:      "Number of HTTP requests.",
		}, func() float64 {
			fmt.Println("Metrics being scraped.")
			return float64(requestCount.Load())
		}),
	}
	reg.MustRegister(m.requests)
	return m
}

// count is a middleware that increments a global counter.
func countWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/metrics" {
			requestCount.Add(1)
		}
		next.ServeHTTP(w, r)
	})
}

func Routes() {
	reg := prometheus.NewRegistry()
	NewMetrics(reg)

	mux := http.NewServeMux()

	promHander := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	mux.Handle("GET /metrics", promHander)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("requestCount: %d\n", requestCount.Load())
		fmt.Fprint(w, "hello!!!")
	})

	mux.HandleFunc("GET /replicasets", func(w http.ResponseWriter, r *http.Request) {
		replicasets, _ := GetReplicasets()
		fmt.Fprint(w, replicasets)
	})

	// count requests to all handlers
	handler := countWrapper(mux)

	// second arg is a "multiplexer", nil is default
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}
