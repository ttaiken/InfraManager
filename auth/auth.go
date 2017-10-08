package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func TokenNew(mySigningKey []byte, username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

func TokenParse(myToken string, myKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err!=nil {
		fmt.Println(err)
	}
	if err == nil && token.Valid {
		//claims := token.Claims

		return token, nil

	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
		return nil,err
	}

}
