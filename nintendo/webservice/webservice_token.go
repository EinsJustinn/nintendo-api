package webservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nintendo-api/nxapi"
)

func GetWebServiceToken(webServiceId int64, fGenerationResponse nxapi.FGenerationResponse, accessToken string) (GetWebServiceTokenResponse, error) {

	config, err := nxapi.GetConfig()
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	webServiceTokenRequest := GetWebServiceTokenRequest{}
	webServiceTokenRequest.Parameter.F = fGenerationResponse.F
	webServiceTokenRequest.Parameter.Timestamp = fGenerationResponse.Timestamp
	webServiceTokenRequest.Parameter.RequestId = fGenerationResponse.RequestId
	webServiceTokenRequest.Parameter.Id = webServiceId
	webServiceTokenRequest.Parameter.RegistrationToken = accessToken

	webServiceTokenRequestJson, err := json.Marshal(webServiceTokenRequest)
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v2/Game/GetWebServiceToken", bytes.NewBuffer(webServiceTokenRequestJson))
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("com.nintendo.znca/%s (Android/14)", config.NsoVersion))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return GetWebServiceTokenResponse{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to read response: %w", err)
	}
	var getWebServiceTokenResponse GetWebServiceTokenResponse
	err = json.Unmarshal(body, &getWebServiceTokenResponse)
	if err != nil {
		return GetWebServiceTokenResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return getWebServiceTokenResponse, nil
}
