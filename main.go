package main

import (
	"fmt"
	"nintendo-api/nintendo/auth"
)

func main() {
	url, s, err := auth.GenerateLoginUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	fmt.Println("Verifier: " + s)
}
