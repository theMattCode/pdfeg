package objects

import (
	"fmt"
	model2 "pdfeg-core/model"
)

type Object interface {
	AsBytes() ([]byte, error)
}

// Reference represents an address of an indirect object. Use NewReference instead of direct init.
// See ISO 32000-2:2017, 7.3.10.
type Reference struct {
	ObjectNumber     int
	GenerationNumber int
}

func NewReference(objectNumber int, generationNumber int) (*Reference, *model2.ValidationError) {
	if err := model2.ValidatePositiveInteger("object number", objectNumber); err != nil {
		return nil, err
	}
	if err := model2.ValidateNonNegativeInteger("generation number", generationNumber); err != nil {
		return nil, err
	}
	return &Reference{objectNumber, generationNumber}, nil
}

func (r *Reference) AsBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%d %d R", r.ObjectNumber, r.GenerationNumber)), nil
}

// Null represents PDF's null object. See ISO 32000-2:2017, 7.3.9.
type Null struct {
	*Reference
}

func NewNull() *Null {
	return &Null{}
}

func NewIndirectNull(objectNumber int, generationNumber int) (*Null, error) {
	reference, err := NewReference(objectNumber, generationNumber)
	if err != nil {
		return nil, err
	}
	return &Null{reference}, nil
}

var nullBytes = []byte{0x6e, 0x75, 0x6c, 0x6c}

func (n Null) AsBytes() ([]byte, error) {
	return nullBytes, nil
}
