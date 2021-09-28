package helper

import (
	"errors"
)
	

func IsAdmin(roleId string) error {
	if roleId == "1" {
		return nil
	}else{
		return errors.New("Not Admin, Access Denied!!!")
	}
}