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

func TokenParse(myToken string, myKey string) (*jwt.Token, string) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	if err == nil && token.Valid {
		//claims := token.Claims
		if token.Valid {
			return token, "ok"
		}
		return nil, "can not get claims"
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
		return nil, "InvaildToken"
	}

}
