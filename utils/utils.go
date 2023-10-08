package utils

import (
	"errors"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(password, hash string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func IsValidEmail(email string) (bool, error) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	isValid := regex.MatchString(email)

	if isValid {
		return true, nil
	}
	return false, errors.New("invalid mail")
}

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	secretKey := []byte("your-secret-key")

	claims["iat"]=time.Now().Unix()
	claims["exp"]=time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)

	return signedToken, err
}

func DecodeJWT(tokenString string) (map[string]interface{}, error) {
	secretKey := []byte("your-secret-key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("not getting claims")
	}

	return claims, nil
}
