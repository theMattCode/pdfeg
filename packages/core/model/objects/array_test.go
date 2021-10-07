package objects_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"testing"
)

var elemInt = objects.Integer{Value: 1}
var elemReal = objects.Real{Value: 2.3}
var elements = []objects.Object{elemInt, elemReal}

func TestArray_Label(t *testing.T) {
	array := objects.Array{Elements: elements}
	assert.Nil(t, array.Label())
}

func TestIndirectArray(t *testing.T) {
	reference, err := objects.NewReference(1,0)
	assert.Nil(t, err)
	array := objects.Array{Reference: reference, Elements: elements}
	assert.Equal(t, len(elements), len(array.Elements))
	assert.NotNil(t, array.Label())
	assert.Equal(t, 1, array.ObjectNumber)
	assert.Equal(t, 0, array.GenerationNumber)
}

func TestArray_AsASCIIBytes(t *testing.T) {
	array := objects.Array{Elements: elements}
	bytes, err := array.AsASCIIBytes()
	assert.Nil(t, err)
	assert.Equal(t, []byte("[1 2.3]"), bytes)
}

func TestArray_AsASCIIBytes_Error(t *testing.T) {
	expectedError := errors.New("expected")
	elemError := objMock{asASCIIBytesMock: func() ([]byte, error) {
		return nil, expectedError
	}}
	array := objects.Array{Elements: []objects.Object{elemInt, elemError}}
	bytes, err := array.AsASCIIBytes()
	assert.Error(t, err)
	assert.Nil(t, bytes)
}
