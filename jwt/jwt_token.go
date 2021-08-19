package jwt

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(token string) (string, error) {
	parser := jwt.Parser{}
	tokenStr, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	jsondata, err := json.MarshalIndent(tokenStr.Claims, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsondata), nil
}

