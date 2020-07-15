package auth

import (
	"testing"
)

func TestGenerateAccessToken(t *testing.T) {
	token1 := GenerateAccessToken("1", true)
	t.Fatalf("token %#v", token1)
}
