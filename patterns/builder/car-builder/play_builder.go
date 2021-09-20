package carBuilder

type PlayCarBuilder struct {
	Seats  int
	Doors  int
	Wheels int
	Engine string
}

func NewPlayCarBuilder() *PlayCarBuilder {
	return &PlayCarBuilder{}
}

func (pcb *PlayCarBuilder) SetSeats(seats int) {
	pcb.Seats = seats + 4
}
func (pcb *PlayCarBuilder) SetDoors(doors int) {
	pcb.Doors = doors + 3
}
func (pcb *PlayCarBuilder) SetWheels(wheels int) {
	pcb.Wheels = wheels + 2
}
func (pcb *PlayCarBuilder) SetEngine(engine string) {
	pcb.Engine = engine + " playing"
}
func (pcb *PlayCarBuilder) Build() Car {
	return Car{
		Seats:  pcb.Seats,
		Doors:  pcb.Doors,
		Wheels: pcb.Wheels,
		Engine: pcb.Engine,
	}
}
