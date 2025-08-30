package nintendo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	nintendoAccountBaseUrl = "https://api.accounts.nintendo.com/2.0.0"
	nintendoBaseUrl        = "https://api-lp1.znc.srv.nintendo.net"
)

func GetUserInfo(token string) (UserResponse, error) {
	request, err := http.NewRequest("GET", nintendoAccountBaseUrl+"/users/me", nil)
	if err != nil {
		return UserResponse{}, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return UserResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return UserResponse{}, fmt.Errorf("status code: %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return UserResponse{}, err
	}
	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return UserResponse{}, err
	}
	return userResponse, nil
}
