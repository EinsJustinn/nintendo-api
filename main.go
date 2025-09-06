package main

import (
	"fmt"
	"github.com/einsjustinn/nintendo-api/nintendo/auth"
)

func main() {
	url, s, err := auth.GenerateLoginUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	fmt.Println("Verifier: " + s)
}
