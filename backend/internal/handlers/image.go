package handlers

import (
	"net/http"
	"strconv"

	"car_db/internal/utils"
)

// Получает все изображения по car_id
func GetImagesByCarID(carID int, db *utils.DB) []string {
	var urls []string
	for _, img := range db.CarImages {
		if img.CarID == carID && img.IsPrimary {
			urls = append(urls, img.URL)
		}
	}
	return urls
}

// Обработчик: GET /api/cars/{id}/images
func GetCarImages(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	db := utils.GetDB()
	var images []string
	for _, img := range db.CarImages {
		if img.CarID == id && img.IsPrimary {
			images = append(images, img.URL)
		}
	}

	utils.WriteJSON(w, http.StatusOK, images)
}
