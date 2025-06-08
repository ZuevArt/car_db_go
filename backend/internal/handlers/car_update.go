package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car_db/internal/models"
	"car_db/internal/utils"
)

type ExtendedCar struct {
	models.Car
	Brand     string `json:"brand"`
	ModelName string `json:"model_name"`
	ImageURL  string `json:"image_url"`
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	carIDStr := r.PathValue("id")
	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar models.Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	found := false

	for i, car := range db.Cars {
		if car.ID == carID {
			db.Cars[i] = updatedCar
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	db.Save("data.json")
	utils.WriteJSON(w, http.StatusOK, db.Cars[carID])
}
