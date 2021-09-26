package helper

import (
	"net/http"
	"errors"
)
	

func IsAdmin(r *http.Request) error {
	if r.Header.Get("RoleId") == "1" {
		return nil
	}else{
		return errors.New("Not Admin, Access Denied!!!")
	}
}