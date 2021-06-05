package objects

import (
	"fmt"
	"pdfeg-core/pkg/model"
)

type Object interface {
	asBytes() ([]byte, error)
}

// Reference represents an address of an indirect object. Use NewReference instead of direct init.
// See ISO 32000-1:2008, 7.3.10.
type Reference struct {
	objectNumber     int
	generationNumber int
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

func (r *Reference) asBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%d %d R", r.objectNumber, r.generationNumber)), nil
}

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

func (n Null) asBytes() ([]byte, error) {
	return nullBytes, nil
}
