package utilService

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword string, password string) bool {
    // Convert the hashed password string to a byte slice
    hashedPasswordBytes := []byte(hashedPassword)
    // Compare the hashed password byte slice with the plaintext password byte slice
    err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(password))
    if err != nil {
        // If the error is not nil, the passwords do not match
        return false
    }
    // If the error is nil, the passwords match
    return true
}

func HashPassword(password string) (string, error) {
    // Generate a salt for the password hash
    // The cost parameter is used to determine the complexity of the hash
    // Increasing the cost parameter makes the hash more computationally expensive to generate
    // and therefore harder to crack
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}