package applience

import "errors"

type Applience interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

//CreateApplience creates object of it's type
func CreateApplience(myType int) (Applience, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("invalid appliance type")

	}
}
