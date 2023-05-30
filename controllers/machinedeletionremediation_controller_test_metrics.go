package controllers

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	reconcilePerCr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "machinedeletionremediation_cr_call_total",
			Help: "Number of reconcile call per given CR",
		},
		[]string{"custom_resource_name"},
	)
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(reconcilePerCr)
}
