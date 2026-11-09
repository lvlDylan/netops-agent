package collector

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

func isIgnoredFsType(fsType string) bool {
	ignored := map[string]bool{
		"tmpfs":    true,
		"devtmpfs": true,
		"devfs":    true,
		"iso9660":  true,
		"squashfs": true,
		"overlay":  true,
		"proc":     true,
		"sysfs":    true,
		"cgroup":   true,
		"autofs":   true,
	}

	return ignored[strings.ToLower(fsType)]
}

func collectDisk() ([]DiskMetrics, error) {

	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, fmt.Errorf("échec de lister les partitions: %w", err)
	}

	var disks []DiskMetrics

	for _, p := range partitions {
		if isIgnoredFsType(p.Fstype) {
			continue
		}

		diskStat, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		if diskStat.Total == 0 {
			continue
		}

		disks = append(disks, DiskMetrics{
			MountPoint:   p.Mountpoint,
			UsagePercent: roundFloat(diskStat.UsedPercent, 2),
			UsedBytes:    diskStat.Used,
			TotalBytes:   diskStat.Total,
		})
	}

	if len(disks) == 0 {
		return nil, fmt.Errorf("aucune partition valide trouvée")
	}

	return disks, nil
}
