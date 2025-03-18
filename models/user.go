package models


type User struct {
	ID       int    
	Username string 
	Password string 
	
}




func (u *User) Authenticate(password string) bool {
	
	
	return u.Password == password 
}