package houseBuilders

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House // Общий метод (значит все продукты будут однотипные)
}

func GetBuilder(builderType string) Builder {
	// если нужен строитель для обычного дома
	if builderType == "normal" {
		return &NormalBuilder{}
	}
	// если нужен строитель для необычного дома
	if builderType == "igloo" {
		return &IglooBuilder{}
	}
	return nil
}
