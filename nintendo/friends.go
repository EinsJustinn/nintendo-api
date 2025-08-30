package nintendo

import (
	"fmt"
	"net/http"
	"nintendo-api/nxapi"
	"nintendo-api/utils"
)

func GetFriendList(webServiceAccessToken string) (*FriendListResponse, error) {

	config, err := nxapi.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v3/Friend/List", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer "+webServiceAccessToken)
	return utils.DoReq[FriendListResponse](request)
}
