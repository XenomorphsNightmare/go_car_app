package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/XenomoprhsNightmare/car_app/pkg/models"
	"github.com/XenomorphsNightmare/car_app/pkg/models"
	"github.com/XenomorphsNightmare/car_app/pkg/utils"
	"github.com/gorilla/mux"
)

var NewCar models.Car


func GetCar(w http.ResponseWriter, r * http.Request) {
	newCars := models.GetAllCars()
	res, _ := json.Marshal(newCars)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetCarByID(w http.ResponseWriter, r * http.Request ){
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0,0)
	if err != nil{
		fmt.Println("error reading id")
	}
	carDetails, _ := models.GetCarByID(ID)
	res, _ := json.Marshal(carDetails)
	w.Header("Conten-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateCar(w http.ResponseWriter, r * http.Request) {
	CreateCar := &models.Car{}
	utils.Parsebody(r, CreateCar)
	b := CreateCar.CreateCar()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r * http.Request) {
	vars mux := mux.vars(r)
	carId = vars["carId"]
	ID, err := strconv.ParseInt(carId)
	if err != nil{
		fmt.Println("parsing error")
	}
	book := models.DeleteCar(ID)
	res, _ := json.Marshal(Car)
	w.Header().Set("Content-Type", "pklication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



