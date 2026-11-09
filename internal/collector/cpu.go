package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func collectCPU() (CPUMetrics, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	cpuUsage := 0.0

	if err == nil && len(cpuPercent) > 0 {
		cpuUsage = cpuPercent[0]
	}

	cpuCores, err := cpu.Counts(true)
	if err != nil {
		cpuCores = 1
	}

	return CPUMetrics{
		UsagePercent: roundFloat(cpuUsage, 2),
		Cores:        cpuCores,
	}, nil
}
