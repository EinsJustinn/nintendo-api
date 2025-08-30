package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		fmt.Println("Failed to generate random string.")
	}
	return base64.URLEncoding.EncodeToString(b)[:n]
}

func GeneratePKCE() (codeVerifier string, codeChallenge string, err error) {
	verifier := make([]byte, 32)
	_, err = rand.Read(verifier)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate random code verifier: %w", err)
	}

	codeVerifier = base64.RawURLEncoding.EncodeToString(verifier)
	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge = base64.RawURLEncoding.EncodeToString(hash[:])

	return codeVerifier, codeChallenge, nil
}

func ExtractSessionTokenCode(userUrl string) (string, error) {
	parse, err := url.Parse(userUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}
	query, err := url.ParseQuery(parse.Fragment)
	if err != nil {
		return "", fmt.Errorf("failed to parse query: %w", err)
	}
	return query.Get("session_token_code"), nil
}
