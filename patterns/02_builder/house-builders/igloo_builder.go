package houseBuilders

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) setWindowType() {
	ib.WindowType = "igloo window"
}

func (ib *IglooBuilder) setDoorType() {
	ib.DoorType = "igloo door"
}

func (ib *IglooBuilder) setNumFloor() {
	ib.Floor = 6
}

func (ib *IglooBuilder) getHouse() House {
	return House{
		WindowType: ib.WindowType,
		DoorType:   ib.DoorType,
		Floor:      ib.Floor,
	}
}
