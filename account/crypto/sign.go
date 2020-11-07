package crypto

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func EmailConfirmationToken(uuid, email, signingToken string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"email": email,
		"uuid":  uuid,
		"exp":   time.Now().Add(time.Duration(2 * time.Hour)).Unix(), // 超时时间2小时
	}).SignedString([]byte(signingToken))

	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyEmailConfirmationToken(tokenString, signingToken string) (uuid, email string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(" Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingToken), nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid = claims["uuid"].(string)
		email = claims["email"].(string)
		return
	}
	return "", "", fmt.Errorf("unable to verify token")
}

func SessionToken(uuid, signingToken string, support bool, dur time.Duration) (string, error) {
	if len(signingToken) == 0 {
		panic("No Signing token present")
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"uuid":    uuid,
		"support": support,
		"exp":     time.Now().Add(dur).Unix(),
	}).SignedString([]byte(signingToken))
	if err != nil {
		return "", err
	}
	return token, nil
}

func RetrieveSessionInformation(tokenString, signingToken string) (uuid string, support bool, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(" Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingToken), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid = claims["uuid"].(string)
		support = claims["support"].(bool)
		return
	}
	err = fmt.Errorf("unable to verify token")
	return
}
