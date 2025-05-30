package main

import (
	"context"
	"fmt"
	"time"

	"logging-tracing-observality/logutil"
	"logging-tracing-observality/metrics"
	"logging-tracing-observality/tracing"
)

func main() {
	logutil.LogInfo("Service is starting...")
	metrics.InitMetrics()

	// Init Jaeger
	tp, err := tracing.InitJaeger("http://localhost:14268/api/traces")
	if err != nil {
		fmt.Println("Jaeger init failed:", err)
		return
	}
	defer tp.Shutdown(context.Background())

	// Simulate request
	for i := 0; i < 3; i++ {
		metrics.RequestsTotal.Inc()
		tracing.TraceJaeger(context.Background(), fmt.Sprintf("operation-%d", i))
		time.Sleep(1 * time.Second)
	}

	logutil.LogError("Finished with errors.")
}
