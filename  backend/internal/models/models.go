package models

import (
	"time"
)

type BodyType struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type DriveType struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type Engine struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type Transmission struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type Brand struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type Model struct {
	ModelID    int        `json:"model_id"`
	Name       string     `json:"name"`
	Brand      string     `json:"brand"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type Color struct {
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

type User struct {
	ID           int       `json:"id"`
	FName        string    `json:"f_name,omitempty"`
	LName        string    `json:"l_name,omitempty"`
	CompanyName  string    `json:"company_name,omitempty"`
	Address      string    `json:"address,omitempty"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	AccountType  string    `json:"account_type"`
	Website      string    `json:"website,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	Role		 string    `json:"role"`
}

type Car struct {
	ID             int        `json:"id"`
	Model          int        `json:"model"`
	BodyType       string     `json:"body_type"`
	DriveType      string     `json:"drive_type"`
	Engine         string     `json:"engine"`
	Transmission   string     `json:"transmission"`
	Color          string     `json:"color"`
	Dealer         int        `json:"dealer"`
	ZeroToHundred  float64    `json:"zero_to_hundred,omitempty"`
	NumSeats       int        `json:"num_seats,omitempty"`
	EngineCapacity float64    `json:"engine_capacity,omitempty"`
	Year           int        `json:"year"`
	Price          float64    `json:"price"`
	Power          int        `json:"power"`
	SteerWheel     string     `json:"steer_wheel"`
	CreatedAt      time.Time  `json:"created_at"`
	ArchivedAt     *time.Time `json:"archived_at,omitempty"`
}

type CarImage struct {
	ID        int       `json:"id"`
	CarID     int       `json:"car_id"`
	URL       string    `json:"url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
}
