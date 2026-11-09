package main

import (
	"flag"
	"netops-agent/internal/collector"
	"netops-agent/internal/identity"
	"netops-agent/internal/logging"
	"netops-agent/internal/sender"
	"time"
)

const AgentVersion = "0.0.1-ALPHA"

func main() {
	apiURL := flag.String("api", "http://localhost:8080", "URL racine de l'API")
	verbose := flag.Bool("v", false, "Activer les logs debugs")
	logFile := flag.Bool("log-file", true, "Activer la sortie de log sur le fichier log temporaire")
	flag.Parse()

	logger, err := logging.NewLogger(*verbose, *logFile)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("démarrage de NetOps Agent v%s", AgentVersion)
	deviceID, err := identity.GetDeviceID()
	if err != nil {
		logger.Error("échec lors de la récupération du device-id: %w", err)
	}
	logger.Info("machine identifiée: %s", deviceID)

	httpSender := sender.New(*apiURL)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	runJob(deviceID, httpSender, logger)
	for range ticker.C {
		runJob(deviceID, httpSender, logger)
	}
}

func runJob(deviceID string, s *sender.Sender, log *logging.Logger) {
	log.Debug("début de la collecte des métriques...")

	payload, err := collector.Collect(deviceID, AgentVersion)
	if err != nil {
		log.Error("erreur lors de la collecte système: %v", err)
		return
	}

	if err := s.SendHeartbeat(payload); err != nil {
		log.Error("échec d'envoie heartbeat: %v", err)
		return
	}

	log.Info("Heartbeat transmis avec succès à l'API. (CPU: %.2f%%, RAM: %.2ff)", payload.Metrics.CPU.UsagePercent, payload.Metrics.Memory.UsagePercent)
}
