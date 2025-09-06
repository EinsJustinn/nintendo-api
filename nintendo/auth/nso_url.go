package auth

import (
	"fmt"
	"github.com/einsjustinn/nintendo-api/utils"

	"github.com/google/go-querystring/query"
)

func GenerateLoginUrl() (string, string, error) {
	verifier, challenge, err := utils.GeneratePKCE()
	if err != nil {
		return "", "", err
	}

	params := NsoLoginOptions{
		State:                           utils.GenerateRandomString(50),
		RedirectURI:                     "npf71b963c1b7b6d119://auth",
		ClientID:                        "71b963c1b7b6d119",
		Scope:                           "openid user user.birthday user.screenName",
		ResponseType:                    "session_token_code",
		SessionTokenCodeChallenge:       challenge,
		SessionTokenCodeChallengeMethod: "S256",
		Theme:                           "login_form",
	}

	queryString, _ := query.Values(params)

	baseURL := "https://accounts.nintendo.com/connect/1.0.0/authorize"
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryString.Encode())
	return fullURL, verifier, nil
}
