package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var oauth2Config = &oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Endpoint:     microsoft.AzureADEndpoint(os.Getenv("TENANT_ID")),
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	Scopes:       []string{"openid", "profile", "email"},
}

// HandleMicrosoftLogin redirects to Microsoft OAuth2 login page
func HandleMicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleMicrosoftCallback
func HandleMicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")

	// Exchange the code for a token
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Validate the token and extract user info
	idToken := token.Extra("id_token").(string)
	// store jwt Token
	jwtToken, err := ValidateTokenAndGenerateJWT(idToken)
	if err != nil {
		http.Error(w, "Failed to validate token", http.StatusInternalServerError)
		return
	}

	// Return the JWT token to the user
	fmt.Fprintf(w, "JWT Token: %s", jwtToken)

}
