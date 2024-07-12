package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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
	// ctx := context.Background()
	code := r.URL.Query().Get("code")

	// Exchange the code for a token
	token, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validating the token and extracting user info
	idToken, ok := token["id_token"].(string)
	if !ok {
		http.Error(w, "No id_token in response", http.StatusInternalServerError)
		return
	}
	// storing jwt Token
	jwtToken, err := ValidateTokenAndGenerateJWT(idToken)
	if err != nil {
		http.Error(w, "Failed to validate token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "JWT Token: %s", jwtToken)
}

func exchangeCodeForToken(authCode string) (map[string]interface{}, error) {
	// Define the token endpoint
	tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", os.Getenv("TENANT_ID"))

	// Prepare the data for the POST request
	data := url.Values{}
	data.Set("client_id", os.Getenv("CLIENT_ID"))
	data.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", os.Getenv("REDIRECT_URL"))

	// Make the POST request
	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", body)
	}

	// Parse the response JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return result, nil
}
