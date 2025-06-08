package handlers

import (
	"encoding/json"
	"net/http"

	"car_db/internal/utils"
)

func GetBodyTypes(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	json.NewEncoder(w).Encode(db.BodyTypes)
}

func GetDriveTypes(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	json.NewEncoder(w).Encode(db.DriveTypes)
}

func GetEngines(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	json.NewEncoder(w).Encode(db.Engines)
}

func GetTransmissions(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	json.NewEncoder(w).Encode(db.Transmissions)
}
