package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pswd := []byte(password)
	cost := 10
	hash, err := bcrypt.GenerateFromPassword(pswd, cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	pswd := []byte(password)
	byte_hash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(pswd, byte_hash)
	if err != nil {
		return err
	} else {
		return nil
	}
}
