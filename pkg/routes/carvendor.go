package routes

import (
	"github.com/XenomoprhsNightmare/car_app/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCarRoutes = func(router *mux.Router) {
	router.HandleFunc("/car/", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car/", controllers.GetCar).Methods("GET")
	router.HandleFunc("/car/{carID}", controllers.GetCarByID).Methods("GET")
	router.HandleFunc("car/{carID}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("car/{carID}", controllers.DeleteCar).Methods("DELETE")

}
