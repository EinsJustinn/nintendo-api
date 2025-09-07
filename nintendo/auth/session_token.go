package auth

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/einsjustinn/nintendo-api/utils"
)

const (
	nintendoConnectBaseUrl = "https://accounts.nintendo.com/connect/1.0.0"
	nintendoGrandType      = "urn:ietf:params:oauth:grant-type:jwt-bearer-session-token"
)

// GetSessionToken valid for 2 years
func GetSessionToken(sessionTokenCode string, sessionTokenCodeVerifier string, client Client) (*SessionTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", client.ID)
	data.Set("session_token_code", sessionTokenCode)
	data.Set("session_token_code_verifier", sessionTokenCodeVerifier)

	request, err := http.NewRequest("POST", nintendoConnectBaseUrl+"/api/session_token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return utils.DoReq[SessionTokenResponse](request)
}
