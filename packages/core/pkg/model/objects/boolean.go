package objects

var (
	TRUE  = Boolean{value: true}
	FALSE = Boolean{value: false}
)

type Boolean struct {
	*Reference
	value bool
}

func NewIndirectBoolean(objectNumber int, generationNumber int, value bool) (*Boolean, error) {
	reference, err := NewReference(objectNumber, generationNumber)
	if err != nil {
		return nil, err
	}
	return &Boolean{reference, value}, nil
}

func (b Boolean) asBytes() ([]byte, error) {
	if b.value {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}
