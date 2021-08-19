package jwt

import (
	"testing"
)

func TestTokenDecode(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	token, err := ParseToken(token)
	if err != nil {
		t.Errorf("unable to parse token: %v", err)
		return
	}

	t.Logf("parse token: %s", token)
}
