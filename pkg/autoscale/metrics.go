package autoscale

import "github.com/armon/go-metrics"

// sendTriggerErrorMetrics is a helper to track autoscaling scale trigger errors. This is done by
// tracking both overall errors, and job specific counters.
func sendTriggerErrorMetrics(job string) {
	metrics.IncrCounter([]string{"autoscale", "trigger", "error"}, 1)
	metrics.IncrCounter([]string{"autoscale", job, "trigger", "error"}, 1)
}

// sendTriggerSuccessMetrics is a helper to track autoscaling scale trigger success. This is done
// by tracking both overall success, and job specific counters.
func sendTriggerSuccessMetrics(job string) {
	metrics.IncrCounter([]string{"autoscale", "trigger", "success"}, 1)
	metrics.IncrCounter([]string{"autoscale", job, "trigger", "success"}, 1)
}
