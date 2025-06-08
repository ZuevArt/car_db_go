package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"car_db/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	found := false
	for i, user := range db.Users {
		if user.ID == id {
			db.Users = append(db.Users[:i], db.Users[i+1:]...)
			db.Save("data.json")
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	email := claims["email"].(string)
	db := utils.GetDB()

	for _, user := range db.Users {
		if user.Email == email {
			utils.WriteJSON(w, http.StatusOK, user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	var updatedData struct {
		FName    string `json:"f_name"`
		LName    string `json:"l_name"`
		Phone    string `json:"phone"`
		Password string `json:"password,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	for i, user := range db.Users {
		if user.Email == email {
			db.Users[i].FName = updatedData.FName
			db.Users[i].LName = updatedData.LName
			db.Users[i].Phone = updatedData.Phone

			// Если указан новый пароль — хэшируем его
			if updatedData.Password != "" {
				hashed, _ := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
				db.Users[i].PasswordHash = string(hashed)
			}

			db.Save("data.json")
			utils.WriteJSON(w, http.StatusOK, db.Users[i])
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
}
