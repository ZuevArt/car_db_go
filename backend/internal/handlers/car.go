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
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
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

	// 1. Фильтр по id
	if carIDStr := r.URL.Query().Get("id"); carIDStr != "" {
		carID, err := strconv.Atoi(carIDStr)
		if err == nil {
			var filteredCars []models.Car
			for _, car := range cars {
				if car.ID == carID {
					filteredCars = append(filteredCars, car)
				}
			}
			cars = filteredCars
		}
	}

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

	// Фильтр: model
	if modelStr := r.URL.Query().Get("model"); modelStr != "" {
		modelID, err := strconv.Atoi(modelStr)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.Model == modelID {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: body_type
	if bodyType := strings.ToLower(r.URL.Query().Get("body_type")); bodyType != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.BodyType) == bodyType {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: drive_type
	if driveType := strings.ToLower(r.URL.Query().Get("drive_type")); driveType != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.DriveType) == driveType {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: engine
	if engine := strings.ToLower(r.URL.Query().Get("engine")); engine != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.Engine) == engine {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: transmission
	if trans := strings.ToLower(r.URL.Query().Get("transmission")); trans != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.Transmission) == trans {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: color
	if color := strings.ToLower(r.URL.Query().Get("color")); color != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.Color) == color {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: dealer
	if dealerStr := r.URL.Query().Get("dealer"); dealerStr != "" {
		dealerID, err := strconv.Atoi(dealerStr)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.Dealer == dealerID {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: zero_to_hundred (ускорение)
	if zthStr := r.URL.Query().Get("zero_to_hundred"); zthStr != "" {
		zth, err := strconv.ParseFloat(zthStr, 64)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.ZeroToHundred == zth {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: num_seats
	if seatsStr := r.URL.Query().Get("num_seats"); seatsStr != "" {
		seats, err := strconv.Atoi(seatsStr)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.NumSeats == seats {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: engine_capacity
	if capStr := r.URL.Query().Get("engine_capacity"); capStr != "" {
		capacity, err := strconv.ParseFloat(capStr, 64)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.EngineCapacity == capacity {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: year
	if yearStr := r.URL.Query().Get("year"); yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err == nil {
			var filtered []models.Car
			for _, car := range cars {
				if car.Year == year {
					filtered = append(filtered, car)
				}
			}
			cars = filtered
		}
	}

	// Фильтр: price
	if minStr := r.URL.Query().Get("min_price"); minStr != "" {
		min, _ := strconv.ParseFloat(minStr, 64)
		var filtered []models.Car
		for _, car := range cars {
			if car.Price >= min {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	if maxStr := r.URL.Query().Get("max_price"); maxStr != "" {
		max, _ := strconv.ParseFloat(maxStr, 64)
		var filtered []models.Car
		for _, car := range cars {
			if car.Price <= max {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: power
	if powerStr := r.URL.Query().Get("power"); powerStr != "" {
		power, _ := strconv.Atoi(powerStr)
		var filtered []models.Car
		for _, car := range cars {
			if car.Power == power {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Фильтр: steer_wheel
	if wheel := strings.ToLower(r.URL.Query().Get("steer_wheel")); wheel != "" {
		var filtered []models.Car
		for _, car := range cars {
			if strings.ToLower(car.SteerWheel) == wheel {
				filtered = append(filtered, car)
			}
		}
		cars = filtered
	}

	// Сортировка
	sortBy := r.URL.Query().Get("sort")

	switch sortBy {
	case "price_asc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Price < cars[j].Price })
	case "price_desc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Price > cars[j].Price })
	case "year_asc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Year < cars[j].Year })
	case "year_desc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Year > cars[j].Year })
	case "power_asc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Power < cars[j].Power })
	case "power_desc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].Power > cars[j].Power })
	case "zero_to_hundred_asc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].ZeroToHundred < cars[j].ZeroToHundred })
	case "zero_to_hundred_desc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].ZeroToHundred > cars[j].ZeroToHundred })
	case "engine_capacity_asc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].EngineCapacity < cars[j].EngineCapacity })
	case "engine_capacity_desc":
		sort.Slice(cars, func(i, j int) bool { return cars[i].EngineCapacity > cars[j].EngineCapacity })
	default:
		sort.Slice(cars, func(i, j int) bool { return cars[i].Year > cars[j].Year }) // По умолчанию: по году
	}

	utils.WriteJSON(w, http.StatusOK, cars)
}
