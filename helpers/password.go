package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHass), err
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
