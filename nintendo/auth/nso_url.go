package auth

import (
	"fmt"

	"github.com/einsjustinn/nintendo-api/utils"

	"github.com/google/go-querystring/query"
)

/*
Client IDs
Nintendo Switch Online App: 71b963c1b7b6d119
Nintendo Music App: a9b03ca3519e14f4
*/

type Client struct {
	ID string
}

var (
	NintendoSwitchOnline = Client{ID: "71b963c1b7b6d119"}
	NintendoMusic        = Client{ID: "a9b03ca3519e14f4"}
)

func GenerateLoginUrl(client Client) (string, string, error) {
	verifier, challenge, err := utils.GeneratePKCE()
	if err != nil {
		return "", "", err
	}

	params := NsoLoginOptions{
		State:                           utils.GenerateRandomString(50),
		RedirectURI:                     fmt.Sprintf("npf%s://auth", client.ID),
		ClientID:                        client.ID, // for nintendo switch online app
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
