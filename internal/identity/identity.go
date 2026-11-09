package identity

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

// GetDeviceID récupère l'identifiant unique natif de la machine (HostID/Machine-ID).
//
// En cas d'échec de lecture du système, elle retourne une erreur.
func GetDeviceID() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("échec lors de la collecte des infos hôtes : %w", err)
	}

	return info.HostID, nil
}
