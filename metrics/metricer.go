package metrics

import "github.com/prometheus/client_golang/prometheus"

func Gauge(method string, fn func() (interface{}, error)) (interface{}, error) {
	labels := prometheus.Labels{"method": method, "op": ""}
	timer := prometheus.NewTimer(Duration.With(labels))
	defer timer.ObserveDuration()
	Executing.With(labels).Inc()
	defer Executing.With(labels).Dec()
	defer Hits.With(labels).Inc()

	i, err := fn()
	if err != nil {
		Failures.With(labels).Inc()
		return nil, err
	}

	return i, nil
}
