package crypto

import "crypto/rand"

func NewSalt() ([]byte, error) {
	b := make([]byte, 60)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
