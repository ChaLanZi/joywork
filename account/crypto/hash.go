package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(salt, password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(append(salt, password...), 12)
}

func CheckPasswordHash(hash, salt, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, append(salt, password...))
}
