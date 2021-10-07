package objects

import "bytes"

type Dictionary struct {
	*Reference
	Entries map[string]Object
}

func (d Dictionary) Label() *Reference {
	return d.Reference
}

func (d Dictionary) AsASCIIBytes() ([]byte, error) {
	buffer := bytes.NewBufferString("<<")
	for k, v := range d.Entries {
		keyBytes, keyErr := Name{Value: k}.AsASCIIBytes()
		if keyErr != nil {
			return nil, keyErr
		}
		buffer.Write(keyBytes)
		buffer.WriteString(" ")

		valueBytes, valueErr := v.AsASCIIBytes()
		if valueErr != nil {
			return nil, valueErr
		}
		buffer.Write(valueBytes)
		buffer.WriteString(" ")
	}
	buffer.WriteString(">>")
	return buffer.Bytes(), nil
}
