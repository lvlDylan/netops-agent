package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"netops-agent/internal/collector"
	"time"
)

type Sender struct {
	apiURL     string
	httpClient *http.Client
}

func New(apiURL string) *Sender {
	return &Sender{
		apiURL: apiURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *Sender) SendHeartbeat(payload *collector.Payload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("échec de la sérialisation JSON: %w", err)
	}

	endpoint := fmt.Sprintf("%s/api/report", s.apiURL)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("échec de la création de la requête HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("Netops-Agent/%s", payload.Identity.AgentVersion))

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("échec de l'envoie de la requête HTTP: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("l'api a renvoyé un statut d'erreur: %d", resp.StatusCode)
	}

	return nil
}
