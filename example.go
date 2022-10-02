package main

import (
	"fmt"
	"os"

	"github.com/gregszalay/firestore-go/firego"
	log "github.com/sirupsen/logrus"
)

var LOG_LEVEL string = os.Getenv("LOGLEVEL")

func main() {

	if LOG_LEVEL == "" {
		setLogLevel("Info")
	} else {
		setLogLevel(LOG_LEVEL)
	}

	log.Info()

	type Car struct {
		Color    string
		MaxSpeed int
	}

	newCar := Car{
		Color:    "blue",
		MaxSpeed: 200,
	}
	firego.Create("cars", newCar, "CAR001")

	newCar2 := Car{
		Color:    "green",
		MaxSpeed: 150,
	}
	firego.Create("cars", newCar2, "CAR002")

	carList := firego.ListAll("cars")
	for index, car := range carList {
		fmt.Printf("Car %d: %+v", index, car)
	}

}
