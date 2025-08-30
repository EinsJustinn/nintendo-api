package nintendo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nintendo-api/nxapi"
)

func GetFriendList(webServiceAccessToken string) (FriendListResponse, error) {

	config, err := nxapi.GetConfig()
	if err != nil {
		return FriendListResponse{}, fmt.Errorf("failed to get config: %w", err)
	}
	configVersion := config.Versions[0]

	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v3/Friend/List", nil)
	if err != nil {
		return FriendListResponse{}, err
	}
	request.Header.Set("X-Platform", configVersion.Platform)
	request.Header.Set("X-ProductVersion", configVersion.Version)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Authorization", "Bearer "+webServiceAccessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return FriendListResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return FriendListResponse{}, fmt.Errorf("status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FriendListResponse{}, err
	}
	var friendListResponse FriendListResponse
	err = json.Unmarshal(body, &friendListResponse)
	if err != nil {
		return FriendListResponse{}, err
	}
	return friendListResponse, nil
}
