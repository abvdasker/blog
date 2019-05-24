package lib

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	username := "username"
	salt := "salt"
	expiresAt := time.Time{}
	expectedToken := "VBjkcLm1u6ug/RBFLBguIK+7HXgmIgjlyHdvY4g7t4Q="

	token := GenerateToken(username, salt, expiresAt)

	if token != expectedToken {
		t.Fatalf("username %s, salt %s and expiration %s was expected to hash to %s but hashed to %s",
			username, salt, expiresAt.String(), expectedToken, token)
	}
}
