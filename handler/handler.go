package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amshashankk/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

var (
	drivers = []models.Driver{

		{Name: "Shashank", License: "India123"},

		{Name: "Tom", License: "India321"},
	}

	cars = []models.Car{

		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},

		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
	}
)

var car models.Car
var driver models.Driver

func GetCars(w http.ResponseWriter, r *http.Request) {

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	db.First(&car, params["id"])

	json.NewEncoder(w).Encode(&car)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	db.First(&driver, params["id"])

	db.Model(&driver).Related(&cars)

	driver.Cars = cars

	json.NewEncoder(w).Encode(&driver)

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	db.First(&car, params["id"])

	db.Delete(&car)

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}
