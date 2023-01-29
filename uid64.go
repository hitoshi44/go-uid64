package uid64

var defaultGenerator, ErrDefaultGeneratorCreation = NewGenerator(0)

func New() (UID, error) {
	return defaultGenerator.Gen()
}

func NewString() (string, error) {
	uid, err := defaultGenerator.Gen()
	return uid.String(), err
}
