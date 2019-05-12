package lib

import (
	"fmt"
	"time"
)

func GenerateToken(username string, salt string, expiresAt time.Time) string {
	expiresAtEpoch := expiresAt.Unix()
	digest := fmt.Sprintf("%s|%s|%d", username, salt, expiresAtEpoch)
	return Sha256EncodeBase64(digest)
}
