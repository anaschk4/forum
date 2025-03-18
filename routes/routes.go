package routes

import (
	"forum/handlers"
	"net/http"
)


func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	return mux
}
