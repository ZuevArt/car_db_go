package handlers

import (
	"net/http"
	"strconv"

	"car_db/internal/utils"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	modelIDStr := r.URL.Query().Get("model_id")

	if modelIDStr != "" {
		modelID, err := strconv.Atoi(modelIDStr)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid model ID"})
			return
		}

		for _, model := range db.Models {
			if model.ModelID == modelID {
				utils.WriteJSON(w, http.StatusOK, model)
				return
			}
		}

		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "Model not found"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, db.Models)
}
