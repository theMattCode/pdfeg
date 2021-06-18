package objects

var (
	True       = Boolean{Value: true}
	False      = Boolean{Value: false}
	TrueBytes  = []byte("true")
	FalseBytes = []byte("false")
)

// Boolean represents PDF's boolean objects holding the values for its keywords true and false.
// For the most cases in building PDFs use the constants True and False if the object is direct.
// See ISO 32000-2:2017, 7.3.2.
type Boolean struct {
	*Reference
	Value bool
}

func (b Boolean) Label() *Reference {
	return b.Reference
}

func (b Boolean) AsASCIIBytes() ([]byte, error) {
	if b.Value {
		return TrueBytes, nil
	}
	return FalseBytes, nil
}
