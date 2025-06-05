package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"car_db/internal/models"
	"car_db/internal/utils"
)

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

func GetCars(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	cars := db.Cars

	//Filtering: by brand
	brandFilter := strings.ToLower(r.URL.Query().Get("brand"))
	if brandFilter != "" {
		validModelIDs := utils.GetModelIDsByBrand(brandFilter)
		if len(validModelIDs) == 0 {
			utils.WriteJSON(w, http.StatusOK, []models.Car{})
			return
		}

		var filteredCars []models.Car
		for _, car := range cars {
			for _, modelID := range validModelIDs {
				if car.Model == modelID {
					filteredCars = append(filteredCars, car)
					break
				}
			}
		}
		cars = filteredCars
	}

	//Filter by model
	modelFilter := r.URL.Query().Get("model")
	if modelFilter != "" {
		var filteredCars []models.Car
		modelID, err := strconv.Atoi(modelFilter)
		if err == nil {
			for _, car := range cars {
				if car.Model == modelID {
					filteredCars = append(filteredCars, car)
				}
			}
		}
		cars = filteredCars
	}

	//Filter by year
	yearFilter := r.URL.Query().Get("year")
	if yearFilter != "" {
		year, err := strconv.Atoi(yearFilter)
		if err == nil {
			var filteredCars []models.Car
			for _, car := range cars {
				if car.Year == year {
					filteredCars = append(filteredCars, car)
				}
			}
			cars = filteredCars
		}
	}

	//Filter by engine
	engineFilter := strings.ToLower(r.URL.Query().Get("engine"))
	if engineFilter != "" {
		var filteredCars []models.Car
		for _, car := range cars {
			if strings.ToLower(car.Engine) == engineFilter {
				filteredCars = append(filteredCars, car)
			}
		}
		cars = filteredCars
	}

	//Filter by body type
	bodyTypeFilter := strings.ToLower(r.URL.Query().Get("body_type"))
	if bodyTypeFilter != "" {
		var filteredCars []models.Car
		for _, car := range cars {
			if strings.ToLower(car.BodyType) == bodyTypeFilter {
				filteredCars = append(filteredCars, car)
			}
		}
		cars = filteredCars
	}

	//Filter by price
	minPrice := r.URL.Query().Get("min_price")
	maxPrice := r.URL.Query().Get("max_price")

	if minPrice != "" || maxPrice != "" {
		var filteredCars []models.Car
		var min float64 = -1
		var max float64 = -1

		if minPrice != "" {
			if val, err := strconv.ParseFloat(minPrice, 64); err == nil {
				min = val
			}
		}
		if maxPrice != "" {
			if val, err := strconv.ParseFloat(maxPrice, 64); err == nil {
				max = val
			}
		}

		for _, car := range cars {
			if (min == -1 || car.Price >= min) && (max == -1 || car.Price <= max) {
				filteredCars = append(filteredCars, car)
			}
		}
		cars = filteredCars
	}

	sortBy := r.URL.Query().Get("sort")

	switch sortBy {
	case "price_asc":
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].Price < cars[j].Price
		})
	case "price_desc":
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].Price > cars[j].Price
		})
	case "year_asc":
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].Year < cars[j].Year
		})
	case "year_desc":
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].Year > cars[j].Year
		})
	case "brand_asc":
		sort.Slice(cars, func(i, j int) bool {
			modelI := utils.GetModelByID(cars[i].Model)
			modelJ := utils.GetModelByID(cars[j].Model)
			return strings.ToLower(modelI.Brand) < strings.ToLower(modelJ.Brand)
		})
	case "brand_desc":
		sort.Slice(cars, func(i, j int) bool {
			modelI := utils.GetModelByID(cars[i].Model)
			modelJ := utils.GetModelByID(cars[j].Model)
			return strings.ToLower(modelI.Brand) > strings.ToLower(modelJ.Brand)
		})
	case "engine_asc":
		sort.Slice(cars, func(i, j int) bool {
			return strings.ToLower(cars[i].Engine) < strings.ToLower(cars[j].Engine)
		})
	case "engine_desc":
		sort.Slice(cars, func(i, j int) bool {
			return strings.ToLower(cars[i].Engine) > strings.ToLower(cars[j].Engine)
		})
	case "body_type_asc":
		sort.Slice(cars, func(i, j int) bool {
			return strings.ToLower(cars[i].BodyType) < strings.ToLower(cars[j].BodyType)
		})
	case "body_type_desc":
		sort.Slice(cars, func(i, j int) bool {
			return strings.ToLower(cars[i].BodyType) > strings.ToLower(cars[j].BodyType)
		})
	default:
		// Default sort: newest first
		sort.Slice(cars, func(i, j int) bool {
			return cars[i].Year > cars[j].Year
		})
	}

	utils.WriteJSON(w, http.StatusOK, cars)
}
