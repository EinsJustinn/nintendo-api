package nintendo_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nintendo-api/nintendo"
	"nintendo-api/nxapi"
)

const (
	nintendoBaseUrl = "https://api-lp1.znc.srv.nintendo.net"
)

func Login(userResponse nintendo.UserResponse, idToken string) (LoginResponse, error) {

	config, err := nxapi.GetConfig()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	fToken, err := nxapi.GenerateFToken(idToken, userResponse.Id, "1", 0)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to generate f token: %w", err)
	}

	loginRequest := LoginRequest{}
	loginRequest.Parameter.NaIdToken = idToken
	loginRequest.Parameter.NaBirthday = userResponse.Birthday
	loginRequest.Parameter.NaCountry = userResponse.Country
	loginRequest.Parameter.Language = userResponse.Language
	loginRequest.Parameter.Timestamp = fToken.Timestamp
	loginRequest.Parameter.RequestId = fToken.RequestId
	loginRequest.Parameter.F = fToken.F

	loginRequestJson, err := json.Marshal(loginRequest)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to marshal json: %w", err)
	}
	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v3/Account/Login", bytes.NewBuffer(loginRequestJson))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return LoginResponse{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to read response: %w", err)
	}
	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return loginResponse, nil
}
