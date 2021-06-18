package objects_test

import (
	"github.com/stretchr/testify/assert"
	"pdfeg-core/model/objects"
	"testing"
)

var (
	expectedTrueBytes  = []byte{0x74, 0x72, 0x75, 0x65}
	expectedFalseBytes = []byte{0x66, 0x61, 0x6c, 0x73, 0x65}
)

func TestBoolean(t *testing.T) {
	assert.Equal(t, true, objects.True.Value)
	actualTrueBytes, trueBytesErr := objects.True.AsASCIIBytes()
	assert.Nil(t, trueBytesErr)
	assert.Equal(t, expectedTrueBytes, actualTrueBytes)

	assert.Equal(t, false, objects.False.Value)
	actualFalseBytes, falseBytesErr := objects.False.AsASCIIBytes()
	assert.Nil(t, falseBytesErr)
	assert.Equal(t, expectedFalseBytes, actualFalseBytes)
}

func TestIndirectBoolean(t *testing.T) {
	reference, err := objects.NewReference(1,0)
	assert.Nil(t, err)
	actual := objects.Boolean{Reference: reference, Value: true}
	assert.Equal(t, true, actual.Value)
	assert.NotNil(t, actual.Label())
	assert.Equal(t, 1, actual.ObjectNumber)
	assert.Equal(t, 0, actual.GenerationNumber)
}
