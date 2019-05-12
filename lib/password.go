package lib

import (
	"crypto/subtle"
	"fmt"
)

const (
	passwordDigestTemplate = "%s|%s|%s"
)

func HashPassword64(username, salt, password string) string {
	digest := createDigest(username, salt, password)
	return Sha256EncodeBase64(digest)
}

func SecureStringsEqual(first, second string) bool {
	result := subtle.ConstantTimeCompare(
		[]byte(first),
		[]byte(second),
	)
	return result == 1
}

func createDigest(username, salt, password string) string {
	return fmt.Sprintf(
		passwordDigestTemplate,
		username, salt, password,
	)
}
