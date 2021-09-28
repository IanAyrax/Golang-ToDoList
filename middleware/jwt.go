package middleware

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"example.com/GolangAPI2/model"
	"time"
	"os"
)

func CreateToken(logged_user model.User) (string, error){
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = logged_user.UserId
	atClaims["role_id"] = logged_user.RoleId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()	
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	return token, err
}

func ExtractToken(r *http.Request) string{
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if(len(strArr) == 2){
		return strArr[1]
	}

	return ""
}

func VerifyToken(r *http.Request) (string, string, error){
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	 })
	 
	 if err != nil {
		return "", "", err
	 }

	 role_id := ""
	 user_id := ""
	 //Verifying Role
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id = fmt.Sprintf("%v", claims["user_id"])
		fmt.Print("Helper RoleId : ")
		fmt.Println(claims["role_id"])
		if fmt.Sprintf("%v", claims["role_id"]) == "1"{
			//r.Header.Set("Role", "admin")
			fmt.Println("Admin")
			role_id = "1"
		}else if fmt.Sprintf("%v", claims["role_id"]) == "2"{
			//r.Header.Set("Role", "user")
			role_id = "2"
		}
	}
	 
	 return user_id, role_id, nil
}