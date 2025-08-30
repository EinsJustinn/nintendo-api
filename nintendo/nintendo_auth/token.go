package nintendo_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetToken(sessionToken string) (TokenResponse, error) {
	tokenRequest := tokenRequest{}
	tokenRequest.ClientId = nintendoClientId
	tokenRequest.SessionToken = sessionToken
	tokenRequest.GrantType = nintendoGrandType

	tokenRequestJson, err := json.Marshal(tokenRequest)
	if err != nil {
		return TokenResponse{}, err
	}

	request, err := http.NewRequest("POST", nintendoConnectBaseUrl+"/api/token", bytes.NewBuffer(tokenRequestJson))
	if err != nil {
		return TokenResponse{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return TokenResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return TokenResponse{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return TokenResponse{}, err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return TokenResponse{}, err
	}

	return tokenResponse, nil
}
