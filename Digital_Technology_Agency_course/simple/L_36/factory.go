package L_36

import (
	"errors"
	"fmt"
)

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	default:
		return nil, errors.New(fmt.Sprintf("Manufacturer %s - not found!", brand))	
	case Asus:
		return &AsusFactory{}, nil
	case Hp:
		return &HpFactory{}, nil

	}
}

