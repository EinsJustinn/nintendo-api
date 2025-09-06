package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"nintendo-api/nintendo"
	"nintendo-api/nxapi"
	"nintendo-api/utils"
)

const nintendoBaseUrl = "https://api-lp1.znc.srv.nintendo.net"

func Login(userResponse nintendo.UserResponse, idToken string) (*LoginResponse, error) {
	config, err := nxapi.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	fToken, err := nxapi.GenerateFToken(idToken, userResponse.Id, "1", 0)
	if err != nil {
		return nil, fmt.Errorf("failed to generate f token: %w", err)
	}

	loginRequestJson, err := json.Marshal(LoginRequest{
		Parameter: LoginRequestParameter{
			NaIdToken:  idToken,
			NaBirthday: userResponse.Birthday,
			NaCountry:  userResponse.Country,
			Language:   userResponse.Language,
			Timestamp:  fToken.Timestamp,
			RequestId:  fToken.RequestId,
			F:          fToken.F,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal loginrequest json: %w", err)
	}
	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v3/Account/Login", bytes.NewBuffer(loginRequestJson))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	return utils.DoReq[LoginResponse](request)
}
