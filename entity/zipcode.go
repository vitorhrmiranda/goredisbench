package entity

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Zipcode struct {
	ID             uint64    `json:"id"`
	Number         string    `fake:"{number:100,101}" json:"number"`
	Kind           string    `json:"kind"`
	Street         string    `json:"street"`
	Neighborhood   string    `json:"neighborhood"`
	Complement     string    `json:"complement"`
	City           string    `json:"city"`
	UF             string    `json:"uf"`
	State          string    `json:"state"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	CountryCode    string    `json:"country_code"`
	MigrateVersion int64     `json:"migrate_version"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FactoryZipcode() Zipcode {
	z := Zipcode{}
	_ = gofakeit.Struct(&z)

	z.Street = gofakeit.Street()
	z.Kind = gofakeit.AppName()
	z.Number = gofakeit.StreetNumber()
	z.City = gofakeit.City()
	z.State = gofakeit.State()
	z.Number = gofakeit.Zip()
	z.Complement = gofakeit.Sentence(3)
	z.CountryCode = gofakeit.CountryAbr()
	z.Latitude = gofakeit.Latitude()
	z.Longitude = gofakeit.Longitude()

	return z
}
