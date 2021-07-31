package applience

type Stove struct {
	typeName string
}

func (fr *Stove) Start() {
	fr.typeName = " Stove! "
}

func (fr *Stove) GetPurpose() string {
	return "I am a " + fr.typeName + "I heat stuff up!"
}
