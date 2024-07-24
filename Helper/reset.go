package helper

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
)

func GenerateToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}

func GenerateLink(token string) string {
	baseURL := "https://bpit-alumni-portal.com/change-password"
	return fmt.Sprintf("%s?token=%s", baseURL, url.QueryEscape(token))
}
