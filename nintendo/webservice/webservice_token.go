package webservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"nintendo-api/nxapi"
	"nintendo-api/utils"
)

func GetWebServiceToken(webServiceId int64, fGenerationResponse nxapi.FGenerationResponse, accessToken string) (*GetWebServiceTokenResponse, error) {

	config, err := nxapi.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	webServiceTokenRequestJson, err := json.Marshal(GetWebServiceTokenRequest{
		Parameter: GetWebServiceTokenRequestParameter{
			F:                 fGenerationResponse.F,
			Timestamp:         fGenerationResponse.Timestamp,
			RequestId:         fGenerationResponse.RequestId,
			Id:                webServiceId,
			RegistrationToken: accessToken,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v2/Game/GetWebServiceToken", bytes.NewBuffer(webServiceTokenRequestJson))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("com.nintendo.znca/%s (Android/14)", config.NsoVersion))
	return utils.DoReq[GetWebServiceTokenResponse](request)
}
