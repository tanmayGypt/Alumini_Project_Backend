package controllers

import (
	"context"
	"fmt"
	"log"
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
	clientID := os.Getenv("CLIENT_ID")
	redirectURL := os.Getenv("REDIRECT_URL")
	tenantID := os.Getenv("TENANT_ID")

	log.Printf("CLIENT_ID: %s, REDIRECT_URL: %s, TENANT_ID: %s\n", clientID, redirectURL, tenantID)

	if clientID == "" || redirectURL == "" || tenantID == "" {
		log.Fatal("Missing required environment variables")
	}

	authURL := fmt.Sprintf(
		"https://login.microsoftonline.com/%s/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=openid profile email",
		tenantID, clientID, redirectURL,
	)

	http.Redirect(w, r, authURL, http.StatusFound)
}

// HandleMicrosoftCallback handles the OAuth2 callback from Microsoft
func HandleMicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")

	// Exchange the code for a token
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Validating the token and extracting user info
	idToken := token.Extra("id_token").(string)
	// storing jwt Token
	jwtToken, err := ValidateTokenAndGenerateJWT(idToken)
	if err != nil {
		http.Error(w, "Failed to validate token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "JWT Token: %s", jwtToken)

}
