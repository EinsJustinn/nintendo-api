package nintendo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetShowSelf(webServiceAccessToken string) (ShowSelfResponse, error) {
	request, err := http.NewRequest("POST", nintendoBaseUrl+"/v4/User/ShowSelf", nil)
	if err != nil {
		return ShowSelfResponse{}, err
	}
	request.Header.Set("Authorization", "Bearer "+webServiceAccessToken)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ShowSelfResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return ShowSelfResponse{}, fmt.Errorf("status code: %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ShowSelfResponse{}, err
	}
	var userResponse ShowSelfResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return ShowSelfResponse{}, err
	}
	return userResponse, nil
}
