package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of requests handled",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(RequestsTotal)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}
