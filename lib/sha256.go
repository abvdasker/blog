package lib

import (
	"crypto/sha256"
	"encoding/base64"
)

func Sha256EncodeBase64(input string) string {
	rawHash := sha256.Sum256([]byte(input))
	rawHashSlice := rawHash[:]
	return base64.StdEncoding.EncodeToString(rawHashSlice)
}
