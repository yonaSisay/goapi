package utilService

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// generate hashed password
	// the cost is to determine the complexity of the hashing algorithm
	// the higher the cost, the longer it takes to generate the hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}