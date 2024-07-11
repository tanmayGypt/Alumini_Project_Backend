package controllers

import (
	"fmt"
	utils "my-go-backend/Utils"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// Validate token and generate a new JWT
func ValidateTokenAndGenerateJWT(idToken string) (string, error) {
	// Parsing the token
	token, _ := jwt.Parse(idToken, nil)
	claims := token.Claims.(jwt.MapClaims)
	kid := token.Header["kid"].(string)
	issuer := claims["iss"].(string)
	audience := claims["aud"].(string)

	// Fetching the JWKS (JSON Web Key Set)
	jwks, err := utils.FetchJWKS("https://login.microsoftonline.com/common/discovery/v2.0/keys")
	if err != nil {
		return "", err
	}

	// Find the key
	key, err := jwks.FindKey(kid)
	if err != nil {
		return "", err
	}

	// Validate the token
	parsedToken, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil || !parsedToken.Valid {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	// Validate issuer and audience
	if !strings.HasPrefix(issuer, "https://login.microsoftonline.com/") || audience != os.Getenv("CLIENT_ID") {
		return "", fmt.Errorf("invalid issuer or audience")
	}

	// Generating a new JWT
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": claims["sub"],
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	jwtString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
