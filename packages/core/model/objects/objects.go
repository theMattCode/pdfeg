package objects

import (
	"fmt"
	"pdfeg-core/model"
)

// Object is the base for each PDF object. PDF includes eight basic types of objects: Boolean values, Integer and Real
// numbers, Strings, Names, Arrays, Dictionaries, Streams, and the null object.
// See ISO 32000-2:2017, 7.3.1.
type Object interface {

	// Label returns a Reference or nil. Objects may be labelled so that they can be referred to by other objects. A
	// labelled object is called an indirect object.
	Label() *Reference

	// AsASCIIBytes returns a raw representation which can be for example written to a PDF file or used to logging.
	AsASCIIBytes() ([]byte, error)
}

// Reference represents a label or address of an indirect object. Use NewReference to init.
// See ISO 32000-2:2017, 7.3.10.
type Reference struct {
	ObjectNumber     int
	GenerationNumber int
}

func NewReference(objectNumber int, generationNumber int) (*Reference, *model.ValidationError) {
	if err := model.ValidatePositiveInteger("object number", objectNumber); err != nil {
		return nil, err
	}
	if err := model.ValidateNonNegativeInteger("generation number", generationNumber); err != nil {
		return nil, err
	}
	return &Reference{objectNumber, generationNumber}, nil
}

func (r *Reference) AsASCIIBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%d %d R", r.ObjectNumber, r.GenerationNumber)), nil
}

// Null represents PDF's null object. See ISO 32000-2:2017, 7.3.9.
type Null struct {
	*Reference
}

func (n Null) Label() *Reference {
	return n.Reference
}

var NullBytes = []byte{0x6e, 0x75, 0x6c, 0x6c}

func (n Null) AsASCIIBytes() ([]byte, error) {
	return NullBytes, nil
}
