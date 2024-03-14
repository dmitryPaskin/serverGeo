package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HandlerCounter *prometheus.CounterVec
	HandlerTime    *prometheus.HistogramVec
	DatabaseTime   *prometheus.HistogramVec
	CacheTime      *prometheus.HistogramVec
	DadataTime     *prometheus.HistogramVec
}

func New() *Metrics {
	hCounter := NewCounterHandler()
	hTimer := NewTimeHandler()
	dbtimer := NewTimeDB()
	cTimer := NewTimeCache()
	dadataTimer := NewTimeDadata()

	prometheus.MustRegister(hCounter, hTimer, dadataTimer, dbtimer, cTimer)
	return &Metrics{
		HandlerCounter: hCounter,
		HandlerTime:    hTimer,
		DatabaseTime:   dbtimer,
		CacheTime:      cTimer,
		DadataTime:     dadataTimer,
	}
}

func NewTimeHandler() *prometheus.HistogramVec {

	reqDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_sec",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{.1, .2, .5, 1, 2, 5, 10},
		},
		[]string{"method", "path"},
	)

	return reqDuration
}

func NewTimeDB() *prometheus.HistogramVec {

	reqDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "DB_request_duration_sec",
			Help:    "Duration of DB requests",
			Buckets: []float64{.1, .2, .5, 1, 2, 5, 10},
		},
		[]string{"method", "path"},
	)

	return reqDuration
}

func NewTimeCache() *prometheus.HistogramVec {

	reqDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "Cache_request_duration_sec",
			Help:    "Duration of cache requests",
			Buckets: []float64{.1, .2, .5, 1, 2, 5, 10},
		},
		[]string{"method", "path"},
	)

	return reqDuration
}

func NewTimeDadata() *prometheus.HistogramVec {

	reqDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "Dadata_request_duration_sec",
			Help:    "Duration of Dadata service requests",
			Buckets: []float64{.1, .2, .5, 1, 2, 5, 10},
		},
		[]string{"method", "path"},
	)

	return reqDuration
}

func NewCounterHandler() *prometheus.CounterVec {

	reqCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Number of HTTP requests",
		},
		[]string{"method", "path"},
	)

	return reqCounter
}
