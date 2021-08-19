package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/prembhaskal/parse-jwt-token/jwt"
)

func main() {
	var input bytes.Buffer
	input.ReadFrom(os.Stdin)

	parsedToken, err := jwt.ParseToken(input.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse token: %v\n", err)
		return
	}

	fmt.Println(parsedToken)
}