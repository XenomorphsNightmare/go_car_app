package models

import (
	"/home/gergo/Documents/go_projects/crud/pkg/config"
	"crud/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	Carmake    string `gorm:""json:"carmake"`
	CarBrand   string `gorm:""json:"carbrand"`
	YearofMake int    `gorm:""json:"yearofmake"`
	Segment    string `gorm:""json:"segment"`
	SerialNo   string `gorm:""json:"serialno"`
	Vendor     Vendor
}

type Vendor struct {
	gorm.Model
	VendorName string `gorm:""json:"vendorname"`
	Address    string `gorm:""json:"adress"`
}

func init() {
	config.Connect()
	db = config.GetDB
	db.AutoMigrate(&Car{}, &Vendor{})

}


func (c *Car) CreateCar() *Car{
	db.NewRecord(c)
	db.Create(&c)
}


func GetAllCars() []Car{
	var Cars []Car 
	db.Find(&Cars)
	return Cars

}

func GetCarByID(Id int64) (*Car, *gorm.DB) {
	var getCar Car
	db:= db.Where("ID=?", Id).Find(&getCar)
	return &getCar, db 

}

func DeleteCar(ID)