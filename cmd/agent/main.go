package main

import (
	"time"

	"github.com/AbramovArseniy/RuntimeMetrics/tree/increment1/internal/agent"
)

const (
	pollRuntimeMetricsInterval = 2 * time.Second
	reportInterval             = 10 * time.Second
)

func main() {
	go agent.Schedule(agent.CollectRuntimeMetrics, pollRuntimeMetricsInterval)
	go agent.Schedule(agent.SendMetrics, reportInterval)
}
