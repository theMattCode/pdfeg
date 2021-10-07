package objects

import (
	"fmt"
)

type Name struct {
	*Reference
	Value string
}

func (n Name) Label() *Reference {
	return n.Reference
}

func (n Name) AsASCIIBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("/%s", n.Value)), nil
}
