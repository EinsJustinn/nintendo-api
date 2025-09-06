package nxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/einsjustinn/nintendo-api/utils"
)

const (
	nxApiBaseUrlZnca = "https://nxapi-znca-api.fancy.org.uk"
	nxApiBaseUrlAuth = "https://nxapi-auth.fancy.org.uk"
	nxApiClientID    = "yfMwOli6yTynugoi3L_7mA"
	nxApiGrantType   = "client_credentials"
	nxApiScope       = "ca:gf"
	userAgent        = "nintendo-api/1.0.0 (+https://github.com/einsjustinn/nintendo-api)"
)

func GetNxApiToken() (*TokenResponse, error) {

	data := url.Values{}
	data.Set("client_id", nxApiClientID)
	data.Set("grant_type", nxApiGrantType)
	data.Set("scope", nxApiScope)

	request, err := http.NewRequest("POST", nxApiBaseUrlAuth+"/api/oauth/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", userAgent)
	return utils.DoReq[TokenResponse](request)
}

func GenerateFToken(token string, naId string, method string, coralUserId int64) (*FGenerationResponse, error) {

	config, err := GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	nxApiToken, err := GetNxApiToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get nx api token: %w", err)
	}

	/*
		Method:
		1 = Coral
		2 = WebService

		CoralUserId only for Webservice
	*/

	jsonBody, err := json.Marshal(fGenerationRequest{
		HashMethod:  method,
		Token:       token,
		NaID:        naId,
		CoralUserId: strconv.FormatInt(coralUserId, 10),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	request, err := http.NewRequest("POST", nxApiBaseUrlZnca+"/api/znca/f", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+nxApiToken.AccessToken)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("X-znca-Platform", configVersion.Platform)
	request.Header.Set("X-znca-Version", configVersion.Version)
	request.Header.Set("X-znca-Client-Version", config.NsoVersion)
	return utils.DoReq[FGenerationResponse](request)
}
