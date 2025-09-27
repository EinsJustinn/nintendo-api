package nintendo

import (
	"net/http"

	"github.com/einsjustinn/nintendo-api/utils"
)

func GetFriendList(webServiceAccessToken string, version string) (*FriendListResponse, error) {

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v3/Friend/List", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Platform", version)
	request.Header.Set("X-ProductVersion", version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer "+webServiceAccessToken)
	return utils.DoReq[FriendListResponse](request)
}
