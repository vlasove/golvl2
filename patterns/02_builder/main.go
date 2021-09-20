// пытаемс построить два дома - обычный и из снега
package main

import (
	"log"

	carBuilder "github.com/vlasove/materials/tasks_2/patterns/builder/car-builder"
	builders "github.com/vlasove/materials/tasks_2/patterns/builder/house-builders"
)

func main() {
	normalBuilder := builders.GetBuilder("normal")
	iglooBuilder := builders.GetBuilder("igloo")

	director := builders.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	log.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	log.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	log.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	log.Printf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	log.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	log.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

	carDirector := carBuilder.NewDirector(&carBuilder.Config{
		Builder: carBuilder.NewSportCarBuilder(),
		Wheels:  4,
		Seats:   4,
		Engine:  "V8 VAG",
		Doors:   2,
	})
	sportCar := carDirector.Build()

	log.Printf("SportCar Doors : %d\n", sportCar.Doors)
	log.Printf("SportCar Wheels : %d\n", sportCar.Wheels)
	log.Printf("SportCar seats : %d\n", sportCar.Seats)
	log.Printf("SportCar engine : %s\n", sportCar.Engine)

	carDirector.SetBuilder(carBuilder.NewPlayCarBuilder())
	playCar := carDirector.Build()

	log.Printf("PlayCar Doors : %d\n", playCar.Doors)
	log.Printf("PlayCar Wheels : %d\n", playCar.Wheels)
	log.Printf("PlayCar seats : %d\n", playCar.Seats)
	log.Printf("PlayCar engine : %s\n", playCar.Engine)
}
