package handlers

import (
    "net/http"
	
    "car_db/internal/models"
    "car_db/internal/utils"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
    db := utils.GetDB()
    utils.WriteJSON(w, http.StatusOK, db.Cars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
    var car models.Car
    if err := utils.ReadJSON(r, &car); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    db := utils.GetDB()
    car.ID = len(db.Cars) + 1
    db.Cars = append(db.Cars, car)
    db.Save("data.json")

    utils.WriteJSON(w, http.StatusCreated, car)
}