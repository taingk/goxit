package Helpers

import (
	"crypto/sha256"
	"encoding/base64"
)

func encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}
