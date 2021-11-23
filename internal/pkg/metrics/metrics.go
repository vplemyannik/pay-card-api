package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cardCountTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: "api",
		Name:      "card_notfound_total",
		Help:      "Total count of not found cards",
	})
)

var (
	cardCUDCountTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "api",
		Name:      "card_cud_total",
		Help:      "Total count of cud operations",
	}, []string{"type"})
)

func IncrementNotFoundCardTotalCount() {
	cardCountTotal.Inc()
}

type CudType string

const (
	Create CudType = "Create"
	Update CudType = "Update"
	Delete CudType = "Delete"
)

func IncrementCUDCardOperationsTotalCount(cudType CudType) {
	cardCUDCountTotal.WithLabelValues(string(cudType)).Inc()
}
