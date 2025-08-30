package nintendo_auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	nintendoConnectBaseUrl = "https://accounts.nintendo.com/connect/1.0.0"
	nintendoClientId       = "71b963c1b7b6d119"
	nintendoGrandType      = "urn:ietf:params:oauth:grant-type:jwt-bearer-session-token"
)

// GetSessionToken valid for 2 years
func GetSessionToken(sessionTokenCode string, sessionTokenCodeVerifier string) (SessionTokenResponse, error) {

	data := url.Values{}
	data.Set("client_id", nintendoClientId)
	data.Set("session_token_code", sessionTokenCode)
	data.Set("session_token_code_verifier", sessionTokenCodeVerifier)

	request, err := http.NewRequest("POST", nintendoConnectBaseUrl+"/api/session_token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return SessionTokenResponse{}, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return SessionTokenResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return SessionTokenResponse{}, fmt.Errorf("status code: %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return SessionTokenResponse{}, err
	}
	var sessionAuthResponse SessionTokenResponse
	err = json.Unmarshal(body, &sessionAuthResponse)
	if err != nil {
		return SessionTokenResponse{}, err
	}
	return sessionAuthResponse, nil
}
