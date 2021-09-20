package gunfactory

import "fmt"

// Конкретная пушка ак47
type AK47 struct {
	Gun
}

func NewAK47(name string, power int) IGun {
	return &AK47{
		Gun: Gun{
			Name:  name,
			Power: power,
		},
	}
}

func (ak *AK47) String() string {
	return fmt.Sprintf("Gun is %s and Power is %d\n", ak.Gun.Name, ak.Gun.Power)
}
