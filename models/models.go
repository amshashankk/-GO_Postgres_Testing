package models

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model

	Name string

	License string

	Cars []Car
}

type Car struct {
	gorm.Model

	Year int

	Make string

	ModelName string

	DriverID int
}
