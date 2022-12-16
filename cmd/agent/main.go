package main

import (
	"time"

	"github.com/AbramovArseniy/RuntimeMetrics/internal/agent"
)

const (
	pollRuntimeMetricsInterval = 2 * time.Second
	reportInterval             = 10 * time.Second
)

func main() {
	go agent.Repeat(agent.CollectRuntimeMetrics, pollRuntimeMetricsInterval)
	go agent.Repeat(agent.SendAllMetrics, reportInterval)
}
