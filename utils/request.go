package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DoReq[Resp any](req *http.Request) (*Resp, error) {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	resp := new(Resp)
	err = json.NewDecoder(response.Body).Decode(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return resp, nil
}
