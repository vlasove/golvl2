package gunfactory

import "errors"

func GetGun(gunType string, power int) (IGun, error) {
	switch gunType {
	case "ak47":
		return NewAK47("AK 47", power), nil
	case "musket":
		return NewMusket("Musket xvf9000", power), nil
	default:
		return nil, errors.New("wrong gun type passed")
	}
}
