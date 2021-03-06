package lib

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
)

const (
	passwordDigestTemplate = "%s|%s|%s"
	saltRandomBytes        = 16
)

func RandomSalt64() string {
	buffer := make([]byte, saltRandomBytes)
	n, err := rand.Read(buffer)
	if n != len(buffer) || err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buffer)
}

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
