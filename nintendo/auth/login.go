package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/einsjustinn/nintendo-api/nintendo"
	"github.com/einsjustinn/nintendo-api/utils"
)

const nintendoBaseUrl = "https://api-lp1.znc.srv.nintendo.net"

func Login(idToken string, fToken nintendo.FParam, birthday string, country string, language string, version string) (*LoginResponse, error) {

	loginRequestJson, err := json.Marshal(LoginRequest{
		Parameter: LoginRequestParameter{
			NaIdToken:  idToken,
			NaBirthday: birthday,
			NaCountry:  country,
			Language:   language,
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
	request.Header.Set("X-Platform", version)
	request.Header.Set("X-ProductVersion", version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	return utils.DoReq[LoginResponse](request)
}
