package utils

import (
	"encoding/base64"
	"fmt"
)

func GenerateSignature(singSource string) string {
	return base64.StdEncoding.EncodeToString([]byte(singSource))
}

func CheckSignature(signature, singSource string) error {
	singB64 := base64.StdEncoding.EncodeToString([]byte(singSource))
	if signature != singB64 {
		return fmt.Errorf("\r\nrequest :%s\r\ngenerate:%s", signature, singB64)
	}
	return nil
}
