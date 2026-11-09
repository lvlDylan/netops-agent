package collector

import (
	"fmt"
	"time"
)

func Collect(deviceID string, agentVersion string) (*Payload, error) {
	identity, err := collectIdentity(agentVersion)
	if err != nil {
		return nil, fmt.Errorf("échec collecte identitée: %w", err)
	}

	cpuMetrics, err := collectCPU()
	if err != nil {
		return nil, fmt.Errorf("échec collecte cpu: %w", err)
	}

	memMetrics, err := collectMemory()
	if err != nil {
		return nil, fmt.Errorf("échec collecte mem: %w", err)
	}

	diskMetrics, err := collectDisk()
	if err != nil {
		return nil, fmt.Errorf("échec collecte disque: %w", err)
	}

	return &Payload{
		DeviceID: deviceID,
		SendAt:   time.Now().UTC(),
		Identity: identity,
		Metrics: Metrics{
			CPU:    cpuMetrics,
			Memory: memMetrics,
			Disk:   diskMetrics,
		},
	}, nil
}
