package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"car_db/internal/models"
)

type DB struct {
	BodyTypes     []models.BodyType     `json:"body_types"`
	DriveTypes    []models.DriveType    `json:"drive_types"`
	Engines       []models.Engine       `json:"engines"`
	Transmissions []models.Transmission `json:"transmissions"`
	Brands        []models.Brand        `json:"brands"`
	Models        []models.Model        `json:"models"`
	Colors        []models.Color        `json:"colors"`
	Users         []models.User         `json:"users"`
	Cars          []models.Car          `json:"cars"`
	CarImages     []models.CarImage     `json:"car_images"`

	mu sync.Mutex
}

var db *DB
var once sync.Once

func GetDB() *DB {
	once.Do(func() {
		db = &DB{}
		db.Load("data.json")
	})
	return db
}

func (db *DB) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read DB file: %v", err)
	}
	return json.Unmarshal(data, db)
}

func (db *DB) Save(path string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, _ := json.MarshalIndent(db, "", "  ")
	return os.WriteFile(path, data, 0644)
}
