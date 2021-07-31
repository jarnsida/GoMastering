package applience

import "errors"

type Applience interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
)

func CreateApplience(myType int) (Applience, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	default:
		return nil, errors.New("Invalid Appliance Type")

	}
}
