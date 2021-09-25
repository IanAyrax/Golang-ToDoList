package helper

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func ExtractToken(r *http.Request) string{
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if(len(strArr) == 2){
		return strArr[1]
	}

	return ""
}

func VerifyToken(r *http.Request) error{
	tokenString := ExtractToken(r)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	 })
	 
	 fmt.Println(err)
	 //fmt.Println(token)
	 if err != nil {
		return err
	 }
	 
	 return nil
}