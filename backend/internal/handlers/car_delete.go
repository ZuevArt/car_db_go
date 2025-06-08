package handlers

import (
	"net/http"
	"strconv"

	"car_db/internal/utils"
)

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Missing car ID", http.StatusBadRequest)
		return
	}

	carID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	found := false

	for i, car := range db.Cars {
		if car.ID == carID {
			db.Cars = append(db.Cars[:i], db.Cars[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	db.Save("data.json")

	w.WriteHeader(http.StatusOK)
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Автомобиль удален"})
}
