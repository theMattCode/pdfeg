package objects

import (
	"fmt"
	"strconv"
)

// Number represents either an Integer or a Real object. See ISO 32000-2:2017, 7.3.3
type Number interface {
	IsReal() bool
	IsInteger() bool
}

// Integer represents PDF's integer object. See ISO 32000-2:2017, 7.3.3
type Integer struct {
	*Reference
	Value int64
}

func NewIndirectInteger(objectNumber int, generationNumber int, value int64) (*Integer, error) {
	reference, err := NewReference(objectNumber, generationNumber)
	if err != nil {
		return nil, err
	}
	return &Integer{Reference: reference, Value: value}, nil
}

func (i Integer) AsBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", i.Value)), nil
}

func (i Integer) IsReal() bool {
	return false
}

func (i Integer) IsInteger() bool {
	return true
}

// Real represents PDF's real object. See ISO 32000-2:2017, 7.3.3
type Real struct {
	*Reference
	Value float64
}

func NewIndirectReal(objectNumber int, generationNumber int, value float64) (*Real, error) {
	reference, err := NewReference(objectNumber, generationNumber)
	if err != nil {
		return nil, err
	}
	return &Real{Reference: reference, Value: value}, nil
}

func (r Real) AsBytes() ([]byte, error) {
	return []byte(strconv.FormatFloat(r.Value, 'f', -1, 64)), nil
}

func (r Real) IsReal() bool {
	return true
}

func (r Real) IsInteger() bool {
	return false
}
