package nxapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetConfig() (ConfigResponse, error) {
	resp, err := http.Get(nxApiBaseUrlZnca + "/api/znca/config")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ConfigResponse{}, fmt.Errorf("failed to read TokenResponse: %w", err)
	}
	var config ConfigResponse
	err = json.Unmarshal(body, &config)
	if err != nil {
		return ConfigResponse{}, fmt.Errorf("failed to unmarshal TokenResponse: %w", err)
	}
	return config, nil
}
