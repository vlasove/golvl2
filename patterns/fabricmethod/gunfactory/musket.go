package gunfactory

import "fmt"

// Конкретная пушка - мушкет
type Musket struct {
	Gun
}

func NewMusket(name string, power int) IGun {
	return &Musket{
		Gun: Gun{
			Name:  name,
			Power: power,
		},
	}
}

func (m *Musket) String() string {
	return fmt.Sprintf("Gun is %s and Power is %d\n", m.Gun.Name, m.Gun.Power)
}
