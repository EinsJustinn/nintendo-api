package auth

type (
	SessionTokenResponse struct {
		SessionToken string `json:"session_token"`
		Code         string `json:"code"`
	}

	tokenRequest struct {
		ClientId     string `json:"client_id"`
		SessionToken string `json:"session_token"`
		GrantType    string `json:"grant_type"`
	}

	TokenResponse struct {
		ExpiresIn   int      `json:"expires_in"`
		AccessToken string   `json:"access_token"`
		Scope       []string `json:"scope"`
		TokenType   string   `json:"token_type"`
		IdToken     string   `json:"id_token"`
	}

	NsoLoginOptions struct {
		State                           string `url:"state"`
		RedirectURI                     string `url:"redirect_uri"`
		ClientID                        string `url:"client_id"`
		Scope                           string `url:"scope"`
		ResponseType                    string `url:"response_type"`
		SessionTokenCodeChallenge       string `url:"session_token_code_challenge"`
		SessionTokenCodeChallengeMethod string `url:"session_token_code_challenge_method"`
		Theme                           string `url:"theme"`
	}

	LoginRequest struct {
		Parameter LoginRequestParameter `json:"parameter"`
	}

	LoginRequestParameter struct {
		NaIdToken  string `json:"naIdToken"`
		NaBirthday string `json:"naBirthday"`
		NaCountry  string `json:"naCountry"`
		Language   string `json:"language"`
		Timestamp  int64  `json:"timestamp"`
		RequestId  string `json:"requestId"`
		F          string `json:"f"`
	}

	LoginResponse struct {
		Status int `json:"status"`
		Result struct {
			User struct {
				Id                int64  `json:"id"`
				NsaId             string `json:"nsaId"`
				ImageUri          string `json:"imageUri"`
				Image2Uri         string `json:"image2Uri"`
				Name              string `json:"name"`
				SupportId         string `json:"supportId"`
				IsChildRestricted bool   `json:"isChildRestricted"`
				Etag              string `json:"etag"`
				Links             struct {
					NintendoAccount struct {
						Membership struct {
							Active bool `json:"active"`
						} `json:"membership"`
					} `json:"nintendoAccount"`
					FriendCode struct {
						Regenerable   bool   `json:"regenerable"`
						RegenerableAt int    `json:"regenerableAt"`
						Id            string `json:"id"`
					} `json:"friendCode"`
				} `json:"links"`
				Permissions struct {
					PlayLog                string `json:"playLog"`
					Presence               string `json:"presence"`
					FriendRequestReception bool   `json:"friendRequestReception"`
				} `json:"permissions"`
				Presence struct {
					State     string `json:"state"`
					UpdatedAt int    `json:"updatedAt"`
					LogoutAt  int    `json:"logoutAt"`
					Game      struct {
					} `json:"game"`
				} `json:"presence"`
			} `json:"user"`
			WebApiServerCredential struct {
				AccessToken string `json:"accessToken"`
				ExpiresIn   int    `json:"expiresIn"`
			} `json:"webApiServerCredential"`
			FirebaseCredential struct {
				AccessToken string `json:"accessToken"`
				ExpiresIn   int    `json:"expiresIn"`
			} `json:"firebaseCredential"`
		} `json:"result"`
		CorrelationId string `json:"correlationId"`
	}
)
