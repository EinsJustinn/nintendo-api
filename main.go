package main

import (
	"fmt"
	"nintendo-api/nintendo/nintendo_auth"
)

func main() {
	url, s, err := nintendo_auth.GenerateLoginUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	fmt.Println("Verifier: " + s)
}
