package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    
	Username string 
	Password string 
}

func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}