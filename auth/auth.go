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
		"role": "admin",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

//func TokenParse(myToken string, myKey string) (*jwt.Token, error) {
func TokenParse(myToken string, myKey string) (*Authentication, error) {

	authinfo := Authentication{Name:"guest",Role:"guest",Token:""}
	if myToken == ""{
		fmt.Print("err: myToken is blan.")
		return &authinfo,nil
	}

	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err!=nil {
		fmt.Println(err)
	}else if token.Valid {
		//claims := token.Claims
		claims := token.Claims.(jwt.MapClaims)

		authinfo.Name,_ = claims["user"].(string)
		authinfo.Role,_ = claims["role"].(string)
		if authinfo.Name=="" || authinfo.Role==""{
			fmt.Println("TokenParse error: cann't get vale auth.Name or auth.Role.")
		}
		authinfo.Token = myToken

	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")

	}
	return &authinfo,err

}


