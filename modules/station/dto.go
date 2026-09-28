package station

import "encoding/json"

type Station struct {
	Id   json.Number `json:"id"`
	Name string      `json:"name"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type StationEnvelope struct {
	Data []Station `json:"data"`
}
