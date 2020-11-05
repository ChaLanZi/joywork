package helper

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/ttacon/libphonenumber"
)

const (
	defaultRegion = "CN"
)

func ParseAndFormatPhoneNumber(in string) (cleanPhoneNumber string, err error) {
	if in == "" {
		return
	}

	p, err := libphonenumber.Parse(in, defaultRegion)
	if err != nil {
		return "", fmt.Errorf(" Invalid phone number")
	}
	cleanPhoneNumber = libphonenumber.Format(p, libphonenumber.E164)
	return
}

func GenerateGravatarURL(email string) string {
	h := md5.New()
	_, _ = io.WriteString(h, email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x.jpg?s=400&d=identicon", h.Sum(nil))
}

func Int64ToBool(in int64) bool {
	if in == 0 {
		return false
	}
	return true
}

func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
