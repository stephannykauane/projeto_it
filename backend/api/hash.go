package api

import (
	"encoding/hex"
	"crypto/sha256"
)

func hash (senha string) string{
	h := sha256.New()
	h.Write([]byte(senha))

	return hex.EncodeToString(h.Sum(nil))
}