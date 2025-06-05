package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"car_db/internal/models"
	"car_db/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("7f1b7e0d3c6a4b2e8f5d9c3b5a8e6f4d7c1b3e5a9f2d4c6b8e1f3a5d7c9b2e4")

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Could not hash password", http.StatusInternalServerError)
        return
    }
    user.PasswordHash = string(hashedPassword)

    if user.Role == "" {
        user.Role = "user"
    }

    db := utils.GetDB()
    user.ID = len(db.Users) + 1 

    db.Users = append(db.Users, user)
    db.Save("data.json")

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser := GetUserFromDB(creds.Email)
	
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(creds.PasswordHash)); err != nil {
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
        if user.Email == email {
            return user
        }
    }
    return models.User{}
}