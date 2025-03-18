package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"forum/config"
	"forum/models"
	"net/http"
	"text/template"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user models.User
		err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			return
		}

		// Use bcrypt instead of custom sha256
		if !user.Authenticate(password) {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		// Create a session with proper expiry
		sessionUUID := generateSessionID()
		expiresAt := time.Now().Add(24 * time.Hour)
		
		// Store session in database (simplified)
		_, err = config.DB.Exec("INSERT INTO sessions (uuid, user_id, username, created_at, expires_at) VALUES (?, ?, ?, ?, ?)",
			sessionUUID, user.ID, user.Username, time.Now(), expiresAt)
		
		if err != nil {
			http.Error(w, "Erreur de création de session", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionUUID,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiresAt,
		})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

// Helper function to generate session ID
func generateSessionID() string {
	// In a real application, you would use a more secure random generator
	hash := sha256.Sum256([]byte(time.Now().String()))
	return hex.EncodeToString(hash[:])
}