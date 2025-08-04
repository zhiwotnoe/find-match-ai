package handlers_test

import (
	"testing"
	"tinder_ai_core/auth"
)

func TestJWT(t *testing.T) {
	token, err := auth.GenerateToken("user123")
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}
}
