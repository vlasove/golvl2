package gunfactory

// Общий интерфейс всех продуктов
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
