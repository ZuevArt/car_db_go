package handlers

import (
	"car_db/internal/models"
	"car_db/internal/utils"
	"encoding/json"

	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Role = ""

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hashedPassword)

	adminKey := r.Header.Get("X-Admin-Key")
	if adminKey == "" {
		adminKey = r.URL.Query().Get("admin_key")
	}

	if adminKey != "" && adminKey == os.Getenv("ADMIN_KEY") {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	db := utils.GetDB()
	user.ID = len(db.Users) + 1
	db.Users = append(db.Users, user)
	db.Save("data.json")

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password_hash"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser := GetUserFromDB(creds.Email)

	if storedUser.ID == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"email": creds.Email,
		"role":  storedUser.Role,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func GetUserFromDB(email string) models.User {
	db := utils.GetDB()
	for _, user := range db.Users {
		if strings.EqualFold(user.Email, email) {
			return user
		}
	}
	return models.User{}
}
