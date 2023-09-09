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