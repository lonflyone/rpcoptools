package utils

import "testing"

func TestGetToken(t *testing.T) {
	var used uint = 123
	token, _ := GenerateToken(used)
	t.Log(token)
	calm, _ := ParseToken((token))
	t.Log(calm.ExpiresAt)
	t.Log(jwtSecret)
}
