package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "car_db/internal/utils"
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