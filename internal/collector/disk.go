package collector

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/disk"
)

func collectDisk() (DiskMetrics, error) {
	mountPoint := "/"
	if runtime.GOOS == "windows" {
		mountPoint = "C:"
	}

	diskStat, err := disk.Usage(mountPoint)
	if err != nil {
		return DiskMetrics{}, fmt.Errorf("échec disque sur '%s' : %w", mountPoint, err)
	}

	return DiskMetrics{
		UsagePercent: roundFloat(diskStat.UsedPercent, 2),
		UsedBytes:    diskStat.Used,
		TotalBytes:   diskStat.Total,
	}, nil
}
