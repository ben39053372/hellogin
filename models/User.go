package models

import (
	"hellogin/config"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllUsers fetch all user
func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUsers create a new user
func CreateUsers(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
