package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestJwt(t *testing.T) {
	var hmacSampleSecret []byte

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"luenci": "123456",
		"time":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	fmt.Println(token)

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret = []byte("test")
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	fmt.Println(token, err)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}
