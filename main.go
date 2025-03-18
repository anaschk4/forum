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
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {}
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.error(w, "Invalid method", er)
		return
	}

	username := rFormValue("username")
	password := rFormValue("password")
	if len(username)<8 || len(password)<8 {
		er := http.StatusNotAcceptable
		http.error(w, "Invalid username or password", er)
		return
	}

	if _, ok := users[username]; ok {
		er := http.StatusConflict
		http.error(w, "Username already exists", er)
		return
	}

	hashPassword, _ := hashPassword(password)
	users[username] = login{
		hashedPassword: hashPassword,
	}

	fmt.Fprintf(w, "User registered successfully ! ")

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func protected(w http.ResponseWriter, r *http.Request) {}