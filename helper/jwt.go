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

func VerifyToken(r *http.Request) (string, error){
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	 })
	 
	 if err != nil {
		return "", err
	 }

	 role_id := ""
	 //Verifying Role
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role_id"] == "1"{
			//r.Header.Set("Role", "user")
			role_id = "1"
		}else if claims["role_id"] == "2"{
			//r.Header.Set("Role", "user")
			role_id = "2"
		}
	}
	 
	 return role_id, nil
}