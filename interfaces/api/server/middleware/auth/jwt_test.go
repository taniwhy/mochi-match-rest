package auth

import (
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	token1, _ := GenerateRefreshToken("1")
	t.Fatalf("token %#v", token1)
}
