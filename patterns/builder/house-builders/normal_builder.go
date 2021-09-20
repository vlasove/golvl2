package houseBuilders

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) setWindowType() {
	nb.WindowType = "normal window"
}

func (nb *NormalBuilder) setDoorType() {
	nb.DoorType = "normal door"
}

func (nb *NormalBuilder) setNumFloor() {
	nb.Floor = 3
}

func (nb *NormalBuilder) getHouse() House {
	return House{
		WindowType: nb.WindowType,
		DoorType:   nb.DoorType,
		Floor:      nb.Floor,
	}
}
