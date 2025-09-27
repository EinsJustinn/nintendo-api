package webservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/einsjustinn/nintendo-api/nintendo"
	"github.com/einsjustinn/nintendo-api/utils"
)

func GetWebServiceToken(webServiceId int64, fParam nintendo.FParam, accessToken string, version string) (*GetWebServiceTokenResponse, error) {

	webServiceTokenRequestJson, err := json.Marshal(GetWebServiceTokenRequest{
		Parameter: GetWebServiceTokenRequestParameter{
			F:                 fParam.F,
			Timestamp:         fParam.Timestamp,
			RequestId:         fParam.RequestId,
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
	request.Header.Set("X-Platform", version)
	request.Header.Set("X-ProductVersion", version)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("com.nintendo.znca/%s (Android/14)", version))
	return utils.DoReq[GetWebServiceTokenResponse](request)
}
