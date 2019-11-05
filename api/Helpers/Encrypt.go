package Helpers

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
