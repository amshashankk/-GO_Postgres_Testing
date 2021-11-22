package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"

	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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

var db *gorm.DB

var err error

var (
	drivers = []Driver{

		{Name: "Shashank", License: "India123"},

		{Name: "Tom", License: "India321"},
	}

	cars = []Car{

		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},

		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
	}
)

func GetCars(w http.ResponseWriter, r *http.Request) {

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	json.NewEncoder(w).Encode(&car)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var driver Driver

	var cars []Car

	db.First(&driver, params["id"])

	db.Model(&driver).Related(&cars)

	driver.Cars = cars

	json.NewEncoder(w).Encode(&driver)

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	db.Delete(&car)

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func main() {

	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=shashank sslmode=disable password=mobcoder@123")

	if err != nil {

		panic("failed to connect database")
	} else {

		fmt.Println("Successfully connected to the database")
	}

	defer db.Close()

	db.AutoMigrate(&Driver{})

	db.AutoMigrate(&Car{})

	for index := range cars {

		db.Create(&cars[index])

	}

	for index := range drivers {

		db.Create(&drivers[index])

	}

	router.HandleFunc("/cars", GetCars).Methods("GET")

	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")

	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")

	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
