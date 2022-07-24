package pkg

import (
	fmt "fmt"
	"errors"
)

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	default:
		return nil, errors.New(fmt.Sprintf("unknown brand %s", brand))
	case Asus:
		return &AsusFactory{}, nil
	case Hp:
		return &HpFactory{}, nil
	}
}
