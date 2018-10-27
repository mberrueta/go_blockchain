package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
)

var (
	Hits      *prometheus.CounterVec
	Failures  *prometheus.CounterVec
	Duration  prometheus.ObserverVec
	Executing *prometheus.GaugeVec
)

// https://prometheus.io/docs/practices/naming/
func SetupMetrics() {
	log.Println("Setting up metrics")

	setDuration()
	setHits()
	setFailures()
	setExecuting()
}

func setDuration() {

	Duration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "blockchain_call_Duration_seconds",
			Help:    "Time taken to execute",
			Buckets: prometheus.LinearBuckets(.01, .05, 10),
		},
		[]string{"service", "method", "op"}).MustCurryWith(labels())

	initLabels(Duration, serviceLabels())
}

func setHits() {
	Hits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "blockchain_call_Hits",
			Help: "How many Hits has",
		},
		[]string{"service", "method", "op"}).MustCurryWith(labels())

	initLabels(Hits, serviceLabels())
}

func setFailures() {
	Failures = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "blockchain_call_Failures",
			Help: "How many Failures has",
		},
		[]string{"service", "method", "op"}).MustCurryWith(labels())

	initLabels(Failures, serviceLabels())
}

func setExecuting() {
	Executing = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "blockchain_call_Executing",
			Help: "How many Executing has",
		},
		[]string{"service", "method", "op"}).MustCurryWith(labels())

	initLabels(Executing, serviceLabels())
}

func labels() prometheus.Labels {
	return prometheus.Labels{
		"service": "blockchain",
	}
}

func serviceLabels() []prometheus.Labels {
	return []prometheus.Labels{
		{"method": "add", "op": "get"},
		{"method": "list", "op": "get"},
	}
}

func initLabels(m interface{}, l []prometheus.Labels) {
	// TODO: how to avoid repeated for?
	switch t := m.(type) {
	case prometheus.CounterVec:
		for _, labels := range l {
			t.With(labels)
		}
	case prometheus.HistogramVec:
		for _, labels := range l {
			t.With(labels)
		}
	case prometheus.GaugeVec:
		for _, labels := range l {
			t.With(labels)
		}
	}
}
