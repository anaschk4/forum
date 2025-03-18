package main

import (
	"fmt"
	"net/http"
	"time"
)

type login struct{
	hashedPassword string
	SessionToken string
	CSRFToken string
}

var users = map[string]login{}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	