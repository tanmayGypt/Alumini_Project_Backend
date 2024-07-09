package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kty   string   `json:"kty"`
	Use   string   `json:"use"`
	KeyId string   `json:"kid"`
	X5t   string   `json:"x5t"`
	N     string   `json:"n"`
	E     string   `json:"e"`
	X5c   []string `json:"x5c"`
}

func FetchJWKS(url string) (*JWKS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks JWKS
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	return &jwks, nil
}

func (jwks *JWKS) FindKey(kid string) (*rsa.PublicKey, error) {
	for _, jwk := range jwks.Keys {
		if jwk.KeyId == kid {
			return jwk.rsaPublicKey()
		}
	}
	return nil, fmt.Errorf("unable to find appropriate key")
}

func (jwk *JWK) rsaPublicKey() (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	e := 0
	for _, b := range eBytes {
		e = e*256 + int(b)
	}

	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: e,
	}

	return pubKey, nil
}
