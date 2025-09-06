package nintendo

import (
	"net/http"

	"github.com/einsjustinn/nintendo-api/utils"
)

func GetShowSelf(webServiceAccessToken string) (*ShowSelfResponse, error) {
	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v4/User/ShowSelf", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+webServiceAccessToken)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	return utils.DoReq[ShowSelfResponse](request)
}
