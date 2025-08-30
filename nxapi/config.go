package nxapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetConfig() (*ConfigResponse, error) {
	resp, err := http.Get(nxApiBaseUrlZnca + "/api/znca/config")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var config ConfigResponse
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ConfigResponse: %w", err)
	}
	return &config, nil
}
