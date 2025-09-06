package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"nintendo-api/utils"
)

func GetToken(sessionToken string) (*TokenResponse, error) {
	tokenRequestJson, err := json.Marshal(tokenRequest{
		ClientId:     nintendoClientId,
		SessionToken: sessionToken,
		GrantType:    nintendoGrandType,
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", nintendoConnectBaseUrl+"/api/token", bytes.NewBuffer(tokenRequestJson))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return utils.DoReq[TokenResponse](request)
}
