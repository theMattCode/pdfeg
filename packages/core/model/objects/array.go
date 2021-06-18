package objects

import (
	"bytes"
)

type Array struct {
	*Reference
	Elements []Object
}

func (a Array) Label() *Reference {
	return a.Reference
}

func (a Array) AsASCIIBytes() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	n := len(a.Elements)
	for i, v := range a.Elements {
		asciiBytes, err := v.AsASCIIBytes()
		if err != nil {
			return nil, err
		}
		buffer.Write(asciiBytes)
		if i < (n - 1) {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}
