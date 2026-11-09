package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func collectMemory() (MemoryMetrics, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return MemoryMetrics{}, fmt.Errorf("échec mémoire: %w", err)
	}

	return MemoryMetrics{
		UsagePercent: roundFloat(vmStat.UsedPercent, 2),
		UsedBytes:    vmStat.Used,
		TotalBytes:   vmStat.Total,
	}, nil
}
