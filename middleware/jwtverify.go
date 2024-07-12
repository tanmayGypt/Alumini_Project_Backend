package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Extract the JWT token from the Authorization header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		jwtToken := tokenParts[1]

		// Parse and validate the JWT token
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			jwtKey := []byte(os.Getenv("JWT_KEY")) // Load your JWT secret key from environment variables
			return jwtKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !claims.VerifyIssuer("https://login.microsoftonline.com/", false) {
			http.Error(w, "Invalid token issuer", http.StatusUnauthorized)
			return
		}

		if claims["aud"] != os.Getenv("CLIENT_ID") {
			http.Error(w, "Invalid token audience", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
