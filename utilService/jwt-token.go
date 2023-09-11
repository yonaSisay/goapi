package utilService

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createJWTToken(payload map[string]interface{}, secretKey string, tokenExpiration int64) (string, error) {
    // Create a new JWT token with the given payload and secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
    // Set the token expiration time to 1 hour from now
    token.Claims.(jwt.MapClaims)["exp"] = tokenExpiration
    // Sign the token with the secret key and return the signed token as a string
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func GetToken(userId string, role string ) (string, error){
	payload := map[string]interface{}{
		"sub": "12345",         // The user ID
		"iat": time.Now().Unix(),       // The token issue time (UNIX timestamp)
		"exp": time.Now().Add(time.Hour * 48).Unix(),  // The token expiration time (UNIX timestamp)
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user", "admin"},  // The allowed roles for the user
			"x-hasura-default-role":  "user",            // The default role for the user
			"x-hasura-user-id":       userId,           // The user ID
			"x-hasura-role":          role,              // The current role for the user
		},	
	}
	tokenExpiration := time.Now().Add(time.Hour * 48).Unix()
	token, err := createJWTToken(payload, os.Getenv("JWT_SECRET"),tokenExpiration )
	if err != nil {
		return "", err
	}
	return token, nil
}
