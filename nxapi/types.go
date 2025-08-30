package nxapi

type (
	TokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		Scope        string `json:"scope"`
		RefreshToken string `json:"refresh_token"`
	}

	fGenerationRequest struct {
		HashMethod  string `json:"hash_method"`
		Token       string `json:"token"`
		NaID        string `json:"na_id"`
		CoralUserId string `json:"coral_user_id,omitempty"`
	}

	FGenerationResponse struct {
		F         string `json:"f"`
		Timestamp int64  `json:"timestamp"`
		RequestId string `json:"request_id"`
	}

	ConfigResponse struct {
		Versions []struct {
			Platform    string `json:"platform"`
			Version     string `json:"version"`
			Build       int    `json:"build"`
			WorkerCount int    `json:"worker_count"`
		} `json:"versions"`
		NsoVersion string `json:"nso_version"`
	}
)
