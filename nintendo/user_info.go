package nintendo

import (
	"github.com/einsjustinn/nintendo-api/utils"
	"net/http"
)

const (
	nintendoAccountBaseUrl = "https://api.accounts.nintendo.com/2.0.0"
	nintendoBaseUrl        = "https://api-lp1.znc.srv.nintendo.net"
)

func GetUserInfo(token string) (*UserResponse, error) {
	request, err := http.NewRequest("GET", nintendoAccountBaseUrl+"/users/me", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	return utils.DoReq[UserResponse](request)
}
