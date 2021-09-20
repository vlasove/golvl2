package carBuilder

type SportCarBuilder struct {
	Seats  int
	Doors  int
	Wheels int
	Engine string
}

func NewSportCarBuilder() *SportCarBuilder {
	return &SportCarBuilder{}
}

func (scb *SportCarBuilder) SetSeats(seats int) {
	scb.Seats = seats
}
func (scb *SportCarBuilder) SetDoors(doors int) {
	scb.Doors = doors
}
func (scb *SportCarBuilder) SetWheels(wheels int) {
	scb.Wheels = wheels
}
func (scb *SportCarBuilder) SetEngine(engine string) {
	scb.Engine = engine + " very fast"
}
func (scb *SportCarBuilder) Build() Car {
	return Car{
		Seats:  scb.Seats,
		Doors:  scb.Doors,
		Wheels: scb.Wheels,
		Engine: scb.Engine,
	}
}
