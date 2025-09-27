package webservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	nintendoBaseUrl = "https://api-lp1.znc.srv.nintendo.net"
)

func GetWebServices(webToken string, version string) (ListWebServicesResponse, error) {

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v1/Game/ListWebServices", nil)
	if err != nil {
		return ListWebServicesResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("X-Platform", version)
	request.Header.Set("X-ProductVersion", version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer "+webToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ListWebServicesResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return ListWebServicesResponse{}, fmt.Errorf("status code: %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ListWebServicesResponse{}, fmt.Errorf("failed to read response: %w", err)
	}
	var responseResponse ListWebServicesResponse
	err = json.Unmarshal(body, &responseResponse)
	if err != nil {
		return ListWebServicesResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return responseResponse, nil
}
