package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/structs"
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

	type Car struct {
		Color    string
		MaxSpeed int
	}

	newCar := Car{
		Color:    "blue",
		MaxSpeed: 200,
	}
	firego.Create("cars", "CAR001", structs.Map(newCar))

	newCar2 := Car{
		Color:    "green",
		MaxSpeed: 150,
	}
	firego.Create("cars", "CAR002", structs.Map(newCar2))

	carList, err := firego.ListAll("cars")
	if err != nil {
		log.Error("failed to get cars from db")
	}

	for index, element := range *carList {
		jsonStr, err := json.Marshal(element)
		if err != nil {
			log.Error("failed to marshal car list element ", index, " error: ", err)
		}
		var car Car
		if err := json.Unmarshal(jsonStr, &car); err != nil {
			log.Error("failed to unmarshal car list element ", index, " error: ", err)
		}
		fmt.Printf("Car %d: %+v\n", index, car)
	}

}

func setLogLevel(levelName string) {
	switch levelName {
	case "Panic":
		log.SetLevel(log.PanicLevel)
	case "Fatal":
		log.SetLevel(log.FatalLevel)
	case "Error":
		log.SetLevel(log.ErrorLevel)
	case "Warn":
		log.SetLevel(log.WarnLevel)
	case "Info":
		log.SetLevel(log.InfoLevel)
	case "Debug":
		log.SetLevel(log.DebugLevel)
	case "Trace":
		log.SetLevel(log.TraceLevel)
	}
}
