package main

import (
	"encoding/json"
	"netops-agent/internal/collector"
	"netops-agent/internal/identity"
	"netops-agent/internal/logging"
)

const AgentVersion = "0.0.1-ALPHA"

func main() {
	logger, err := logging.NewLogger(true, false)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("=== DEBUT TEST ===")
	deviceID, err := identity.GetDeviceID()
	if err != nil {
		logger.Error("échec lors de la récupération du device-id: %w", err)
	}
	logger.Info("deviceID identifié: %s", deviceID)

	logger.Info("lancement de la collecte système")
	payload, err := collector.Collect(deviceID, AgentVersion)
	if err != nil {
		logger.Error("échec lors de la collecte système: %w", err)
		return
	}

	jsonBytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		logger.Error("échec lors de la sérialization en json: %w", err)
		return
	}

	logger.Info("Payload: %s", string(jsonBytes))
}
