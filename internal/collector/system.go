package collector

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/host"
)

func collectIdentity(agentVersion string) (Identity, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return Identity{}, fmt.Errorf("échec info hôte : %w", err)
	}

	platformStr := hostInfo.Platform
	if hostInfo.PlatformVersion != "" {
		platformStr = fmt.Sprintf("%s %s", hostInfo.Platform, hostInfo.PlatformVersion)
	} else {
		platformStr = fmt.Sprintf("%s (kernel %s)", hostInfo.Platform, hostInfo.KernelVersion)
	}

	return Identity{
		Hostname:      hostInfo.Hostname,
		OS:            hostInfo.OS,
		Platform:      platformStr,
		Arch:          runtime.GOARCH,
		IPAddress:     getLocalIP(),
		AgentVersion:  agentVersion,
		UptimeSeconds: hostInfo.Uptime,
	}, nil
}
