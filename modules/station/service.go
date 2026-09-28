package station

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arifwp/mrt-schedules/common/client"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() ([]StationResponse, error) {
	url := "https://beweb-dev.jakartamrt.co.id/middleware/api/datum"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var env StationEnvelope
	if err := json.Unmarshal(byteResponse, &env); err != nil {
		return nil, fmt.Errorf("decode stations: %w", err)
	}

	response := make([]StationResponse, 0, len(env.Data))
	for _, item := range env.Data {
		response = append(response, StationResponse{Id: item.Id.String(), Name: item.Name})
	}
	return response, nil
}
