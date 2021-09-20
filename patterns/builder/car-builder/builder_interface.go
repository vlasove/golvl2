package carBuilder

type Builder interface {
	SetSeats(int)
	SetDoors(int)
	SetWheels(int)
	SetEngine(string)
	Build() Car
}
